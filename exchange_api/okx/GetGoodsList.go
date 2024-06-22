package okx

import (
	"coin-candle/global"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
)

var BaseUrlArr = []string{
	"https://www.okx.com",
	"https://aws.okx.com",
}

func GetGoodsList() {
	GetGoodsList_SPOT()
	GetGoodsList_SWAP()
}

func GetGoodsList_SPOT() {
	var DataMap = map[string]any{
		"instType": "SPOT",
	}
	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrlArr[0],
		Path:      "/api/v5/public/instruments",
		Data:      m_json.ToJson(DataMap),
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		global.LogErr("okx.GetGoodsList_SPOT Err", err)
	}

	//将请求结果写入目录
	m_file.Write(global.Path.Okx.Dir+"/goods_list-spot.json", m_json.JsonFormat(resData))
}
func GetGoodsList_SWAP() {
	var DataMap = map[string]any{
		"instType": "SWAP",
	}
	resData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrlArr[0],
		Path:      "/api/v5/public/instruments",
		Data:      m_json.ToJson(DataMap),
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()

	if err != nil {
		global.LogErr("okx.GetGoodsList_SPOT Err", err)
	}

	//将请求结果写入目录
	m_file.Write(global.Path.Okx.Dir+"/goods_list-swap.json", m_json.JsonFormat(resData))
}
