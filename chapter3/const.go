package main

import "fmt"

// 定数
// const -> globalにするのが一般的

// 頭文字を大文字にすると，別パッケージから呼び出せることができる．
const PI = 3.14

const (
	URL      = "http://xxx.co.jp"
	Sitename = "test"
)

//
const (
	A = 1
	B
	C
	D = "A"
	E
	F
)

const (
	c0 = iota
	c1
	c2
)

// var Big int = 9223372036854775807 + 1
const Big = 9223372036854775807 + 1

func main() {
	fmt.Println(PI)

	// PI = 3
	// fmt.Println(PI)

	fmt.Println(URL, Sitename)

	fmt.Println(A, B, C, D, E, F)

	fmt.Println(Big - 1)

	fmt.Println(c0, c1, c2)

}
