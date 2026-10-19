package main
import "fmt"
func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    n := 100
    var i int
    for i = 1; i <= n; i++ {
        if i % 3 == 0 && i % 5 == 0 {
            fmt.Println("FizzBuzz")
        } else if  i % 3 == 0 {
            fmt.Println("Fizz")
        } else if i % 5 == 0 {
            fmt.Println("Buzz")
        } else {
            fmt.Println(i)
        } 
    }
    
}