package main

import (
	"time"

	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"

	"github.com/handy-golang/go-tools/m_cycle"
)

// ####### 系统初始化基本数据，必须在所有Api调用之前执行完毕 #######
func SysInit() {
	// 初始化系统
	global.SysInit(global.SysInitOpt{
		ProxyURLs: []string{"http://127.0.0.1:10809"},
	})

	// 全局的定时任务
	m_cycle.New(m_cycle.Opt{
		Func: func() {
			// exchange_api.UpdateLocalGoodsList()  // 更新本地的商品列表
			// exchange_api.UpdateLocalTicker()  // 更新本地的榜单
			exchange_api.GetGoodsList() // 读取获取商品列表 并存入内存中
		},
		SleepTime: time.Hour * 4, // 执行一次后 每隔 4h 再执行一次
	}).Start()
}

func main() {
	SysInit() //  系统的初始化
}
