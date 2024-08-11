package mock_trade

import (
	"fmt"
	"os"

	"github.com/handy-golang/go-tools/m_str"
	"github.com/otter-trade/coin-candle/global"
)

// 删除一个 MockServe
func DeleteMockServe(opt global.FindMockServeOpt) (resErr error) {
	resErr = nil
	// 检查 StrategyID 和 MockName 并获取存储目录
	mockPath, err := CheckMockName(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	err = os.RemoveAll(mockPath.MockDataFullDir)
	if err != nil {
		resErr = err
		return
	}

	return
}

// 删除一个 Strategy
func ClearStrategy(StrategyID string) (resErr error) {
	resErr = nil
	StrategyDir := m_str.Join(
		global.Path.MockTradeDir,
		os.PathSeparator,
		StrategyID,
	)
	files, err := os.ReadDir(StrategyDir)
	if err != nil {
		resErr = err
		return
	}
	if len(files) > 0 {
		resErr = fmt.Errorf("请先删除该策略下的所有 MockServe")
		return
	}

	err = os.RemoveAll(StrategyDir)
	if err != nil {
		resErr = err
		return
	}
	return
}
