package binance

import (
	"coin-candle/global"
	"os"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
)

var BaseUrlArr = []string{
	"https://api.binance.com",
	"https://api1.binance.com",
	"https://api2.binance.com",
	"https://api3.binance.com",
	"https://api4.binance.com",
	"https://api-gcp.binance.com",
}

var CacheFilePath_GoodsList = m_str.Join(
	global.Dir.DataPath,
	m_str.ToStr(os.PathSeparator),
	"binance-exchangeInfo",
)

func GetGoodsList() {
	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrlArr[2],
		Path:      "/api/v3/exchangeInfo",
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	}).Get()

	if err != nil {
		global.LogErr("binance.GetGoodsList Err", err)
	}
	m_file.Write(global.Dir.DataPath+"/exchangeInfo.json", m_json.JsonFormat(resData))
}
