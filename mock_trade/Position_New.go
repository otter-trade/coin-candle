package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_count"
	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"
)

type PositionType struct {
	GoodsId   string // OtterTrade 的 商品 ID ，从 exchange_api.GetGoodsList 获取
	TradeType string // 交易种类，可选值 global.TradeTypeList
	TradeMode string // 交易模式，可选值 global.TradeModeList
	Leverage  string // 杠杆倍率，缺省值 1 ，只有 TradeMode = SWAP 时有效
	Side      string // 下单方向，global.SideList, 只有 TradeMode = SWAP 时有效
	Amount    string // 下单金额，不可超过账户结余
}

// New 一个 Position
func NewPosition(opt global.NewPositionOpt) (resData PositionType, resErr error) {
	resErr = nil
	resData = PositionType{}

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

// 新建一个活动
type MockActionType struct {
	MockConfigFullPath string
	MockConfig         global.MockServeConfigType
	ActionTime         int64
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
	resData.MockConfigFullPath = mockPath.ConfigFullPath

	MockConfig, err := ReadMockServeInfo(mockPath.ConfigFullPath)
	if err != nil {
		resErr = err
		return
	}
	resData.MockConfig = MockConfig

	// 获取活动进行的时间
	nowTime := m_time.GetUnixInt64()
	var UpdateTime int64
	// 只有 回测模式， UpdateTime 才有效
	if MockConfig.RunMode.Value == 1 {
		UpdateTime = opt.Time
	}
	// 小于系统最老时间 或者 大于当前 则重置为当前时间
	if UpdateTime < global.TimeOldest || UpdateTime > nowTime {
		UpdateTime = nowTime
	}
	// 你不能跳回到过去下单， UpdateTime 必须大于上一次更新时间 30 秒
	lastUpdateTime := MockConfig.LastPositionUpdateTime
	if UpdateTime-lastUpdateTime < m_time.UnixTimeInt64.Seconds*30 {
		resErr = fmt.Errorf("UpdateTime必须大于上一次更新时间")
		return
	}

	resData.ActionTime = UpdateTime

	return
}
