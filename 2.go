package main

import (
    "fmt"
    "strconv"
)

type Temp int

func (t Temp) String() string {
    return strconv.Itoa(int(t)) + " С"
}

func main() {
    x := Temp(24)
    fmt.Printf("%v %T\n", x, x)
}