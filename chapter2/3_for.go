package chapter2

import f "fmt"

func forSectionHandler() {
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
