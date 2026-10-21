package main
import "fmt"
import "bufio"
import "os"
import "strconv"




func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    n := 5
    var array1 [5]int
    sc := bufio.NewScanner(os.Stdin)
    
    min_value := 101
    
    for i := 0 ; i < n ; i++ {
        sc.Scan()
        data , _ := strconv.Atoi(sc.Text())
        array1[i] = data
    }
    
    for j := 0 ; j < n ; j++ {
        if array1[j] < min_value {
            min_value = array1[j]
        }
    }
    
    fmt.Println(min_value)
}