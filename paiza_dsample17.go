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
    d := sc.Text()
    array1 := strings.Split(d," ")
    a , _ := strconv.ParseFloat(array1[0],64)
    b , _ := strconv.ParseFloat(array1[1],64)
    c , _ := strconv.ParseFloat(array1[2],64)
    
    fmt.Println(a + 0.1)
    fmt.Println(b + 0.1)
    fmt.Println(c + 0.1)
}