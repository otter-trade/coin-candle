package mock_trade

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"
)

// 更新一次仓位状态
func UpdatePosition(opt global.UpdatePositionOpt) (resErr error) {
	resErr = nil

	// 读取 Config 信息
	MockConfig, err := GetMockServeInfo(global.FindMockServeOpt{
		StrategyID: opt.StrategyID,
		MockName:   opt.MockName,
	})
	if err != nil {
		resErr = err
		return
	}

	// UpdateTime
	nowTime := m_time.GetUnixInt64()
	var UpdateTime int64
	// 回测模式， UpdateTime 才有效
	if MockConfig.RunMode.Value == 1 {
		UpdateTime = opt.UpdateTime
	}
	// 小于系统最老时间 和 大于当前
	if UpdateTime < global.TimeOldest || UpdateTime > nowTime {
		UpdateTime = nowTime
	}

	var NewPosition []global.NewPositionType

	for _, item := range opt.NewPosition {
		if len(item.GoodsId) > 1 {
			position, err := NewPositionFunc(item)
			if err != nil {
				resErr = fmt.Errorf("%+v,%+v", item.GoodsId, err) // 只要有一个持仓有问题，则该次持仓判定为失效
				return
			}
			NewPosition = append(NewPosition, position)
		}
	}

	fmt.Println("UpdateTime", UpdateTime, NewPosition)

	return
}

func NewPositionFunc(opt global.NewPositionType) (resData global.NewPositionType, resErr error) {
	resData = global.NewPositionType{}
	resErr = nil

	// 检查参数
	_, err := global.GetTradeType(opt.TradeType)
	if err != nil {
		resErr = err
		return
	}
	resData.TradeType = opt.TradeType

	GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
		GoodsId: opt.GoodsId,
	})
	if err != nil {
		resErr = err
		return
	}
	if GoodsDetail.State != "live" {
		resErr = fmt.Errorf("%+v,%+v", GoodsDetail.GoodsId, GoodsDetail.State)
		return
	}
	resData.GoodsId = GoodsDetail.GoodsId

	TradeModeValue := opt.TradeMode
	if len(TradeModeValue) > 1 {
		_, err := global.GetTradeMode(TradeModeValue)
		if err != nil {
			resErr = err
			return
		}
	} else {
		TradeModeValue = global.TradeModeList[0].Value
	}

	resData.TradeMode = TradeModeValue

	m_json.Println(resData)

	// fmt.Println("item", item, GoodsDetail, TradeMode)

	return
}
