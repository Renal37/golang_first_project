package main

import "fmt"

type A struct{
	C int
}
func main(){
	c:= []int{100,2,3}
	a := &c[0]
	c[0]++
	fmt.Printf("%v %v\n",c,*a)
	*a++
	fmt.Printf("%v %v\n",c,*a)
	c = append(c, 9,8,7)
	fmt.Printf("%v %v\n",c,*a)
	c[0]++
	fmt.Printf("%v %v\n",c,*a)
}