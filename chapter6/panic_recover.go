package main

import "fmt"

func main() {

	// panic -> 実行中のプログラムを強制的に中断させる
	// recover -> panicによって強制終了するプログラムを復旧させる関数
	// 下記のdefer func()の場合，panicによって強制終了する前に，deferに登録されている無名関数を実行する．
	// その後，panic()によりrecover()が呼び出されることから，xにはpanic("runtime error")の"runtime error"が代入される．
	// x != nilではないので，fmt.Println("runtime error")となる．
	defer func() {
		if x := recover(); x != nil {
			fmt.Println(x)
		}
	}()

	panic("runtime error")
	fmt.Println("Start")
}
