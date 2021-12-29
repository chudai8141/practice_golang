package main

import "fmt"

func anything(a interface{}) {
	switch v := a.(type) {
	case string:
		fmt.Println(v + "!?")
	case int:
		fmt.Println(v + 10000)
	}
}

func main() {
	// func anything -> assertion
	anything("aaa")
	anything(1)

	var x interface{} = 3
	// 型アサーション
	i := x.(int)
	fmt.Println(i + 2)
	// ./switch_of_switch.go:16:16: invalid operation: x + 2 (mismatched types interface {} and int)
	// fmt.Println(x + 2)

	f, isFloat64 := x.(float64)
	// 変換に失敗すると2つ目の変数に，falseが返る
	fmt.Println(f, isFloat64)

	if x == nil {
		fmt.Println("None")
	} else if i, isInt := x.(int); isInt {
		fmt.Println(i, "x is Int")
	} else if s, isString := x.(string); isString {
		fmt.Println(s, isString)
	} else {
		fmt.Println("I don't Know")
	}

	// 推奨された型アサーション
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	default:
		fmt.Println("I don't Know")
	}

	switch v := x.(type) {
	case bool:
		fmt.Println(v, "bool")
	case int:
		fmt.Println(v, "int")
	case string:
		fmt.Println(v, "string")
	default:
		fmt.Println("I don't Know")
	}

}
