package demo

/*
MockServe 相关的 行为 api


*/

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/otter-trade/coin-candle/global"
	"github.com/otter-trade/coin-candle/mock_trade"
)

// MockAction 流程
func MockAction_demo() {
	// New 一个 Action 对象
	ActionObj, err := mock_trade.NewMockAction(mock_trade.NewMockActionOpt{
		StrategyID: "mo7_StrategyID_002",
		MockName:   "测试_MockName_1",
	})
	if err != nil {
		fmt.Println("新建 MockAction 失败", err)
		return
	}

	// 读取配置文件的方法
	err = ActionObj.ReadMockServeConfig()
	if err != nil {
		fmt.Println("读取 MockServeConfig 失败", err)
		return
	}

	// 存储配置文件的方法 将配置文件更新到本地
	err = ActionObj.StoreMockServeConfig()
	if err != nil {
		fmt.Println("写入保存 MockServeConfig 失败", err)
		return
	}

	// 读取 PositionIndex 也就是 持仓列表的索引
	err = ActionObj.ReadPositionIndex()
	if err != nil {
		fmt.Println("写入保存 MockServeConfig 失败", err)
		return
	}

	// 新添加一个持仓
	err = ActionObj.NewPositionAdd(global.NewPositionType{
		GoodsId:   "BTC-USDT", // GoodsId
		TradeMode: "SWAP",     //  交易方式  SWAP  永续合约
		TradeType: "Coin",     //  交易类型  Coin  币币
		Leverage:  "1",        //  杠杆倍率
		Side:      "Buy",      // 方向，买入
		Amount:    "12",       // 买入金额
	})
	if err != nil {
		fmt.Println("添加仓位失败1", err)
		return
	}
	// 再添加另一个持仓 , (可以无限之的添加)
	err = ActionObj.NewPositionAdd(global.NewPositionType{
		GoodsId:   "ETH-USDT",
		TradeMode: "SWAP",
		TradeType: "Coin",
		Leverage:  "1",
		Side:      "Sell",
		Amount:    "24",
	})
	if err != nil {
		fmt.Println("添加仓位失败2", err)
		return
	}

	err = ActionObj.OpenPosition(mock_trade.OpenPositionOpt{
		OrderTime: 0,
	})
	if err != nil {
		fmt.Println("开仓失败", err)
		return
	}

	m_file.Write(global.Path.LogPath+"/ActionObj.json", m_json.ToStr(ActionObj))

	// 下单
}
