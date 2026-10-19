package main
import  "fmt"
import  "bufio"
import  "os"
import  "strconv"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    
    n,_ := strconv.Atoi(sc.Text())
    
    if n == 7 {
        fmt.Println("Yes")
    } else {
        fmt.Println("No")
    }    
}