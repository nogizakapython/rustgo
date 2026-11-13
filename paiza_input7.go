package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 自分の得意な言語で
	// Let's チャレンジ！！

	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()

	name := sc.Text()
	fmt.Println(name)

	sc.Scan()

	name2 := sc.Text()
	fmt.Println(name2)

	//fmt.Println("XXXXXX")
}
