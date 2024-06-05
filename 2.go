package main

import (
	"fmt"
	"flag"
)

var msg = flag.String("msg", "default-msg","")

func init(){
	flag.Parse()
}

func main() {
    fmt.Println(msg,*msg)
}