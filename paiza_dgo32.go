package main
import "fmt"
import "os"
import "bufio"
import "strings"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    name := sc.Text()
    
    array1 := strings.Split(name,"")
    
    n := len(array1)
    
    fmt.Println(array1[n-1])
}