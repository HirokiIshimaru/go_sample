// 関数
package main

import "fmt"

func sayHello() {
	fmt.Println("Hello World")
}

// 引数が含まれた関数
func sayGreeting(greeting string) {
	fmt.Println(greeting)
}

// 戻り値がある場合
func cal(x, y int) int {
	return (x / y)
}

func main() {
	sayHello()

	// 引数greetingに文字列「good morning」を渡す
	sayGreeting("Good morning")

	// 戻り値がある場合
	result := cal(6, 3)
	fmt.Println(result)
}
