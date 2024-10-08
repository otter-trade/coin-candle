package global

/*
所有和持仓相关的 type 定义
*/

import "fmt"

// 交易模式
type KeyDescType struct {
	Value       string
	Description string
}

func GetKeyDescObj(value string, list []KeyDescType) (resData KeyDescType, resErr error) {
	resErr = nil
	resData = KeyDescType{}
	for _, v := range list {
		if v.Value == value {
			resData = v
			break
		}
	}
	if len(resData.Value) < 1 {
		resErr = fmt.Errorf("参数 %+v 不正确", value)
		return
	}
	return
}

// 持仓模式
var TradeModeList = []KeyDescType{
	{
		Value:       "SPOT",
		Description: "现货买入，现货卖出赚取差价。",
	},
	{
		Value:       "SWAP",
		Description: "永续合约，杠杆借币，买入做多，卖出做空。",
	},
}

// 持仓种类 将来会增加 股票、期货 等
var TradeTypeList = []KeyDescType{
	{
		Value:       "Coin",
		Description: "数字货币，成熟且自由的交易市场，允许永续合约做多做空",
	},
}

// 持仓方向
var SideList = []KeyDescType{
	{
		Value:       "Buy",
		Description: "现货中代表买入，合约中代表做多",
	},
	{
		Value:       "Sell",
		Description: "现货中代表卖出，合约中代表做空",
	},
}

// 新建一个持仓
type NewPositionType struct {
	GoodsId   string // OtterTrade 的 商品 ID ，从 exchange_api.GetGoodsList 获取
	TradeType string // 交易种类，可选值 global.TradeTypeList
	TradeMode string // 交易模式，可选值 global.TradeModeList
	Leverage  string // 杠杆倍率，缺省值 1 ，只有 TradeMode = SWAP 时有效
	Side      string // 下单方向，global.SideList, 只有 TradeMode = SWAP 时有效
	Amount    string // 下单金额，不可超过账户结余
}

var MaxLeverage = "30" // 支持的最大杠杆倍率
