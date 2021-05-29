package main

import "fmt"

func main() {
	// 条件分岐	 if文
	x := 10
	y := 12

	// 簡易文付きif文
	if age := x + y; age >= 20 {
		fmt.Println("adult", age)
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
