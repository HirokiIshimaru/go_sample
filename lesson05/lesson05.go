package main

import "fmt"

func main() {
	// 条件分岐	 if文
	x := 10
	y := 12

	if age := x + y; age >= 20 {
		fmt.Println("adult")
	} else if age == 0 {
		fmt.Println("baby")
	} else {
		fmt.Println("child")
	}
}
