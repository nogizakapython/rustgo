package main
import "fmt"
import "os"
import "bufio"
import "strconv"


func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    array1 := [5]string{"A","B","C","D","E"}
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    
    n , _ := strconv.Atoi(sc.Text())
    
    
    
    fmt.Println(array1[n - 1])
}