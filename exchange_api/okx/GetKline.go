package okx

import "coin-candle/global"

func GetKline(opt global.GetKlineOpt) {

	if global.CheckGetKlineOpt(opt) != nil {
		return
	}

}
