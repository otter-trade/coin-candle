package main

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/global"
)

func main() {
	// 初始化系统
	global.Start(global.Opt{})

	// 获交易所交易对信息

	binance.GetGoodsList()

}
