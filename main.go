package main

import "github.com/handy-golang/go-tools/m_str"

func main() {

	a := []rune("Hello, World!")
	str := m_str.ToStr(a)
	println(str)
}
