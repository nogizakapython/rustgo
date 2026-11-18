package main
import "fmt"
import "strings"
import "strconv"
import "os"
import "bufio"


func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    d := sc.Text()
    array1 := strings.Split(d," ")
    a , _ := strconv.Atoi(array1[0])
    b , _ := strconv.Atoi(array1[1])
    c , _ := strconv.Atoi(array1[2])
    
    fmt.Println(a + 1)
    fmt.Println(b + 1)
    fmt.Println(c + 1)
}