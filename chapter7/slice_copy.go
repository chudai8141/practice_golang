package main

import "fmt"

func main() {
	/*
		sl := []int{100, 200}
		sl2 := sl

		sl2[0] = 1000

		// sl[0] と　sl2[0]が参照するメモリのアドレスが同じため，sl2[0]に1000を代入すると，slにも代入した値が反映される
		fmt.Println(sl)

		var i int = 10
		i2 := i
		i2 = 100
		fmt.Println(i, i2)

	*/

	sl := []int{1, 2, 3, 4, 5}
	sl2 := make([]int, 5, 10)
	fmt.Println(sl2)

	n := copy(sl2, sl)

	fmt.Println(n, sl2)

}
