package main
import "fmt"
import "bufio"
import "os"
import "strconv"


func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    n , _ := strconv.Atoi(sc.Text())
    
    for i := 0 ; i < n ; i++{
        sc.Scan()
        data := sc.Text()
        fmt.Println(data)
    }
    
    //fmt.Println("XXXXXX")
}