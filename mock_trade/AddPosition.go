package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"
)

// ## 添加一个仓位 ##
/*
GoodsId   string // OtterTrade 的 商品 ID ，从 exchange_api.GetGoodsList 获取
TradeType string // 交易种类，可选值 global.TradeTypeList
TradeMode string // 交易模式，可选值 global.TradeModeList
Leverage  string // 杠杆倍率，缺省值 1 ，只有 TradeMode = SWAP 时有效
Side      string // 下单方向，global.SideList, 只有 TradeMode = SWAP 时有效
Amount    string // 下单金额，不可超过账户结余
*/

type NewPositionType struct {
	GoodsDetail global.GoodsType // OtterTrade 的 商品 详情
	GoodsId     string           // OtterTrade 的交易品ID 以 OKX 为准 如 BTC-USDT
	TradeType   string           // 交易种类，Coin
	TradeMode   string           // 持仓模式，SPOT  SWAP
	Leverage    string           // 杠杆倍率，1-30
	Side        string           // 下单方向，Buy 和 Sell
	Amount      string           // 下单金额，
}

// 添加一个新的持仓
func (obj *MockActionObj) NewPositionAdd(opt global.NewPositionType) (resErr error) {
	resErr = nil
	position := NewPositionType{}

	// 检查参数 TradeType
	TradeType_obj, err := global.GetKeyDescObj(opt.TradeType, global.TradeTypeList)
	if err != nil {
		resErr = err
		return
	}
	position.TradeType = TradeType_obj.Value
	// 检查 TradeMode
	TradeMode_obj, err := global.GetKeyDescObj(opt.TradeMode, global.TradeModeList)
	if err != nil {
		resErr = err
		return
	}
	position.TradeMode = TradeMode_obj.Value

	if TradeMode_obj.Value == "SWAP" {
		// 检查 Side
		Side_obj, err := global.GetKeyDescObj(opt.Side, global.SideList)
		if err != nil {
			resErr = err
			return
		}
		position.Side = Side_obj.Value

		// 检查 Leverage
		Leverage := m_count.Sub(opt.Leverage, "0")
		Leverage = m_count.Cent(Leverage, 0)
		if m_count.Le(Leverage, "1") < 0 {
			Leverage = "1"
		}
		if m_count.Le(Leverage, global.MaxLeverage) > 0 {
			Leverage = global.MaxLeverage
		}
		position.Leverage = Leverage
	} else {
		position.Side = "Buy"
		position.Leverage = "1"
	}

	// 检查 GoodsId
	GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
		GoodsId: opt.GoodsId,
	})
	if err != nil {
		resErr = err
		return
	}
	position.GoodsDetail = GoodsDetail
	position.GoodsId = GoodsDetail.GoodsId

	Amount := m_count.Sub(opt.Amount, "0")
	if m_count.Le(Amount, "0") < 0 {
		Amount = "0" // 最小值为 0
	}
	position.Amount = Amount

	for _, item := range obj.NewPosition {
		if item.GoodsId == position.GoodsId {
			resErr = fmt.Errorf("当前商品已添加")
			return
		}
	}
	// 这里应当计算一下余额还有多少

	obj.NewPosition = append(obj.NewPosition, position)
	return
}
