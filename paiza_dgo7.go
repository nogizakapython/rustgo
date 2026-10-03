package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	var a int
	var b int
	a = 437326
	b = 9085
	ans1 := a / b
	ans2 := a % b
	fmt.Println(strconv.Itoa(ans1) + " " + strconv.Itoa(ans2))
}
