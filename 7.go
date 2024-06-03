package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Dog struct{ name string }

func (d *Dog) Bark() { fmt.Printf("%s\n", d.name) }

var dogPack = sync.Pool{
	New: func() interface{} { return &Dog{} },
}

func tryDog() {
	dog := dogPack.Get().(*Dog)
	if rand.Intn(1000)%10 == 0 {
		dog.name = "ivan"
	} else {
		dog.name = "leni"
	}
	dog.Bark()
	dogPack.Put(dog)
}

func main() {
	for i := 0; i < 1000; i++ {
		tryDog()
	}
}
