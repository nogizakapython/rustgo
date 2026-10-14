package main
import "fmt"
import "bufio"
import "os"
import "strconv"


func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    point, _ := strconv.Atoi(sc.Text())
    
    if point <= 100 {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
    
    //fmt.Println("XXXXXX")
}