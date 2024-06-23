package main

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/global"
)

func main() {
	// 初始化系统
	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	// 获交易所交易对信息

	// binance.GetGoodsList()

	binance.GetTicker()

	// okx.GetGoodsList()

	// okx.GetTicker()

	// okxKline, errOkx := okx.GetKline(okx.GetKlineOpt{
	// 	Okx_instId: "BTC-USDT",
	// 	Bar:        "1m",
	// 	// Before:     m_time.GetUnixInt64() - m_time.UnixTimeInt64.Day*365, // 一年前
	// })

	// fmt.Println("Okx", okxKline, errOkx)

	// binanceKline, errBinance := binance.GetKline(binance.GetKlineOpt{
	// 	Binance_symbol: "BTCUSDT",
	// 	Bar:            "1m",
	// 	// Before:         m_time.GetUnixInt64() - m_time.UnixTimeInt64.Day*365, // 一年前
	// })

	// fmt.Println("binance", binanceKline, errBinance)

}
