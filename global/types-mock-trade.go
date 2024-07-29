package global

import (
	"fmt"
	"os"
	"regexp"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_str"
)

// 2-24位字母数字下划线和中文
func IsMockNameReg(str string) bool {
	pattern := "^[a-zA-Z0-9_\u4e00-\u9fa5]{2,24}$"
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(str)
}

type MockPathType struct {
	ConfigPath  string
	MockDataDir string
}

func CheckMockName(opt FindPositionOpt) (resData MockPathType, resErr error) {
	resData = MockPathType{}
	resErr = nil
	// StrategyID 不能为空
	if len(opt.StrategyID) < 1 {
		resErr = fmt.Errorf("StrategyID 不能为空")
		return
	}
	// MockName 必须为2-24位字母数字下划线和中文
	isMockNameReg := IsMockNameReg(opt.MockName)
	if !isMockNameReg {
		resErr = fmt.Errorf("MockName必须为2-24位字母数字下划线和中文")
		return
	}

	MockDataDir := m_str.Join(
		Path.MockTradeDir,
		os.PathSeparator,
		opt.StrategyID,
		os.PathSeparator,
		opt.MockName,
		os.PathSeparator,
	)

	ConfigPath := m_str.Join(
		MockDataDir,
		"config.json",
	)

	resData = MockPathType{
		ConfigPath:  ConfigPath,
		MockDataDir: MockDataDir,
	}

	return
}

// map 类型
type MapAny map[string]any

type CreatePositionOpt struct {
	StrategyID   string // 策略的Id，从 OtterTrade 用户数据中读取，不可为空
	MockName     string // 模拟交易的名称，策略开发者自定义，不可为空，1-12位中文，字母或数字。
	RunMode      string // 运行模式，缺省值 1 ，可选值 RunTypeList
	InitialAsset string // 初始资产(USDT)  最小值/缺省值 1000
	FeeRate      string // 手续费率 缺省值 0.001 参考 https://www.okx.com/zh-hans/fees
}

type RunModeType struct {
	Key         int
	Description string
}

// 运行模式可选参数
var RunModeList = []RunModeType{
	{
		Key:         1,
		Description: "回测模式，交易只会在本地产生订单，方便查询和以及辅助策略的修改。",
	},
	{
		Key:         2,
		Description: "社区模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。需要OtterTrade平台的权限。",
	},
	{
		Key:         3,
		Description: "生产模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。策略需要通过 OtterTrade 平台的审核。",
	},
}

func GetRunMode(key string) (resData RunModeType, resErr error) {
	if len(key) < 1 {
		key = "1" // 缺省值为 1
	}
	resErr = nil
	for _, v := range RunModeList {
		if m_count.Le(key, m_str.ToStr(v.Key)) == 0 {
			resData = v
			break
		}
	}
	if resData.Key < 1 {
		resErr = fmt.Errorf("RunMode 不正确")
		return
	}
	return
}

type PositionConfigType struct {
	CreatePositionOpt
	RunMode RunModeType
}

type FindPositionOpt struct {
	StrategyID string
	MockName   string
}
