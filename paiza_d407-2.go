package main
import "fmt"
import "os"
import "bufio"
func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    
    sc := bufio.NewScanner(os.Stdin)
    
    
    for i := 0;i < 3;i++{
        sc.Scan()
        name := sc.Text()
        fmt.Println(name)
    }
    
}