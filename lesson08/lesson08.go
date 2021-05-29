// 構造体とは
// 異なるデータ型の変数を1つにまとめたもの
// 構造体に定義した変数を「フィールド」という

package main

import "fmt"

// 構造体名はStudentとして、生徒の名前を格納するnameというフィールドと、
// 数学と英語の点数を格納するmathとenglishというフィールドを定義する

type Student struct {
	name          string
	math, english float64
}

func main() {
	// s := Student{"sato", 80, 70}
	// データを代入する書き方
	s := Student{name: "sato", math: 80, english: 70}
	fmt.Println(s)
}
