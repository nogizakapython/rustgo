package main
import "fmt"
import "os"
import "bufio"
func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    a := ""
    b := ""
    
    sc.Scan()
    a = sc.Text()
    
    sc.Scan()
    b = sc.Text()
    
    if a == b{
        fmt.Println("OK")
    } else {
        fmt.Println("NG")
    }    
}