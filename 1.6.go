package main

import "fmt"

func Sum(a ...int)(int){
	b :=[]int{9,8}
	b = append(b,a...)
	fmt.Println(b)
    s :=0
    for _, i := range a {
        s+=i
    }
    return s
}
func main() {
    s := Sum(234,43,12,4,3,2,5,6,123,4,1)
    fmt.Println(s)
}