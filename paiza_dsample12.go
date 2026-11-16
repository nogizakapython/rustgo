package main
import "fmt"
import "os"
import "bufio"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    n := 2 
    sc := bufio.NewScanner(os.Stdin)
    for i := 0 ; i < n ; i++ {
        sc.Scan()
        s := sc.Text()
        fmt.Println(s)
    }
    //fmt.Println("XXXXXX")
}