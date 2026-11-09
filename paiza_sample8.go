package main
import "fmt"
import "os"
import "strconv"
import "bufio"
import "strings"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    data := sc.Text()
    array1 :=  strings.Split(data," ")
    m , _ := strconv.Atoi(array1[0])
    n , _ := strconv.Atoi(array1[1])
    fmt.Println(m % n)
}