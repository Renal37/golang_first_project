package main

import "fmt"
import "github.com/pkg/errors"

func main() {
	err := errors.New("error")
	err = errors.Wrap(err,"open failed")
	err = errors.Wrap(err,"read failed")
	fmt.Println(err)
}
// Не работает 