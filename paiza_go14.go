package main

import ("fmt"
        //"bufio"
        "strconv"
        //"os"
        )

func main() {
	greeting := "Hello paiza"
	for i := 0 ; i < 5 ; i++ {
	    fmt.Println(greeting + strconv.Itoa(i))
	}
}
