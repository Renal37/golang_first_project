package main

import (
	"fmt"
	"reflect"
)

type MyError struct {}

func (e MyError) Error() string {
    return "my error"
}
func main(){
	var e error
	e = MyError{}
	fmt.Println(reflect.TypeOf(e).Name())
	fmt.Printf("%T\n",e)
}