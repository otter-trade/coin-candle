package binance

import (
	"coin-candle/global"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
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
	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrlArr[2],
		Path:      "/api/v3/exchangeInfo",
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		global.LogErr("binance.GetGoodsList Err", err)
	}

	//将请求结果写入目录
	m_file.Write(global.Path.Binance.Dir+"/goods_list-original.json", m_json.JsonFormat(resData))
}
