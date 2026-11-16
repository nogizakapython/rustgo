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
    
    d := sc.Text()
    array1 := strings.Split(d," ")
    for _ ,  v := range array1{
        fmt.Println(v)
    }
    //fmt.Println("XXXXXX")
}