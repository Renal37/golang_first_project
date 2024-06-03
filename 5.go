// Выдаст ошибку 

package main

import (
	"sync"
)

func main() {

	var wg sync.WaitGroup
	m :=  make(map[string]int)

	for x := 0; x < 12; x++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
				m["hello"] = 1
		}(&wg)
	}
	wg.Wait()
}