package okx

import (
	"coin-candle/global"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
)

func GetGoodsList() {
	GetGoodsList_SPOT()
	GetGoodsList_SWAP()
}

func GetGoodsList_SPOT() {
	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: BaseUrl,
		Path:   "/api/v5/public/instruments",
		DataMap: map[string]any{
			"instType": "SPOT",
		},
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		global.LogErr("okx.GetGoodsList_SPOT Err", err)
	}

	//将请求结果写入目录
	m_file.Write(global.Path.Okx.Dir+"/goods_list-spot.json", m_json.JsonFormat(resData))
}
func GetGoodsList_SWAP() {
	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin: BaseUrl,
		Path:   "/api/v5/public/instruments",
		DataMap: map[string]any{
			"instType": "SWAP",
		},
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		global.LogErr("okx.GetGoodsList_SPOT Err", err)
	}

	//将请求结果写入目录
	m_file.Write(global.Path.Okx.Dir+"/goods_list-swap.json", m_json.JsonFormat(resData))
}
