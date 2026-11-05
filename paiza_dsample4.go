package main
//package math

import "fmt"
import "os"
import "bufio"
import "strconv"
import "math"



func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    n , _ := strconv.Atoi(sc.Text())
    
    m := float64(n)
    
    ans := math.Pow(m,5)
    fmt.Println(ans)
}