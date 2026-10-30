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
    
    x := sc.Text()
    array1 := strings.Split(x," ")
    
    s,_ := strconv.ParseFloat(array1[0],64)
    a,_ := strconv.ParseFloat(array1[1],64)
    t,_ := strconv.ParseFloat(array1[2],64)
    ans := int((( a * t) / ( s * a)) * 100)
    
    fmt.Println(ans)
}