package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"
)

type NewPositionType struct {
	GoodsDetail global.GoodsType // OtterTrade 的 商品 详情
	TradeType   string           // 交易种类，Coin
	TradeMode   string           // 持仓模式，SPOT  SWAP
	Leverage    string           // 杠杆倍率，1-30
	Side        string           // 下单方向，Buy 和 Sell
	Amount      string           // 下单金额，
}

// New 一个 Position
func NewPositionParam(opt global.NewPositionType) (resData NewPositionType, resErr error) {
	resErr = nil
	resData = NewPositionType{}

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
	resData.GoodsDetail = GoodsDetail

	// 买入金额
	Amount := m_count.Sub(opt.Amount, "0")
	if m_count.Le(Amount, "0") < 0 {
		Amount = "0" // 最小值为 0
	}
	resData.Amount = Amount

	return
}

// 新建一个活动
type MockActionType struct {
	MockPath   MockPathType
	MockConfig global.MockServeConfigType
	UpdateTime int64
}

// New 一个新的 Action
func NewMockAction(opt global.NewMockActionOpt) (resData MockActionType, resErr error) {
	resErr = nil
	resData = MockActionType{}
	// 读取 Config 相关的目录信息
	mockPath, err := CheckMockName(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}
	resData.MockPath = mockPath

	MockConfig, err := ReadMockServeInfo(mockPath.ConfigFullPath)
	if err != nil {
		resErr = err
		return
	}
	resData.MockConfig = MockConfig

	// 获取活动进行的时间
	var UpdateTime int64
	nowTime := m_time.GetUnixInt64()
	// 只有 回测模式， UpdateTime 才有效
	if MockConfig.RunMode.Value == 1 {
		UpdateTime = opt.Time
	}

	if UpdateTime == 0 {
		// 如果时间为 0 , 则为当前时间
		UpdateTime = nowTime
	} else {
		//  小于系统最老时间
		if UpdateTime < global.TimeOldest {
			UpdateTime = global.TimeOldest
		}
	}
	// 或者 大于当前 则重置为当前时间
	if UpdateTime > nowTime {
		UpdateTime = nowTime
	}

	lastUpdateTime := MockConfig.LastPositionUpdateTime
	if UpdateTime-lastUpdateTime > m_time.UnixTimeInt64.Seconds*30 {
		// 本次更新时间 - 上次更新时间 必须 大于 30 秒
		// UpdateTime - lastUpdateTime > 30s
	} else {
		resErr = fmt.Errorf("UpdateTime必须大于上一次更新时间")
		return
	}

	resData.UpdateTime = UpdateTime

	return
}
