package binance

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_fetch"
)

var BaseUrlArr = []string{
	"https://api.binance.com",
	"https://api1.binance.com",
	"https://api2.binance.com",
	"https://api3.binance.com",
	"https://api4.binance.com",
	"https://api-gcp.binance.com",
}

func GetGoodsList() {
	m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: BaseUrlArr[2],
		Path:   "/api/v3/exchangeInfo",
		Event: func(s string, data any) {
			fmt.Println("GetGoodsList", s, data)
		},
	}).Get()
}
