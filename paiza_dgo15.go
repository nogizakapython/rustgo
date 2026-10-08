package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	data := sc.Text()
	array1 := strings.Split(data, " ")

	A, _ := strconv.Atoi(array1[0])
	B, _ := strconv.Atoi(array1[1])
	ans1 := A - B
	ans2 := A * B

	fmt.Println(strconv.Itoa(ans1) + " " + strconv.Itoa(ans2))
}
