package mock_trade

type AddPositionOpt struct {
	GoodsId   string
	TradeType string // 交易种类，可选值 global.TradeTypeList
	TradeMode string // 交易模式，可选值 global.TradeModeList
	Leverage  string // 杠杆倍率，缺省值 1 ，只有 TradeMode = SWAP 时有效
	Side      string // 下单方向，global.SideList, 只有 TradeMode = SWAP 时有效
	Amount    string // 下单金额，不可超过账户结余
}

// ## 添加一个仓位 ##
func (obj *MockActionObj) AddPosition(opt AddPositionOpt) (resErr error) {
	resErr = nil

	return
}
