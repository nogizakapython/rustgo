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
    
    data := sc.Text()
    array1 := strings.Split(data," ")
    name := array1[1]
    fmt.Println("Hello " + name + " !" )
}