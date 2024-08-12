package lib_x

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/global"
	"github.com/otter-trade/coin-candle/mock_trade"
)

// 读取当前时间持仓情况，包括收益，策略情况评估等。
func ReadPosition(opt global.ReadPositionOpt) (resData any, resErr error) {
	// 读取 Config 信息
	MockConfig, err := mock_trade.GetMockServeInfo(global.FindMockServeOpt{
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

// 读取某个时间段内的仓位记录列表
type GetHistoryPositionOpt struct {
	StrategyID string
	MockName   string
	StartTime  int64 // 开始时间
	EndTime    int64 // 结束时间
}

func GetPositionList(opt GetHistoryPositionOpt) {
	// 读取持仓历史
}

// 查询一个MockServe 下的持仓列表
func GetMockServePosition() {
}
