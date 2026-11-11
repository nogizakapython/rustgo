package main
import "fmt"
import "os"
import "bufio"
import "strconv"


func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！
    
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    num := sc.Text()
    
    ans, _ := strconv.ParseFloat(num,64)
    fmt.Println(ans + 0.1)
}