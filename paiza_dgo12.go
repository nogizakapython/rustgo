package main
import "fmt"
import "bufio"
import "os"
func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    name := sc.Text()
    fmt.Println(name)
}