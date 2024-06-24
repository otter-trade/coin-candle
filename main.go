package main

import (
	"coin-candle/exchange_api"
	"coin-candle/global"
	"fmt"
	"time"

	"github.com/handy-golang/go-tools/m_cycle"
	"github.com/handy-golang/go-tools/m_time"
)

func main() {
	Start()
	Api()
}

func Start() {

	// 初始化系统的各项参数
	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})
	m_cycle.New(m_cycle.Opt{
		Func: func() {
			// 更新本地的商品列表
			exchange_api.UpdateLocalGoodsList()
			// 更新本地的榜单
			exchange_api.UpdateLocalTicker()
		},
		SleepTime: time.Hour * 4, // 每4小时执行一次更新
	}).Start()

	// 获取榜单数据
	// Ticker, err := exchange_api.GetTickerList()
	// if err != nil {
	// 	fmt.Println("获取榜单数据失败", err)
	// }
	// fmt.Println("Ticker", Ticker)

	// // 获取 商品列表
	// GoodsList, err := exchange_api.GetTickerList()
	// if err != nil {
	// 	fmt.Println("获取商品列表失败", err)
	// }
	// fmt.Println("GoodsList", GoodsList)

	// // 获取商品详情
	// GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
	// 	GoodsId: "BTC-USDT",
	// })
	// if err != nil {
	// 	fmt.Println("获取商品详情失败", err)
	// }
	// fmt.Println("GoodsDetail", GoodsDetail)

}

func Api() {

	time := m_time.TimeParse(m_time.LaySP_ss, "2023-05-06 18:56:43")

	kline, err := exchange_api.GetKline(global.GetKlineOpt{
		GoodsId:  "BTC-USDT",
		Bar:      "1m",
		EndTime:  time, // 一年前
		Limit:    382,
		Exchange: []string{"okx", "binance"},
	})

	fmt.Println("kline", kline, err)

}
