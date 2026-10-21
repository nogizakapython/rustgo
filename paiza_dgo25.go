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
	a, _ := strconv.Atoi(array1[0])
	b, _ := strconv.Atoi(array1[1])
	fmt.Println(a + b)
}
