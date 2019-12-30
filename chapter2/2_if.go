package chapter2

import f "fmt"

func ifSectionHandler() {
	// if 構文
	// 三項演算子は使えない
	// {} の省略記法はコンパイルエラーになる
	a, b := 10, 100
	if a > b {
		f.Println("a は b よりも大きい")
	} else if a < b {
		f.Println("b は a よりも大きい")
	} else {
		f.Println("a と b は等しい")
	}
}
