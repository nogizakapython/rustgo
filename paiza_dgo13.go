import "fmt"
import "bufio"
import "os"
import "strings"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    sc := bufio.NewScanner(os.Stdin)
    sc.Scan()
    s := sc.Text()
    array1 := strings.Split(s," ")
    n := len(array1)
    
    for i := 0 ;i<n;i++ {
        fmt.Println(array1[i])
    }
   // fmt.Println("XXXXXX")
}