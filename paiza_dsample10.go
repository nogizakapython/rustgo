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
    data1 := sc.Text()
    
    array1 := strings.Split(data1," ")
    a , _ := strconv.Atoi(array1[0])
    b , _ := strconv.Atoi(array1[1])
    
        sc.Scan()
    h , _ := strconv.Atoi(sc.Text())
    
    
    fmt.Println((a+b) * h / 2)
}