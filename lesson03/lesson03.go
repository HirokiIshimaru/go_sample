package main

import "fmt"

func main() {
	// 配列

	//変数名 := [要素の数] データ型 {データ1, データ 2, データ3,...}
	// a := [3]string{"sato", "suzuki", "takahashi"}

	// 要素の数は[...]で省略することができる
	// a := [...]string{"sato", "suzuki", "takahashi"}

	// 配列のデータ変更
	// a[1] = "tanaka"

	// 多次元配列
	a := [2][2]string{{"sato", "suzuki"}, {"takahashi", "tanaka"}}
	fmt.Println(a[0][0])
	fmt.Println(a[0][1])
	fmt.Println(a[1][0])
	fmt.Println(a[1][1])
}
