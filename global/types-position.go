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

var TradeModeList = []KeyDescType{
	{
		Value:       "SPOT",
		Description: "现货，买入卖出赚取差价。",
	},
	{
		Value:       "SWAP",
		Description: "永续合约，杠杆做多，卖出做空。",
	},
}

func GetTradeMode(Value string) (resData KeyDescType, resErr error) {
	resErr = nil
	for _, v := range TradeModeList {
		if v.Value == Value {
			resData = v
			break
		}
	}
	if len(resData.Value) < 1 {
		resErr = fmt.Errorf("TradeMode不正确")
		return
	}
	return
}

// 交易种类
var TradeTypeList = []KeyDescType{
	{
		Value:       "Coin",
		Description: "数字货币，成熟且自由的交易市场，允许永续合约做多做空",
	},
}

func GetTradeType(Value string) (resData KeyDescType, resErr error) {
	resErr = nil
	for _, v := range TradeTypeList {
		if v.Value == Value {
			resData = v
			break
		}
	}
	if len(resData.Value) < 1 {
		resErr = fmt.Errorf("TradeType不正确")
		return
	}
	return
}

// 更新一次持仓状态
type NewPositionType struct {
	GoodsId   string // OtterTrade 的 商品 ID ，从 exchange_api.GetGoodsList 获取
	TradeMode string // 交易模式，缺省值 SPOT 可选值 TradeModeList
	TradeType string // 交易种类，可选值 TradeTypeList
	Leverage  string // 杠杆倍率，缺省值 1 ，只有 TradeMode = SWAP 时有效
	Side      string // 下单方向 Buy , Sell , 只有 TradeMode = SWAP 时有效
	Amount    string // 下单金额，不可超过账户结余
}

type UpdatePositionOpt struct {
	StrategyID  string            // 策略的Id
	MockName    string            // 本次回测的名称
	UpdateTime  int64             // 更新本次仓位的时间(13位毫秒时间戳)，只有在 RunType 为 1 时 才会读取。也就是只有在回测模式下才允许在任意时间更新仓位，否则只能在当前时间点更新仓位。
	NewPosition []NewPositionType // 允许多个不同品类的仓位持仓，空代表清空所有仓位。
}

var MaxLeverage = "30" // 支持的最大杠杆倍率

type ReadPositionOpt struct {
	StrategyID string
	MockName   string
	Timestamp  int64 // 读取任意时间点的持仓情况(13位毫秒时间戳)，0 或 空 则为当前时间。
}

// 持仓结算
type PositionClose struct{}
