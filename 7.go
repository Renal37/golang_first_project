package main

import "fmt"

func main() {
	var ch = make(chan int,3)
	for i:=0; i<3 ; i++{
		ch <- i
		fmt.Println("Записано!")
	}
	for j := 0; j < 3; j++ {
	fmt.Printf("Читаем: %v",<-ch)
	}
}