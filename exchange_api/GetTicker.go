package exchange_api

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/global"
	"fmt"
)

func UpdateLocalTicker() {

	binanceTicker, err := binance.GetTicker()
	if err != nil {
		global.LogErr("错误:exchange_api.GetTicker -> binance.GetTicker", err)
	}

	fmt.Println("xxx", len(binanceTicker))
	// okxTicker, err := okx.GetTicker()
	// if err != nil {
	// 	global.LogErr("错误:exchange_api.GetTicker -> okx.GetTicker", err)
	// }

	// if len(binanceTicker) < 6 || len(okxTicker) < 6 {
	// 	global.LogErr("exchange_api.GetTicker 数据条目不正确", len(binanceTicker), len(okxTicker))
	// 	return
	// }

	// var tickerList = []global.TickerType{}

	// for _, okx := range okxTicker {
	// 	for _, binance := range binanceTicker {
	// 		if okx.InstID == binance.InstID {
	// 			ticker := TickerCount(okx, binance)
	// 			if len(ticker.InstID) > 2 {
	// 				tickerList = append(tickerList, ticker)
	// 			}
	// 			break
	// 		}
	// 	}
	// }

}
