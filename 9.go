package main

import "fmt"

func main() {
	var ch = make(chan int,3)
	for i:=0; i<3 ; i++{
		ch <- i
		fmt.Println("Записано!")
	}
	for x:=range ch{
		fmt.Printf("x: %v\n",x)
	}
	fmt.Printf("len: %d cap:%d\n",len(ch),cap(ch))
}