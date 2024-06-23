package exchange_api

import (
	"coin-candle/global"
)

func GetKline(opt global.GetKlineOpt) (resData []global.KlineType, resErr error) {

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
	return
}
