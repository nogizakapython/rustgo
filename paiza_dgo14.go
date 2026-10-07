package main
import "fmt"
import "os"
import "bufio"
func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    data1 := sc.Text()
    
    sc.Scan()
    data2 := sc.Text()
    
    fmt.Println(data1)
    fmt.Println(data2)
}