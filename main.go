package main

import (
	"fmt"
	"time"

	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/exchange_api/binance"
	"github.com/otter-trade/coin-candle/exchange_api/okx"
	"github.com/otter-trade/coin-candle/global"
	"github.com/otter-trade/coin-candle/mock_trade"

	"github.com/handy-golang/go-tools/m_cycle"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_time"
)

func main() {
	SysInit() //  系统的初始化
	// KlineActionTest()
	// MarketFunc()    // 市场相关的数据
	// KlineFunc()     // K线相关的数据
	// MockServeFunc() // MockServe 的增删查
	PositionFunc() // 持仓管理
}

func SysInit() {
	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	m_cycle.New(m_cycle.Opt{
		Func: func() {
			// // 更新本地的商品列表
			// exchange_api.UpdateLocalGoodsList()
			// // 更新本地的榜单
			// exchange_api.UpdateLocalTicker()
			// 读取获取商品列表 并存入内存中
			exchange_api.GetGoodsList()
		},
		SleepTime: time.Hour * 4, // 执行一次后 每隔 4h 再执行一次
	}).Start()
}

func MockServeFunc() {
	//  ####### 创建一个 MockServe #######
	// _, err := mock_trade.CreateMockServe(global.CreateMockServeOpt{
	// 	StrategyID:   "mo7_StrategyID_001",
	// 	MockName:     "测试_MockName_3",
	// 	RunMode:      "1",
	// 	InitialAsset: "1000",
	// })
	// if err != nil {
	// 	fmt.Println("创建持仓失败", err)
	// }

	// #######  查看 MockServe 列表 #######
	mockServeList := mock_trade.GetMockServeList(global.FindMockServeListOpt{
		StrategyID: "mo7_StrategyID_001",
	})
	m_json.Println(mockServeList)

	// #######  查看 MockServe 详情 #######
	// mockServeInfo, err := mock_trade.GetMockServeInfo(global.FindMockServeOpt{
	// 	StrategyID: "mo7_StrategyID_001",
	// 	MockName:   "测试_MockName_1",
	// })
	// if err != nil {
	// 	fmt.Println("获取MockServe信息失败", err)
	// }
	// m_json.Println(mockServeInfo)

	// ####### 删除一个 MockServe #######
	// err := mock_trade.DeleteMockServe(global.FindMockServeOpt{
	// 	StrategyID: "mo7_StrategyID_001",
	// 	MockName:   "测试_MockName_4",
	// })
	// if err != nil {
	// 	fmt.Println("删除虚拟持仓失败", err)
	// }

	// ####### 删除一个 策略 #######
	// err := mock_trade.ClearStrategy("mo7_StrategyID_001")
	// if err != nil {
	// 	fmt.Println("删除策略失败", err)
	// }
}

// 更新持仓状态
func PositionFunc() {
	time := m_time.TimeParse(m_time.LaySP_ss, "2024-07-30 11:00:00")
	// err := mock_trade.UpdatePosition(global.UpdatePositionOpt{
	// 	StrategyID: "mo7_StrategyID_001",
	// 	MockName:   "测试_MockName_1",
	// 	UpdateTime: time,
	// 	NewPosition: []global.NewPositionType{
	// 		{
	// 			GoodsId:   "BTC-USDT",
	// 			TradeMode: "SWAP",
	// 			TradeType: "Coin",
	// 			Leverage:  "0.134",
	// 			Side:      "Buy",
	// 			Amount:    "12",
	// 		},
	// 		{
	// 			GoodsId:   "ETH-USDT",
	// 			TradeMode: "SPOT",
	// 			TradeType: "Coin",
	// 			Leverage:  "10",
	// 			Side:      "Sell",
	// 			Amount:    "10",
	// 		},
	// 	},
	// })
	// if err != nil {
	// 	fmt.Println("更新持仓失败", err)
	// }

	// 读取 任意时间点 的持仓结果
	mock_trade.ReadPosition(global.ReadPositionOpt{
		StrategyID: "mo7_StrategyID_001",
		MockName:   "测试_MockName_1",
		Timestamp:  time,
	})
}

// 市场数据
func MarketFunc() {
	// ####### 获取商品列表 #######
	GoodsList, err := exchange_api.GetGoodsList()
	if err != nil {
		fmt.Println("获取商品列表失败", err)
	}
	fmt.Println("GoodsList 商品总数:", len(GoodsList))

	// ####### 获取商品详情 #######
	GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
		GoodsId: "BTC-USDT",
	})
	if err != nil {
		fmt.Println("获取商品详情失败", err)
	}
	fmt.Println("GoodsDetail 最后更新时间", GoodsDetail.GoodsId, GoodsDetail.UpdateStr)

	// ####### 获取榜单数据  #######
	TickerList, err := exchange_api.GetTickerList()
	if err != nil {
		fmt.Println("获取榜单数据失败", err)
	}
	fmt.Println("TickerList 上榜币种数量:", len(TickerList))
}

// K线数据
func KlineFunc() {
	// time := m_time.TimeParse(m_time.LaySP_ss, "2023-05-06 18:56:43")
	// time := m_time.TimeParse(m_time.LaySP_ss, "2024-07-26 16:00:00")

	mo7_start_time := m_time.GetUnixInt64()

	time := m_time.GetUnixInt64()
	klineMap, err := exchange_api.GetKline(global.GetKlineOpt{
		GoodsId: "BTC-USDT",
		Bar:     "1h",
		EndTime: time,
		Limit:   350,
		// Exchange: []string{"okx", "binance"},
		Exchange: []string{"okx"},
	})
	if err != nil {
		fmt.Println("获取K线数据失败", err)
	}

	kline := klineMap["okx"]
	fmt.Println("kline 最新时间", len(kline), kline[len(kline)-1][0])

	m_file.WriteByte(global.Path.DataPath+"/kline-test-h.json", m_json.ToJson(klineMap))

	mo7_end_time := m_time.GetUnixInt64()
	fmt.Println("耗时", mo7_end_time-mo7_start_time)
}

// ########### 交易所K线数据行为一致性检测 ###########
func KlineActionTest() {
	// 早于这个时间，则欧意交易所没数据， 也就是当前时间6年起算
	// var TimeOldest = m_time.TimeParse(m_time.LaySP_ss, "2018-01-11 22:00:00")
	// diff := m_time.GetUnixInt64() - TimeOldest
	// fmt.Println("diff", diff)
	// time := m_time.TimeParse(m_time.LaySP_ss, "2028-01-01 00:00:00")
	time := m_time.TimeParse(m_time.LaySP_ss, "2018-01-11 22:00:00")
	okxKline, err := okx.GetKline(okx.GetKlineOpt{
		Okx_instId: "BTC-USDT",
		Bar:        "1m",
		EndTime:    time,
	})
	if err != nil {
		fmt.Println("okx Err:", err)
	}
	fmt.Println("数据获取成功", len(okxKline), err)
	m_file.WriteByte(global.Path.Okx.Dir+"/kline.json", m_json.ToJson(okxKline))

	binanceKline, err := binance.GetKline(binance.GetKlineOpt{
		Binance_symbol: "BTCUSDT",
		Bar:            "1m",
		EndTime:        time,
	})
	if err != nil {
		fmt.Println("binance Err:", err)
	}
	fmt.Println("数据获取成功", len(binanceKline), err)
	m_file.WriteByte(global.Path.Binance.Dir+"/kline.json", m_json.ToJson(binanceKline))
}
