// method
package main

import "fmt"

type Student struct {
	name string
	// math, english float64
}

// methodの書き方
func (s Student) avg(math, english float64) float64 {
	// 数学と英語の平均点を求めるmethod
	// fmt.Println((math + english) / 2)

	// methodで戻り値を使う場合
	return (math + english) / 2
}

func main() {
	// a001 := Student{"sato", 80, 70}

	// methodに引数を渡す場合
	// a001 := Student{"sato"}
	// a001.avg(80, 70)

	a001 := Student{"satp"}
	fmt.Println(a001.avg(80, 70))
}
