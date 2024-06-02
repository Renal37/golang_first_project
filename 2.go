package main

import "fmt"

func main() {
	var ch = make(chan int)
    go func() {
        fmt.Println("Hello, world")
    	ch <- 1
    }()
	fmt.Printf("x: %v\n",<-ch)
}