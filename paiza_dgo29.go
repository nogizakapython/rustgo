package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！

	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()

	a, _ := strconv.Atoi(sc.Text())

	sc.Scan()

	b, _ := strconv.Atoi(sc.Text())

	fmt.Println(a - b)
}
