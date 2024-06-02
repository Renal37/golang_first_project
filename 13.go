package main

import "fmt"

func t1() error {
	for {
	}
}

func t2() error {
	for {
	}
}

func main() {
	chErr := make(chan error, 2)

	go func() {
		chErr <- t1()
	}()

	go func() {
		chErr <- t2()
	}()

	for err := range chErr {
	   fmt.Println(err)
	}
}
