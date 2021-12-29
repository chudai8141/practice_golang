package main

import (
	"fmt"
	"os"
)

func TestDefer() {

	// defer -> 一番最後に実行する．
	// 関数が終了したときに処理を行う
	defer fmt.Println("END")
	fmt.Println("START")
}

func RunDefer() {
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")
	// 一番最後のdeferから処理を行う．
}

func main() {
	TestDefer()

	/*
		defer func() {
			fmt.Println("1")
			fmt.Println("2")
			fmt.Println("3")
		}()

	*/

	RunDefer()

	file, err := os.Create("test.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	file.Write([]byte("Hello"))

}
