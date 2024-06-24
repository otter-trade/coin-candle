package exchange_api

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_time"
)

func GetKline(opt global.GetKlineOpt) (resData []global.KlineType, resErr error) {
	resData = nil
	resErr = nil

	if len(opt.GoodsId) < 1 {
		resErr = fmt.Errorf("缺少 GoodsId")
		return
	}

	GoodsDetail, err := GetGoodsDetail(GetGoodsDetailOpt{
		GoodsId: opt.GoodsId,
	})
	if err != nil {
		resErr = err
		return
	}

	if len(opt.Bar) < 1 {
		resErr = fmt.Errorf("缺少 Bar")
		return
	}

	// 当前时间
	now := m_time.GetUnixInt64()
	Before := now
	// 时间 传入的时间戳 必须大于6年前 才有效，否则等价与当前时间戳
	if opt.Before > now-m_time.UnixTimeInt64.Day*2190 {
		Before = opt.Before
	}

	Limit := 100
	if Limit > 1 && Limit < 500 {
		Limit = opt.Limit
	} else {
		resErr = fmt.Errorf("Limit必须为1-500的正整数")
		return
	}

	fmt.Println(Before, Limit, GoodsDetail.GoodsId)
	return
}
