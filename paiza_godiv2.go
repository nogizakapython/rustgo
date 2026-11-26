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
    n, _ := strconv.Atoi(sc.Text())
    
    sc.Scan()
    data := sc.Text()
    array1 := strings.Split(data," ")
    
    
    for i := 0 ; i < n ; i++ {
        fmt.Println(array1[i])
    }
    //fmt.Println("XXXXXX")
}