package main

import "fmt"

// 構造体を定義
type Student struct {
	// フィールドnameを定義
	name string
}

// 平均点を計算するメソッド
func (s Student) calAvg(data []float64) (avgResult float64) {
	sum := 0.0
	for i := 0; i < len(data); i++ {
		sum += data[i]
	}
	avgResult = sum / float64(len(data))
	return
}

// 結果を判定するメソッド
func (s Student) judge(avg float64) (judgeResult string) {
	if avg >= 60 {
		judgeResult = "passed"
	} else {
		judgeResult = "failed"
	}
	return
}

func main() {
	a001 := Student{"sato"}
	data := []float64{70, 65, 60, 50, 10, 30}
	var avg float64 = a001.calAvg(data)
	result := a001.judge(avg)
	fmt.Println(avg)
	fmt.Println(a001.name + "" + result)
}
