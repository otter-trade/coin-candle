package mock_trade

import (
	"fmt"
	"os"
	"regexp"

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

// 注册虚拟一个持仓服务
func CreatePosition(opt global.CreatePositionOpt) (resData any, resErr error) {
	resData = nil
	resErr = nil
	// 检查参数
	if len(opt.StrategyID) < 1 {
		resErr = fmt.Errorf("StrategyID 不能为空")
		return
	}

	if len(opt.MockName) < 1 {
		resErr = fmt.Errorf("MockName 不能为空")
		return
	}

	reg := regexp.MustCompile(global.MockNamePattern)
	match := reg.MatchString(opt.MockName)
	if !match {
		resErr = fmt.Errorf("MockName 只能由1-12位汉字、字母、数字、下划线组成")
		return
	}

	fmt.Println("正则结果", opt.MockName, match)
	// 存储目录为
	configFile := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		opt.StrategyID,
		os.PathSeparator,
		opt.MockName,
		os.PathSeparator,
		"config.json",
	)

	fmt.Println("配置文件为", configFile)

	// 配置文件有，则抛错，表示该虚拟持仓已建立。无，则创建，并返回成功状态。

	return
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
