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

func importSectionHandler() {
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
}
