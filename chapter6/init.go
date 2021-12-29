package main

import "fmt"

// init -> initialize 初期化メソッド
func init() {
	fmt.Println("init")
}

// 出現順
func init() {
	fmt.Println("init2")
}

func main() {
	fmt.Println("Main")
}
