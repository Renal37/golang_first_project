package main

import "fmt"

func main() {
	s:="hello"
	b := []byte(s)
	fmt.Printf("%p %p",&s,&b[0])
}