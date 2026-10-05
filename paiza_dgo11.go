package main
import  ( "fmt"
          "bufio"
          "os"
        )  
func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    
    
    for i := 0 ; i < 2 ; i++{
        sc.Scan()
        name := sc.Text()
        fmt.Println(name)
    }
    
}