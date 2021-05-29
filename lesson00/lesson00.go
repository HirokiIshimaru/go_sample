package main //mainパッケージに必ず属す必要がある
import (
	"fmt"
)

func main() { //同一フォルダ内にfunc main()を使用した別の「.go」ファイルがあるとエラーがでる
	fmt.Println("おはよう")
	fmt.Println("こんにちは")
	fmt.Println("こんばんわ")
}

// goファイルをターミナルで動かす方法
// go run ファイル名
