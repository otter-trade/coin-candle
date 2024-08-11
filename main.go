package main

import (
	"time"

	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"

	"github.com/handy-golang/go-tools/m_cycle"
)

func main() {
	SysInit() //  系统的初始化
	// MarketFunc()    // 市场相关的数据
	// KlineFunc()     // K线相关的数据
	// MockServeFunc() // MockServe 的增删查
	// PositionFunc() // 持仓管理
}

// ####### 系统初始化，必须在所有函数之前执行完毕 #######
func SysInit() {
	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	m_cycle.New(m_cycle.Opt{
		Func: func() {
			// // 更新本地的商品列表
			// exchange_api.UpdateLocalGoodsList()
			// // 更新本地的榜单
			// exchange_api.UpdateLocalTicker()
			// 读取获取商品列表 并存入内存中
			exchange_api.GetGoodsList()
		},
		SleepTime: time.Hour * 4, // 执行一次后 每隔 4h 再执行一次
	}).Start()
}
