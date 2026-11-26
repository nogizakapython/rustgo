package main
import "fmt"
import "os"
import "bufio"
import "strconv"
import "strings"


func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    L, _ := strconv.Atoi(sc.Text())
    fmt.Println(L+1)
    sc.Scan()
    data := sc.Text()
    array1 := strings.Split(data," ")
    u , _ := strconv.Atoi(array1[0])
    a , _ := strconv.Atoi(array1[1])
    b , _ := strconv.Atoi(array1[2])
    fmt.Println(u + 1)
    fmt.Println(a + 1)
    fmt.Println(b + 1)
    sc.Scan()
    v, _ := strconv.Atoi(sc.Text())
    fmt.Println(v+1)
}