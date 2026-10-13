package main
import "fmt"
import "bufio"
import "os"
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
    
    N := 0
    
    N = N + A 
    N = N * B 
    N = N % C
    
    fmt.Println(N)
}