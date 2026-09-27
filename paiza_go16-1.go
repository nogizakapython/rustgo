// ゼロ・プラス・マイナスを繰り返し判定する
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	count, _ := strconv.Atoi(sc.Text())
	fmt.Println(count)

	for i := 0; i < count; i++ {
		sc.Scan()
		number, _ := strconv.Atoi(sc.Text())
		if number == 0 {
			fmt.Println(strconv.Itoa(number) + "は0")
		} else if number > 0 {
			fmt.Println(strconv.Itoa(number) + "はプラス")
		} else {
			fmt.Println(strconv.Itoa(number) + "はマイナス")
		}

	}
}
