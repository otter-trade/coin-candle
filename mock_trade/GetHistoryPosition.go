package mock_trade

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
