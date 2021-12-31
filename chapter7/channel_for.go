package main

import "fmt"

func main() {
	ch1 := make(chan int, 3)

	ch1 <- 1
	ch1 <- 2
	ch1 <- 3
	close(ch1)
	// チャネルのfor文は，値を入れ終わったら，closeするのが基本．

	for i := range ch1 {
		fmt.Println(i)
	}

}
