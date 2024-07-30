package mock_trade

import (
	"github.com/handy-golang/go-tools/m_json"
	"github.com/otter-trade/coin-candle/global"
)

// 更新一次仓位状态
func UpdatePosition(opt global.UpdatePositionOpt) (resErr error) {
	resErr = nil

	// 读取 Config 信息
	MockConfig, err := GetMockServeInfo(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	m_json.Println(MockConfig)

	return
}
