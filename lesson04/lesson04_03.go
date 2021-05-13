package main

import "fmt"

func main() {
	// 論理演算子
	// 複数の条件を判断させる演算子のこと(and && やor || のこと)

	x := 8
	y := 3

	// and &&
	fmt.Println(x >= 5 && x <= 10)
	fmt.Println(y >= 5 && y <= 10)

	// or ||
	fmt.Println(x == 3 || y == 3)
	fmt.Println(x == 1 || y == 1)

}
