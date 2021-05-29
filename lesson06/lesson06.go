package main

import "fmt"

func main() {

	// for構文
	for i := 0; i <= 4; i++ {
		fmt.Println(i)
	}

	for i := 0; i <= 4; i++ {
		if i == 2 {
			break
			// break文	繰り返し処理を終了させる
		}
		if i == 3 {
			continue
			// continue文	処理をスキップする
		}
		fmt.Println(i)
	}

	// for文のネスト
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 2; j++ {
			fmt.Println(i, "-", j)
		}
	}

	// for文を使って配列の中身を表示
	arr := [...]int{2, 4, 6, 8, 10}
	for i := 0; i <= 4; i++ {
		fmt.Println(arr[i])
	}

	// for文をシンプルに書く方法
	i := 0 //スタート値、条件、増減式を書かなくてもfor文は動く
	for i <= 4 {
		fmt.Println(i)
		i++
	}
}
