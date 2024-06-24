package exchange_api

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_time"
)

func GetKline(opt global.GetKlineOpt) (resData []global.KlineType, resErr error) {
	resData = nil
	resErr = nil

	// 检查参数 GoodsId
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

	// 检查参数 Bar
	BarObj := global.GetBarOpt(opt.Bar)
	if BarObj.Interval < m_time.UnixTimeInt64.Minute {
		resErr = fmt.Errorf("参数 Bar 错误")
		return
	}

	// Before 缺省值
	now := m_time.GetUnixInt64()
	Before := now
	// 时间 传入的时间戳 必须大于6年前 才有效，否则等价与当前时间戳
	if opt.Before > now-m_time.UnixTimeInt64.Day*2190 {
		Before = opt.Before
	}

	// Limit 缺省值
	Limit := 10
	if opt.Limit > 1 && opt.Limit < 500 {
		Limit = opt.Limit
	} else {
		resErr = fmt.Errorf("Limit必须为1-500的正整数")
		return
	}

	// Exchange 缺省值, 过滤有效值
	var Exchange []string
	if len(opt.Exchange) < 1 {
		Exchange = []string{"okx"}
	} else {
		for _, item := range global.ExchangeOpt {
			for _, exchange := range opt.Exchange {
				if exchange == item {
					Exchange = append(Exchange, item)
					break
				}
			}
		}
	}
	if len(Exchange) < 1 {
		resErr = fmt.Errorf("参数 Exchange 无效")
		return
	}

	// 计算出起止时间  // 结束时间 - 时间间隔 * 条数
	EmdTime := Before
	StartTime := EmdTime - BarObj.Interval*int64(Limit)

	GetKlineFilePath(GetKlineFilePathOpt{
		StartTime: StartTime,
		EmdTime:   EmdTime,
		BarObj:    BarObj,
		Goods:     GoodsDetail,
		Exchange:  Exchange,
	})

	return
}

type GetKlineFilePathOpt struct {
	StartTime int64
	EmdTime   int64
	BarObj    global.KlineBarType
	Goods     global.GoodsType
	Exchange  []string
}

func GetKlineFilePath(opt GetKlineFilePathOpt) {
	for _, exchange := range opt.Exchange {
		fmt.Println(exchange)
	}
}
