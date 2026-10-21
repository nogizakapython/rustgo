package main

import "fmt"

func main() {
    // 1. 宣言のみ（要素はすべて 0 で初期化される）
    var arr1 [3]int 
    fmt.Println(arr1) // [0 0 0]

    // 2. 宣言と同時に初期化
    var arr2 = [3]int{10, 20, 30}
    // 暗黙的な型定義（:=）を使用する場合
    arr3 := [3]string{"Go", "Python", "Java"}

    // 3. [...] を使うと、初期化子の数から自動で長さを計算してくれる
    arr4 := [...]int{1, 2, 3, 4, 5}
    
    fmt.Println(arr2, arr3, arr4)
}
