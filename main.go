package main

import (
	"fmt"
	"time"

	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/exchange_api/binance"
	"github.com/otter-trade/coin-candle/exchange_api/okx"
	"github.com/otter-trade/coin-candle/global"

	"github.com/handy-golang/go-tools/m_cycle"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_time"
)

func main() {
	Start()
	KlineActionTest()
}

func Start() {

	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	m_cycle.New(m_cycle.Opt{
		Func: func() {
			// 更新本地的商品列表
			exchange_api.UpdateLocalGoodsList()
			// 更新本地的榜单
			exchange_api.UpdateLocalTicker()
			// 读取本地并缓存
			exchange_api.GetGoodsList()
		},
		SleepTime: time.Hour * 4, // 每4小时执行一次更新
	}).Start()

	// // ####### 获取商品列表 #######
	// GoodsList, err := exchange_api.GetGoodsList()
	// if err != nil {
	// 	fmt.Println("获取商品列表失败", err)
	// }
	// fmt.Println("GoodsList 商品总数:", len(GoodsList))

	// // ####### 获取商品详情 #######
	// GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
	// 	GoodsId: "BTC-USDT",
	// })
	// if err != nil {
	// 	fmt.Println("获取商品详情失败", err)
	// }
	// fmt.Println("GoodsDetail 最后更新时间", GoodsDetail.GoodsId, GoodsDetail.UpdateStr)

	// // ####### 获取榜单数据  #######
	// TickerList, err := exchange_api.GetTickerList()
	// if err != nil {
	// 	fmt.Println("获取榜单数据失败", err)
	// }
	// fmt.Println("TickerList 上榜币种数量:", len(TickerList))

	// ####### K线数据 #######
	// time := m_time.TimeParse(m_time.LaySP_ss, "2023-05-06 18:56:43")
	// time := m_time.TimeParse(m_time.LaySP_ss, "2024-05-21 18:55:43")
	time := m_time.GetUnixInt64()
	klineMap, err := exchange_api.GetKline(global.GetKlineOpt{
		GoodsId: "BTC-USDT",
		Bar:     "1m",
		EndTime: time, // 一年前
		Limit:   5,
		// Exchange: []string{"okx", "binance"},
		Exchange: []string{"okx"},
	})

	if err != nil {
		fmt.Println("获取K线数据失败", err)
	}
	// fmt.Println("kline 总长度", len(klineMap["binance"]))

	m_file.WriteByte(global.Path.DataPath+"/kline-test1.json", m_json.ToJson(klineMap))

}

func KlineActionTest() {
	// ########### 交易所K线数据行为一致性检测 ###########

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
