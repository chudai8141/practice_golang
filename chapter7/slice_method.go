package main

import "fmt"

func main() {

	sl := []int{100, 200}
	fmt.Println(sl)

	sl = append(sl, 300)
	fmt.Println(sl)

	sl = append(sl, 400, 500, 600)
	fmt.Println(sl)

	sl2 := make([]int, 5)
	fmt.Println(sl2)

	fmt.Println(len(sl2))

	fmt.Println(cap(sl2))

	// 第二引数は要素数，第三引数は容量．
	// capはプログラムのメモリを気にする場合は，使用する．
	sl3 := make([]int, 5, 10)

	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))

	sl3 = append(sl3, 1, 2, 3, 4, 5, 6, 7)

	fmt.Println(len(sl3))

	fmt.Println(cap(sl3))

	// ↑のoutputs
	// 12 要素数
	// 20 メモリ数
}
