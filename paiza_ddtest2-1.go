// 標準入力から1行の整数列を読み取り、合計値から最大値と最小値を1つずつ除いた値を出力する。
// 採点結果から最高点・最低点を除外するような、「外れ値を除いた合計」を求めるプログラム。
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	input := sc.Text()
	values := strings.Split(input, " ")

	// 最大値・最小値は、固定の番兵値ではなく先頭要素で初期化する。
	// NOTE: 入力値の制約が変わっても、初期値が原因で誤判定しないようにするため。
	firstValue, _ := strconv.Atoi(values[0])
	sum := 0
	maxValue := firstValue
	minValue := firstValue

	for i := 0; i < len(values); i++ {
		value, _ := strconv.Atoi(values[i])

		sum += value

		if value > maxValue {
			maxValue = value
		}

		if value < minValue {
			minValue = value
		}
	}

	// 合計から最大値と最小値を1つずつ除外した値を求める。
	answer := sum - maxValue - minValue

	fmt.Println(answer)
}
