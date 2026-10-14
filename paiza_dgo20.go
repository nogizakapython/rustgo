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
    
    data := sc.Text()
    array1 := strings.Split(data," ")
    
    A,_ := strconv.Atoi(array1[0])
    B,_ := strconv.Atoi(array1[1])
    C,_ := strconv.Atoi(array1[2])
    
    if A * B <= C {
        fmt.Println("YES")
    } else {
        fmt.Println("NO")
    }
    
    //fmt.Println("XXXXXX")
}