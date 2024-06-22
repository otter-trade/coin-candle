package binance

import (
	"coin-candle/global"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
)

func GetTicker() {

	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      "/api/v3/ticker/24hr",
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		global.LogErr("binance.GetTicker Err", err)
	}

	//将请求结果写入目录
	m_file.Write(global.Path.Binance.Dir+"/ticker-original.json", m_json.JsonFormat(resData))

}
