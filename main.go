package main

import (
	"coin-candle/global"
	"fmt"

	"github.com/handy-golang/go-tools/m_json"
)

func main() {

	global.Start(global.Opt{})

	fmt.Println("global.Dir", m_json.Format(global.Dir))

}
