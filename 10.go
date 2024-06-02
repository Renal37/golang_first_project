package main

import(
	"fmt"
	"time"
)

func main(){
	ch := make(chan string,1)
	go func() {
		time.Sleep(9 * time.Second)
		ch<-"Hello"
	}()
	timer := time.NewTimer(10 * time.Second)

	select{
	case data := <-ch:
		fmt.Printf("received: %v\n",data)
	case <-timer.C:
		fmt.Printf("failed to receive in 10s\n")
	}
}