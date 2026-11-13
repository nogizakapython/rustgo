package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()

	d := sc.Text()

	array1 := strings.Split(d, " ")
	l := len(array1)

	for i := 0; i < l; i++ {
		fmt.Println(array1[i])
	}

}
