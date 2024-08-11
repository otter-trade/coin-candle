package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"
)

func NewPosition(opt global.NewPositionOpt) (resData UpdatePositionType, resErr error) {
	resErr = nil
	resData = UpdatePositionType{}

	// 检查参数 TradeType
	TradeType_obj, err := global.GetKeyDescObj(opt.TradeType, global.TradeTypeList)
	if err != nil {
		resErr = err
		return
	}
	resData.TradeType = TradeType_obj.Value

	// 检查 TradeMode
	TradeMode_obj, err := global.GetKeyDescObj(opt.TradeMode, global.TradeModeList)
	if err != nil {
		resErr = err
		return
	}
	resData.TradeMode = TradeMode_obj.Value

	// 当 TradeMode 为 SWAP 时
	if TradeMode_obj.Value == "SWAP" {
		// 检查 Side
		Side_obj, err := global.GetKeyDescObj(opt.Side, global.SideList)
		if err != nil {
			resErr = err
			return
		}
		resData.Side = Side_obj.Value

		// 检查 Leverage
		// 将来每个币种都会有自己独立的最大持仓范围
		Leverage := m_count.Sub(opt.Leverage, "0")
		Leverage = m_count.Cent(Leverage, 0)
		if m_count.Le(Leverage, "1") < 0 {
			Leverage = "1"
		}
		if m_count.Le(Leverage, global.MaxLeverage) > 0 {
			Leverage = global.MaxLeverage
		}
		resData.Leverage = Leverage
	} else {
		resData.Side = "Buy"
		resData.Leverage = "1"
	}

	// 检查 GoodsId
	GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
		GoodsId: opt.GoodsId,
	})
	if err != nil {
		resErr = err
		return
	}
	if GoodsDetail.State != "live" {
		resErr = fmt.Errorf("该商品交易状态存在问题")
		return
	}

	resData.GoodsId = GoodsDetail.GoodsId

	// 买入金额
	Amount := m_count.Sub(opt.Amount, "0")
	if m_count.Le(Amount, "0") < 0 {
		Amount = "0" // 最小值为 0
	}
	resData.Amount = Amount

	return
}
