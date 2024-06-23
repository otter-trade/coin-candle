package main

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/exchange_api/okx"
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_time"
)

func main() {
	// 初始化系统
	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	// 获交易所交易对信息

	// binance.GetGoodsList()

	// binance.GetTicker()

	// okx.GetGoodsList()

	// okx.GetTicker()

	_, errOkx := okx.GetKline(okx.GetKlineOpt{
		Okx_instId: "BTC-USDT",
		Bar:        "1m",
		Before:     m_time.GetUnixInt64() - m_time.UnixTimeInt64.Day*1, // 一天前
	})

	fmt.Println("errOkx", errOkx)

	_, errBinance := binance.GetKline(binance.GetKlineOpt{
		Binance_symbol: "BTCUSDT",
		Bar:            "1m",
		Before:         m_time.GetUnixInt64() - m_time.UnixTimeInt64.Day*1, // 一天前
	})

	fmt.Println("errBinance", errBinance)

}
