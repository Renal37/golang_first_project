package main

import "fmt"

func main() {
	var ch = make(chan int,3)
	for i:=0; i<3 ; i++{
		ch <- i
		fmt.Println("Записано!")
	}
	fmt.Printf("len: %d cap:%d\n",len(ch),cap(ch)) 
}