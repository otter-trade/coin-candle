package exchange_api

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/global"
	"fmt"
)

func UpdateLocalTicker() {
	GetBinanceTicker()
}

func GetBinanceTicker() {
	binanceTicker, err := binance.GetTicker()
	if err != nil {
		global.LogErr("错误:exchange_api.GetTicker -> binance.GetTicker", err)
	}

	for _, item := range binanceTicker {
		fmt.Println(item.Symbol)
	}
}
