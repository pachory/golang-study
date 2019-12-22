package chapter2

import (
	// パッケージを利用しない
	f "fmt"
	// パッケージを利用しない
	_ "golang-study/gosample"
	// パッケージ名の省略
	. "strings"
)

// 変数宣言
var message string = "hello world"

// 一度に複数の変数を宣言する
// パターン1
var foo, bar, baz string = "foo", "bar", "baz"

// パターン2
var (
	a string = "aaa"
	b        = "bbb"
	c        = "ccc"
	d        = "ddd"
)

func Exec() {
	// strings パッケージの import 時に省略記法をとってるので、 strings.ToUpper() と記述しなくてよい
	f.Println(ToUpper(message))
	f.Println(foo, bar, baz)
	f.Println(a, b, c, d)

	// 変数内部では型の指定を省略し、型推論を用いた変数宣言が可能
	// 関数の外では無理っぽい
	hoge := "hoge"
	f.Println(hoge)

	// 定数
	const hello string = "hello"
	f.Println(hello)

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

	// for 構文
	// while や do/while は存在しない、全て for を利用する
	// 無限ループは for {...} で実行できる
	for i := 0; i < 5; i++ {
		f.Println(i)
	}

	// ループにおける break, continue
	n := 0
	for {
		n++
		if n > 10 {
			// ループを抜ける
			break
		}

		if n%2 == 0 {
			// 偶数の場合は次のループ処理に進む
			continue
		}

		// 奇数のみ表示
		f.Println(n)
	}
}
