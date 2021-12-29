package main

import "fmt"

func main() {
	/*
		i := 0
		for {
			i++
			if i == 3 {
				break
			}
			fmt.Println("Loop")
		}
	*/

	/*
		like python
			point := 0
			for point < 10 {
				fmt.Println(point)
				point++
			}

	*/

	/*
		like clang
		for i := 0; i < 10; i++ {
				if i == 3 {
					continue
				}
				if i == 6 {
					break
				}
				fmt.Println(i)
			}
	*/

	/*
		like c
		arr := [3]int{1, 2, 3}
		for i := 0; i < len(arr); i++ {
			fmt.Println(arr[i])
		}
	*/

	/*
		arr := [3]int{1, 2, 3}
		// i -> index number
		// v -> value for arr
		for i, v := range arr {
			fmt.Println(i, v)
		}

		for i := range arr {
			fmt.Println(i)
		}

	*/

	/*
		sl := []string{"Python", "PHP", "GO"}
		for i, v := range sl {
			fmt.Println(i, v)
		}
	*/

	// k のみ， vのみもできる
	m := map[string]int{"apple": 100, "banana": 200}
	for k, v := range m {
		fmt.Println(k, v)
	}

}
