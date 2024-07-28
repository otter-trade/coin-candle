package mock_trade

import (
	"fmt"
	"os"

	"github.com/handy-golang/go-tools/m_str"
	"github.com/otter-trade/coin-candle/global"
)

/*
	模拟交易模块。
	设计了一套仓位管理模式，可以将策略逻辑和交易所 进行有效隔离。内建模拟交易好处:
	1. 有效简化策略开发流程。因为是内建的虚拟仓位，所以在编写策略时，只需要关注策略逻辑即可。
	2. 方便回测。在回测模式下，开发者可以在1分钟之内完成 一整年的交易订单的模拟，并查看结果，方便策略的修改。
	3. 回测-实盘 无缝对接。开发者在写完策略后，只需要简单修改一个参数，即可上架实盘。保证实盘效果和回测结果的无限接近。

	基本逻辑：
  注册一个虚拟持仓，并标注初始资产。
	不断地更新仓位，系统会记录下仓位变化，如果为 实盘模式 则同步给下单模块更新至交易所。
	可以读取任意时间点的仓位情况，包括当前时间点的收益。并根据持仓情况，对这个时间点的策略情况进行打分和评估。

	支持读取全部的持仓记录。

	基于本地文件系统，支持超高频率的读取和更新。
*/

type MapAny map[string]any

// 运行模式可选参数
var RunTypeList = MapAny{
	"1": MapAny{
		"description": "回测模式，交易只会在本地产生订单，方便查询和以及辅助策略的修改。",
	},
	"2": map[string]any{
		"description": "社区模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。需要OtterTrade平台的权限。",
	},
	"3": map[string]any{
		"description": "生产模式，交易不仅会在本地产生订单，还会同时将订单结果同步至交易所完成下单。策略需要通过 OtterTrade 平台的审核。",
	},
}

// 注册虚拟一个持仓服务
type CreatePositionOpt struct {
	StrategyID   string // 策略的Id，从 OtterTrade 用户数据中读取，不可为空
	MockName     string // 模拟交易的名称，策略开发者自定义，不可为空，1-12位中文，字母和数字。
	RunType      string // 虚拟持仓的运行模式，缺省值 1 ，可选值 RunTypeList
	InitialAsset string // 初始资产(USDT)  缺省值 10000
	FeeRate      string // 手续费率 缺省值 0.001 参考 https://www.okx.com/zh-hans/fees
}

func CreatePosition(opt CreatePositionOpt) {
	// 存储目录为
	dir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		opt.StrategyID,
		os.PathSeparator,
		opt.MockName,
	)
	fmt.Println("存储目录为", dir)

	// 配置文件为
	configFile := m_str.Join(
		dir,
		os.PathSeparator,
		"config.json",
	)
	fmt.Println("配置文件为", configFile)

	// 配置文件有，则抛错，表示该虚拟持仓已建立。无，则创建，并返回成功状态。
}

// 删除一个持仓服务
type DeletePositionOpt struct {
	StrategyID string // 策略的Id,
	MockName   string // 需要删除
}

func DeletePosition(opt DeletePositionOpt) {
	// RunType 1 则可以随便删除
	// RunType 2 或者 3 则需要权限，才能删除。
}
