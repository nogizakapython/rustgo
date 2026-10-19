package main
import "fmt"
import "os"
import "bufio"
import "strconv"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    n,_ := strconv.Atoi(sc.Text())
    
    for i := 1; i <= n;i++{
        fmt.Println(i)
    }
    //fmt.Println("XXXXXX")
}