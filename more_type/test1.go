package main

import "fmt"

func main() {
	a := 15

	// ポインタを取得
	b := &a
	fmt.Println(*b) // 15

	// ポインタの中身を書き換え
	*b = 21
	fmt.Println(a) // 21
}
