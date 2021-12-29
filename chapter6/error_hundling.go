package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string = "100"
	// var s string = "A"

	/*
		s に数字の文字列が代入されれば，errにエラー文は入らない．
		s に数字以外の文字列が代入されると，errにエラー文が代入されて，if文の処理が動作する．
		エラーハンドリング処理の一連．
		Atoiの場合...
	*/
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("i = %T\n", i)

}
