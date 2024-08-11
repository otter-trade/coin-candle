package demo

import (
	"fmt"

	"github.com/handy-golang/go-tools/m_file"
	"github.com/handy-golang/go-tools/m_json"
	"github.com/handy-golang/go-tools/m_time"
	"github.com/otter-trade/coin-candle/exchange_api"
	"github.com/otter-trade/coin-candle/global"
)

/*
获取K线数据
可以自行选择获取哪个交易所的数据
*/
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
