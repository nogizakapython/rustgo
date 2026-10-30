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
    
    x := sc.Text()
    array1 := strings.Split(x," ")
    array2 := make([]int , 0 , len(array1))
    
    sum_v := 0
    max_v := -1
    min_v := 10001
    
    for i := 0;i<len(array1);i++{
        d, _ := strconv.Atoi(array1[i])
        array2 = append(array2,d)
        if d > max_v {
            max_v = d
        }
        
        if d < min_v {
            min_v = d
        }
        
    }
    
    for j := 0; j < len(array2) ; j++ {
        sum_v += array2[j]
        
    }
    
    ans := sum_v - max_v - min_v
    
    
    
    fmt.Println(ans)
    
    
}