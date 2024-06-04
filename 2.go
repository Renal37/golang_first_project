package main

import (
	"fmt"
	"os"
)

func main() {
	b := make([]byte, 1000)
	f, err := os.Open("dev/random")
	fmt.Println(err)
	d, err := f.ReadAt(b, 10)
	written, err := io.Copy()
	fmt.Println(d,err)
}
