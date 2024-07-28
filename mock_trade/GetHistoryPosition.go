package mock_trade

// 读取某个时间段内的仓位记录列表
type GetHistoryPositionOpt struct {
	StrategyID string
	MockName   string
	StartTime  int64 // 开始时间
	EndTime    int64 // 结束时间
}

func GetHistoryPosition(opt GetHistoryPositionOpt) {
	// 读取持仓历史
}
