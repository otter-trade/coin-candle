package main

import (
	"coin-candle/exchange_api"
	"coin-candle/global"
)

func main() {
	// 初始化系统
	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	// 更新本地的商品列表
	exchange_api.UpdateLocalGoodsList()
	// 更新本地的榜单
	exchange_api.UpdateLocalTicker()

}
