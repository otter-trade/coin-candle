package exchange_api

import (
	"coin-candle/exchange_api/binance"
	"coin-candle/exchange_api/okx"
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_str"
	"github.com/handy-golang/go-tools/m_time"
	jsoniter "github.com/json-iterator/go"
)

// 读取 GoodsList
func GetGoodsList() (resData []global.GoodsType, resErr error) {

	resData = nil
	resErr = nil

	var fileData = m_file.ReadFile(global.Path.GoodsListFile)
	if len(fileData) < 2 {
		resErr = fmt.Errorf("文件读取失败: %s", global.Path.GoodsListFile)
		return
	}
	var GoodsListData []global.GoodsType
	jsoniter.Unmarshal(fileData, &GoodsListData)
	if len(GoodsListData) < 10 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(GoodsListData))
		return
	}
	resData = GoodsListData

	// 如果没有初始化 global.GoodsMap
	if len(global.GoodsMap) < 10 {
		UpdateGoodsMap(resData)
	}

	return
}

// 更新 GoodsList 至本地
func UpdateLocalGoodsList() {

	var CoinStatusErrTemp = `
${InstID} OkxState:${OkxState} BinanceStatus:${BinanceStatus}`

	binanceGoodsList, err := binance.GetGoodsList()
	if err != nil {
		global.LogErr("错误:exchange_api.GetGoodsList -> binance.GetGoodsList", err)
	}

	// m_file.Write(global.Path.Binance.Dir+"/binance_GoodsList.json", m_json.ToStr(binanceGoodsList))

	okx_GoodsList_SPOT, err := okx.GetGoodsList_SPOT()
	if err != nil {
		global.LogErr("错误:exchange_api.GetGoodsList -> okx.GetGoodsList_SPOT", err)
	}
	// m_file.Write(global.Path.Okx.Dir+"/okx_GoodsList_SPOT.json", m_json.ToStr(okx_GoodsList_SPOT))

	okx_GoodsList_SWAP, err := okx.GetGoodsList_SWAP()
	if err != nil {
		global.LogErr("错误:exchange_api.GetGoodsList -> okx.GetGoodsList_SPOT", err)
	}
	// m_file.Write(global.Path.Okx.Dir+"/okx_GoodsList_SWAP.json", m_json.ToStr(okx_GoodsList_SWAP))

	// 整理现货 ， 基于现货做产品基础 结构 的构建
	var GoodsList = []global.GoodsType{}

	var WarningList = []string{}

	var UpdateTime = m_time.GetTime()
	for _, item := range okx_GoodsList_SPOT {
		if item.QuoteCcy == global.SystemSettleCcy { // 只有 现货商品的 QuoteCcy 才等于 USDT
			Symbol := m_str.Join(item.BaseCcy, global.SystemSettleCcy)
			for _, item2 := range binanceGoodsList {
				if item2.Symbol == Symbol { // 提取两家交易所均有的货币

					var OkxInfo = item
					var BinanceInfo = item2

					var State = "live"

					if OkxInfo.State == "live" && BinanceInfo.Status == "TRADING" {

					} else {
						State = "warning:商品现货状态异常;"
						WarningInfo := m_str.Temp(CoinStatusErrTemp, map[string]string{
							"InstID":        OkxInfo.InstID,
							"OkxState":      OkxInfo.State,
							"BinanceStatus": BinanceInfo.Status,
						})
						WarningList = append(WarningList, WarningInfo)
					}

					var goods global.GoodsType = global.GoodsType{
						GoodsId:       item.InstID,
						State:         State,
						UpdateUnix:    UpdateTime.TimeUnix,
						UpdateStr:     UpdateTime.TimeStr,
						QuoteCcy:      item.QuoteCcy,
						BaseCcy:       item.BaseCcy,
						Okx_SPOT_Info: item,
						BinanceInfo:   item2,
					}
					GoodsList = append(GoodsList, goods)

					break
				}
			}
		}
	}

	// 把合约信息也整理进去
	var NewGoodsList = []global.GoodsType{}
	for _, item := range GoodsList {
		var SWAP_InstId = m_str.Join(item.Okx_SPOT_Info.InstID, "-SWAP")
		var NewItem = item

		// 只有现货状态没问题的，才有资格去整理合约
		if NewItem.State == "live" {
			for _, item2 := range okx_GoodsList_SWAP {
				if SWAP_InstId == item2.InstID {

					if item2.State == "live" {

					} else {
						NewItem.State += "warning:OKX的合约状态异常;" // 增加异常状态
						WarningInfo := m_str.Temp(CoinStatusErrTemp, map[string]string{
							"InstID":         item2.InstID,
							"Okx_SWAP_State": item2.State,
						})
						WarningList = append(WarningList, WarningInfo)
					}
					NewItem.Okx_SWAP_Info = item2
					break
				}
			}
		}

		NewGoodsList = append(NewGoodsList, NewItem)
	}

	// 合约信息标记
	var NewGoodsList2 = []global.GoodsType{}
	for _, item := range NewGoodsList {
		var NewItem2 = item
		if len(NewItem2.Okx_SWAP_Info.InstID) < 2 {
			NewItem2.State += "warning:没有OKX合约信息"
		}
		NewGoodsList2 = append(NewGoodsList2, NewItem2)
	}

	// 打印警告
	if len(WarningList) > 0 {
		global.ExchangeLog.Println(m_str.Join("警告:以下商品状态存在问题\n", WarningList))
	}

	// 如果数量不对则发出警告
	if len(NewGoodsList2) > 10 {
		UpdateGoodsMap(NewGoodsList2) // 更新 global.GoodsMap
		m_file.Write(global.Path.GoodsListFile, m_json.ToStr(NewGoodsList2))
		global.RunLog.Println("商品列表更新完成", global.Path.GoodsListFile)
	} else {
		global.LogErr("exchange_api.GetGoodsList 数量不足", len(NewGoodsList2))
	}
}

// 将数据更新至内存中
func UpdateGoodsMap(GoodsList []global.GoodsType) {
	global.GoodsMap = make(map[string]global.GoodsType)
	for _, item := range GoodsList {
		global.GoodsMap[item.GoodsId] = item
	}
}

// 读取商品详情
type GetGoodsDetailOpt struct {
	GoodsId        string // 三个值任选其一
	Okx_InstID     string
	Binance_Symbol string
}

func GetGoodsDetail(opt GetGoodsDetailOpt) (resData global.GoodsType, resErr error) {

	resData = global.GoodsType{}
	resErr = nil

	if len(global.GoodsMap) < 10 {
		resErr = fmt.Errorf("没有初始化 global.GoodsMap")
		return
	}

	// 如果  opt.GoodsId
	if len(opt.GoodsId) > 1 {
		resData = global.GoodsMap[opt.GoodsId]
	}

	// 如果  opt.Okx_InstID
	if len(opt.Okx_InstID) > 1 {
		for _, item := range global.GoodsMap {
			if item.Okx_SPOT_Info.InstID == opt.Okx_InstID {
				resData = item
				break
			}
		}
	}

	// 如果  opt.Binance_Symbol
	if len(opt.Binance_Symbol) > 1 {
		for _, item := range global.GoodsMap {
			if item.BinanceInfo.Symbol == opt.Binance_Symbol {
				resData = item
				break
			}
		}
	}

	if len(resData.GoodsId) < 2 {
		resErr = fmt.Errorf("没有找到商品")
		return
	}

	return
}
