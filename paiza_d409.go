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

	n, _ := strconv.Atoi(sc.Text())

	if n == 3 || n == 5 || n == 7 {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}
	//fmt.Println("XXXXXX")
}
