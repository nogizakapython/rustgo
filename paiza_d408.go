package main
import "fmt"
import "os"
import "strconv"
import "bufio"


func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    n , _ := strconv.Atoi(sc.Text())
    
    sc.Scan()
    o , _ := strconv.Atoi(sc.Text())
    
    fmt.Println(n % o)
}