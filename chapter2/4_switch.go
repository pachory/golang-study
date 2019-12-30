package chapter2

import f "fmt"

func switchSectionHandler() {
	// switch 構文
	n := 10
	switch n {
	case 15:
		f.Println("FizzBuzz")
	case 5, 10:
		f.Println("Buzz")
	case 3, 6, 9:
		f.Println("Fizz")
	default:
		f.Println(n)
	}

	// Go の Switch は case に break を記述しなくても次の case に進まない
	// 次の case に進めたい場合は fallthrough を記述する
	n = 3
	switch n {
	case 3:
		n = n - 1
		fallthrough
	case 2:
		n = n - 1
		fallthrough
	case 1:
		n = n - 1
		f.Println(n)
	}

	// 評価式でも switch を利用できる
	n = 10
	switch {
	case n%15 == 0:
		f.Println("FizzBuzz")
	case n%5 == 0:
		f.Println("Buzz")
	case n%3 == 0:
		f.Println("Fizz")
	default:
		f.Println(n)
	}
}
