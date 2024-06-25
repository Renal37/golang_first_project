package main

import "fmt"

func main() {
	c:=10
	s :=make([]int,0, c)
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
}