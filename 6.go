package main

import (
	"fmt"
	"sync"
)

func main() {

	var counters sync.Map

	counters.Store("habr", 42)

	v, ok := counters.Load("otus")
	if ok {
		val := v.(int)
		val++
	}
	counters.Range(func(k, v interface{}) bool {
		fmt.Println("key:", k, ", val:", v)
		return true
	})

	counters.Delete("otus")
}
