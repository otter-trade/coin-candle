package demo

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/global"
	"github.com/otter-trade/coin-candle/mock_trade"
)

// 新建一个持仓

func NewPosition_demo() {
	resData, err := mock_trade.NewPosition(global.NewPositionType{
		GoodsId:   "BTC-USDT",
		TradeMode: "SWAP",
		TradeType: "Coin",
		Leverage:  "0.134",
		Side:      "Buy",
		Amount:    "12",
	})
	if err != nil {
		fmt.Println("新建持仓失败", err)
		return
	}
	m_json.Println(resData)
}

// 更新持仓
func UpdatePosition_demo() {
	time := m_time.TimeParse(m_time.LaySP_ss, "2024-08-11 17:20:00")

	BTC_position := global.NewPositionType{
		GoodsId:   "BTC-USDT",
		TradeMode: "SWAP",
		TradeType: "Coin",
		Leverage:  "0.134",
		Side:      "Buy",
		Amount:    "12",
	}

	ETH_position := global.NewPositionType{
		GoodsId:   "ETH-USDT",
		TradeMode: "SPOT",
		TradeType: "Coin",
		Leverage:  "10",
		Side:      "Sell",
		Amount:    "10",
	}

	err := mock_trade.UpdatePosition(global.UpdatePositionOpt{
		StrategyID: "mo7_StrategyID_001",
		MockName:   "测试_MockName_1",
		UpdateTime: time,
		NewPosition: []global.NewPositionType{
			BTC_position,
			ETH_position,
		},
	})
	if err != nil {
		fmt.Println("更新持仓失败", err)
	}
}

// 读取 任意时间点 的持仓状态
func ReadPosition_demo() {
	time := m_time.TimeParse(m_time.LaySP_ss, "2024-07-30 11:00:00")
	mock_trade.ReadPosition(global.ReadPositionOpt{
		StrategyID: "mo7_StrategyID_001",
		MockName:   "测试_MockName_1",
		Timestamp:  time,
	})
}
