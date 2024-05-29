package main

import "fmt"

func main() {
	var a int =1
	var b *int = &a
	fmt.Printf("%d %d\n",a,*b)
	a++
	fmt.Printf("%d %d\n",a,*b)
	*b++
	fmt.Printf("%d %d\n",a,*b)

	c:=100
	b = &c

	fmt.Printf("%d %d\n",a,*b)
	
}