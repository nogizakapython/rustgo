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
    d1 := sc.Text()
    array1 := strings.Split(d1," ")
    m, _ := strconv.Atoi(array1[0])
    n, _ := strconv.Atoi(array1[1])
    
    sc.Scan()
    d2 := sc.Text()
    array2 := strings.Split(d2," ")
    a, _  := strconv.Atoi(array2[0])
    b, _ := strconv.Atoi(array2[1])
    
    fmt.Println(m + 1)
    fmt.Println(n + 1)
    fmt.Println(a + 1)
    fmt.Println(b + 1)
    
}