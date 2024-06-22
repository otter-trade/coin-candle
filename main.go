package main

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/exchange_api/okx"
	"coin-candle/global"
)

func main() {
	// 初始化系统
	global.Start(global.Opt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	// 获交易所交易对信息

	binance.GetGoodsList()

	okx.GetGoodsList()

}
