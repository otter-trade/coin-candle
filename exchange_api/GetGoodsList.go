package exchange_api

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/exchange_api/okx"
	"coin-candle/global"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
)

func GetGoodsList() {
	binanceGoodsList, err := binance.GetGoodsList()
	if err != nil {
		global.LogErr("错误:exchange_api.GetGoodsList -> binance.GetGoodsList", err)
	}
	m_file.Write(global.Path.Binance.Dir+"/binance_GoodsList.json", m_json.ToStr(binanceGoodsList))

	okx_GoodsList_SPOT, err := okx.GetGoodsList_SPOT()
	if err != nil {
		global.LogErr("错误:exchange_api.GetGoodsList -> okx.GetGoodsList_SPOT", err)
	}
	m_file.Write(global.Path.Okx.Dir+"/okx_GoodsList_SPOT.json", m_json.ToStr(okx_GoodsList_SPOT))

	okx_GoodsList_SWAP, err := okx.GetGoodsList_SWAP()
	if err != nil {
		global.LogErr("错误:exchange_api.GetGoodsList -> okx.GetGoodsList_SPOT", err)
	}
	m_file.Write(global.Path.Okx.Dir+"/okx_GoodsList_SWAP.json", m_json.ToStr(okx_GoodsList_SWAP))
}
