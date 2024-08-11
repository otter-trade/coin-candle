package my_test

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/exchange_api/binance"
	"github.com/otter-trade/coin-candle/exchange_api/okx"
	"github.com/otter-trade/coin-candle/global"
)

// ########### 交易所K线数据行为一致性检测 ###########
func KlineActionTest() {
	// 早于这个时间，则欧意交易所没数据， 也就是当前时间6年起算
	// var TimeOldest = m_time.TimeParse(m_time.LaySP_ss, "2018-01-11 22:00:00")
	// diff := m_time.GetUnixInt64() - TimeOldest
	// fmt.Println("diff", diff)
	// time := m_time.TimeParse(m_time.LaySP_ss, "2028-01-01 00:00:00")
	time := m_time.TimeParse(m_time.LaySP_ss, "2018-01-11 22:00:00")
	okxKline, err := okx.GetKline(okx.GetKlineOpt{
		Okx_instId: "BTC-USDT",
		Bar:        "1m",
		EndTime:    time,
	})
	if err != nil {
		fmt.Println("okx Err:", err)
	}
	fmt.Println("数据获取成功", len(okxKline), err)
	m_file.WriteByte(global.Path.Okx.Dir+"/kline.json", m_json.ToJson(okxKline))

	binanceKline, err := binance.GetKline(binance.GetKlineOpt{
		Binance_symbol: "BTCUSDT",
		Bar:            "1m",
		EndTime:        time,
	})
	if err != nil {
		fmt.Println("binance Err:", err)
	}
	fmt.Println("数据获取成功", len(binanceKline), err)
	m_file.WriteByte(global.Path.Binance.Dir+"/kline.json", m_json.ToJson(binanceKline))
}

// K线数据
func KlineFunc() {
	// time := m_time.TimeParse(m_time.LaySP_ss, "2023-05-06 18:56:43")
	// time := m_time.TimeParse(m_time.LaySP_ss, "2024-07-26 16:00:00")

	mo7_start_time := m_time.GetUnixInt64()

	time := m_time.GetUnixInt64()
	klineMap, err := exchange_api.GetKline(global.GetKlineOpt{
		GoodsId: "BTC-USDT",
		Bar:     "1h",
		EndTime: time,
		Limit:   350,
		// Exchange: []string{"okx", "binance"},
		Exchange: []string{"okx"},
	})
	if err != nil {
		fmt.Println("获取K线数据失败", err)
	}

	kline := klineMap["okx"]
	fmt.Println("kline 最新时间", len(kline), kline[len(kline)-1][0])

	m_file.WriteByte(global.Path.DataPath+"/kline-test-h.json", m_json.ToJson(klineMap))

	mo7_end_time := m_time.GetUnixInt64()
	fmt.Println("耗时", mo7_end_time-mo7_start_time)
}
