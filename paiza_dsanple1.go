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
    
    h,_ := strconv.Atoi(array1[0])
    //m,_ := strconv.Atoi(array1[1])
    
    if ( h >= 7 && h <= 12) {
        fmt.Println("lunch")
    } else if ((h >= 13 && h <= 18)){
        fmt.Println("dinner")
    } else {
        fmt.Println("breakfast")
    }    
    //fmt.Println("XXXXXX")
}