package main


import "fmt"
import "os"
import "bufio"
import "strconv"
import "strings"
import "math"

func main(){
    // 自分の得意な言語で
    // Let's チャレンジ！！
    
    sc := bufio.NewScanner(os.Stdin)
    
    sc.Scan()
    data1 := sc.Text()
    array1 := strings.Split(data1," ")
    
    t_sum := 0
    b_sum := 0
    
    sc.Scan()
    data2 := sc.Text()
    array2 := strings.Split(data2," ")
    
    for i := 0;i < 5;i++{
        num1, _ := strconv.Atoi(array1[i])
        flag1 := array2[i]
        
        t_sum += num1
        
        if flag1 == "o" {
            b_sum += num1
        }
        
        
    }
    
    var bb_sum float64
    var tt_sum float64
    
    bb_sum = float64(b_sum)
    tt_sum = float64(t_sum)
    
    //fmt.Println(bb_sum)
    //fmt.Println(tt_sum)
    result1 := float64(bb_sum / tt_sum) 
    result := math.Floor(result1 * 100)
    
    fmt.Println(result)
}