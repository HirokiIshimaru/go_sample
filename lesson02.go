// Go言語は静的型付け言語
// 変数に入れるデータの種類
// ・数値型
// ・文字列型
// ・ブール型

package main

import (
	"fmt"
	"reflect"
)

func main() {
	// 数値型
	// 整数型 「Int8」「Int16」「Int32」「Int64」がある
	// 数が大きくなるにつれて使える桁数が増える
	// 後ろの数字を省略して「Int」とすることもできる
	// ↑実装するOSやCPUに依存した整数型である
	num01 := 123
	var num02 int = 1234567890

	// 小数点型「float32」「float64」がある
	// 数が大きくなるにつれて使える桁数が増える
	// 通常は「float64」で問題ない
	num03 := 1.23
	var num04 float64 = 1.23456789

	fmt.Println(num01)
	fmt.Println(num02)
	fmt.Println(num03)
	fmt.Println(num04)

	// Goのデータ型は「reflect」のパッケージの「TypeOf」を使用して確認できる
	// fmt.Println(reflect.TypeOf(num01))
	// fmt.Println(reflect.TypeOf(num02))
	// fmt.Println(reflect.TypeOf(num03))
	// fmt.Println(reflect.TypeOf(num04))

	// 文字列型＝String型
	var string_a string = "Hello World!"
	string_b := "Hello World!"
	fmt.Println(string_a)
	fmt.Println(reflect.TypeOf(string_a))
	fmt.Println(string_b)
	fmt.Println(reflect.TypeOf(string_b))

	// ブール型
	// True or Falseの2つのうち、どちらか1つを持つ型
	a := 10
	b := 1
	// var num_bool bool = a > b
	num_bool := a > b

	fmt.Println(num_bool)
	fmt.Println(reflect.TypeOf(num_bool))
}
