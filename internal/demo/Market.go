package demo

import (
	"fmt"

	"github.com/otter-trade/coin-candle/exchange_api"
)

/*
#### 获取商品列表 ####
*/
func GetGoodsList_demo() {
	GoodsList, err := exchange_api.GetGoodsList()
	if err != nil {
		fmt.Println("获取商品列表失败", err)
	}
	fmt.Println("GoodsList 商品总数:", len(GoodsList))
}

/*
#### 获取商品详情 ####
*/
func GetGoodsDetail_demo() {
	GoodsDetail, err := exchange_api.GetGoodsDetail(exchange_api.GetGoodsDetailOpt{
		GoodsId: "BTC-USDT",
	})
	if err != nil {
		fmt.Println("获取商品详情失败", err)
	}
	fmt.Println("GoodsDetail 最后更新时间", GoodsDetail.GoodsId, GoodsDetail.UpdateStr)
}

/*
#### 获取榜单数据 ####
综合了两家交易所的可交易币种交易量排名几的币种榜单，包括起 24h 成交量等
*/
func GetTickerList_demo() {
	TickerList, err := exchange_api.GetTickerList()
	if err != nil {
		fmt.Println("获取榜单数据失败", err)
	}
	fmt.Println("TickerList 上榜币种数量:", len(TickerList))
}
