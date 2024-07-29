package mock_trade

import "github.com/otter-trade/coin-candle/global"

// 更新一次仓位

var TradeModeList = global.MapAny{
	"SWAP": map[string]any{
		"description": "永续合约",
	},
	"SPOT": map[string]any{
		"description": "现货",
	},
}

var TradeTypeList = global.MapAny{
	"Coin": map[string]any{
		"description": "数字货币",
	},
}

/*

*交易对(InstName):  币安标准 BTC_USDT  欧意的标准 btc-usdt   OtterTrade的标准 BTC-USDT  格式:   xxxx-USDT
*交易模式(tradeType)： 永续合约 SWAP  现货 SPOT
杠杆倍数(Leverage)： 只有 永续合约 才存在 杠杆倍数  ， 现货这个值始终为1  缺省值为1 （计算方式: 下单资金[或下单数量]*杠杆倍数）
下单方向(Side)： 买多，买空， null
开仓资金(xxx)： 开仓资金 (如果开仓资金>实际余额) 报错并返回说明无效参数原因，你没有那么多资金

如果是回测的话
只需要加上下单的时间即可
Timestamp(毫秒)：比如  2024-06-08 14:00:00 的unix 时间戳 ; 如果为 0 则为当前时间
这个时候应该采用 当前交易对 这个时间点的价格进行计算

*/

type NewPositionType struct {
	GoodsId   string // OtterTrade 的交易品ID 以 OKX 为准 如 BTC-USDT
	TradeMode string // 交易模式，缺省值 SPOT 可选值 TradeModeList
	TradeType string // 交易种类，缺省值 Coin 可选值 TradeTypeList
	Leverage  string // 杠杆倍率，缺省值 1 ，只有 TradeMode = SWAP 时有效
	Side      string // 下单方向 Buy , Sell
	Amount    string // 下单金额
}

type UpdatePositionOpt struct {
	StrategyID  string // 策略的Id
	MockName    string // 本次回测的名称
	Timestamp   int64  // 更新本次仓位的时间(13位毫秒时间戳)，只有在 RunType 为 1 时 才会读取。也就是只有在回测模式下才允许在任意时间更新仓位，否则只能在当前时间点更新仓位。
	NewPosition []NewPositionType
}

// 更新一次仓位状态
func UpdatePosition(opt UpdatePositionOpt) {
}
