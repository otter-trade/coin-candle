package global

/*
所有和 MockServe 相关的 type 定义都放在这里
*/

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_str"
)

// map 类型
type MapAny map[string]any

type CreateMockServeOpt struct {
	StrategyID   string // 策略的Id，从 OtterTrade 用户数据中读取，不可为空
	MockName     string // 模拟交易的名称，策略开发者自定义，不可为空，1-12位中文，字母或数字。
	RunMode      string // 运行模式，缺省值 1 ，可选值 RunTypeList
	InitialAsset string // 初始资产(USDT)  最小值/缺省值 1000
	FeeRate      string // 手续费率 缺省值 0.001(0.1%)  参考 https://www.okx.com/zh-hans/fees
	Description  string // 用户对本次 Mock 的描述，缺省值 空
}

var DefaultFeeRate = "0.001" // 默认手续费率

var DefaultInitialAsset = "1000" // 默认的初始资产

type RunModeType struct {
	Value       int
	Description string
}

// 运行模式可选参数
var RunModeList = []RunModeType{
	{
		Value:       1,
		Description: "回测模式，交易只会在本地产生订单，方便查询和以及辅助策略的修改。",
	},
	{
		Value:       2,
		Description: "社区模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。需要OtterTrade平台的权限。",
	},
	{
		Value:       3,
		Description: "生产模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。策略需要通过 OtterTrade 平台的审核。",
	},
}

func GetRunMode(key string) (resData RunModeType, resErr error) {
	if len(key) < 1 {
		key = "1" // 缺省值为 1
	}
	resErr = nil
	for _, v := range RunModeList {
		if m_count.Le(key, m_str.ToStr(v.Value)) == 0 {
			resData = v
			break
		}
	}
	if resData.Value < 1 {
		resErr = fmt.Errorf("RunMode 不正确")
		return
	}
	return
}

type MockServeConfigType struct {
	CreateMockServeOpt
	RunMode                RunModeType
	MockDataDir            string // 数据存储路径
	ConfigPath             string // 配置文件存放路径
	PositionIndexPath      string // 仓位索引存放路径
	CreateTime             int64  // 创建时间，方便查询
	LastPositionUpdateTime int64  // 最后一次更新持仓的时间，方便查询
	SysWarn                string // 系统提示
}

type FindMockServeOpt struct {
	StrategyID string
	MockName   string
}

type FindMockServeListOpt struct {
	StrategyID             string
	CreateTime             [2]int64 // 依据创建时间范围查询 开始-结束时间的 13 位毫秒时间戳
	LastPositionUpdateTime [2]int64 // 依据 最后一次更新持仓 的时间范围查询
}

type PositionIndexType []int64

var MaxMockServeCount = 60 // 每个策略允许的最大 MockServe 数量，将来可以适当增加。
