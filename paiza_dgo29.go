package main
import "fmt"
import "strconv"
import "os"
import "bufio"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    n , _ := strconv.Atoi(sc.Text())
    
    ans := ""
    
    for i := 0 ; i < n ; i++ {
        ans = ans + "*"
    }
    fmt.Println(ans)
}