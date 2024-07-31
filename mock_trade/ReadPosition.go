package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/global"
)

// 读取当前时间持仓情况，包括收益，策略情况评估等。
func ReadPosition(opt global.ReadPositionOpt) (resData any, resErr error) {
	// 读取 Config 信息
	MockConfig, err := GetMockServeInfo(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	// 获取 查询时间
	nowTime := m_time.GetUnixInt64()
	ReadTime := opt.Timestamp
	// 小于系统最老时间 或者 大于当前 则重置为当前时间
	if ReadTime < global.TimeOldest || ReadTime > nowTime {
		ReadTime = nowTime
	}

	fmt.Println("ReadTime", ReadTime, MockConfig.PositionIndexPath)

	return
}
