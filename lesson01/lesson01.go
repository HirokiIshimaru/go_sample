// 変数の宣言
package main

import "fmt"

func main() { //同一フォルダ内にfunc main()を使用した別の「.go」ファイルがあるとエラーがでる
	// var num int      //varで変数を宣言・変数名をつける・データの種類を入れる(この場合数値を入れるのでintを入れる)
	// num = 1          //変数numに1を代入

	// var num = 1      //変数の宣言と同時に代入することができる
	// num := 1         //より簡単に変数宣言する方法
	// fmt.Println(num) //変数の中を参照するには、Println(プリントライン)関数を使う

	num := 1
	num01 := 2
	num_01 := 3

	fmt.Println(num)
	fmt.Println(num01)
	fmt.Println(num_01)
}
