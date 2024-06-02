package main

import "fmt"

func main() {
	var ch = make(chan int)
	close(ch)
	x,ok :=<-ch
	fmt.Printf("x: %v,ok: %v\n",x,ok)
}