package binance

import (
	"coin-candle/global"
	"fmt"
	"strings"

	"github.com/handy-golang/go-tools/m_fetch"
	"github.com/handy-golang/go-tools/m_json"
	jsoniter "github.com/json-iterator/go"
)

func GetTicker() (resData []global.BinanceTickerType, resErr error) {

	resData = nil
	resErr = nil

	fetchData, err := m_fetch.NewHttp(m_fetch.HttpOpt{
		Origin:    BaseUrl,
		Path:      "/api/v3/ticker/24hr",
		ProxyURLs: global.Path.ProxyURLs,
	}).Get()
	if err != nil {
		resErr = err
		return
	}

	var result []global.BinanceTickerType
	jsoniter.Unmarshal(fetchData, &result)

	if len(result) < 2 {
		resErr = fmt.Errorf("错误:结果返回不正确 %+v", m_json.ToStr(result))
		return
	}
	resData = result
	SetInstID(resData)

	return
}

func SetInstID(data []global.BinanceTickerType) (TickerList []global.BinanceTickerType) {

	// var list []global.BinanceTickerType
	for _, val := range data {
		find := strings.Contains(val.Symbol, global.SystemSettleCcy)
		if find {
			InstID := strings.Replace(val.Symbol, global.SystemSettleCcy, "-"+global.SystemSettleCcy, 1)

			fmt.Println(111, InstID)
			// SPOT := okxInfo.Inst[InstID]
			// val.InstID = SPOT.InstID
			// if len(SPOT.Symbol) > 3 {
			// 	list = append(list, val)
			// }
		}
	}

	return
}
