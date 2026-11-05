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
    n := sc.Text()
    m,_ := strconv.Atoi(n)
    ans := ""
    
    
    if m < 10 {
        ans = "0" + n
    } else {
        ans = n
    }
    
    fmt.Println(ans)
}