package main

import "fmt"

func add(i *string){
	*i = "PET"
	{
		i:="TEST"
		fmt.Println(i)
	}
	fmt.Println(*i)
}
func main() {
	i := "CHECK"
	fmt.Println(i)
	add(&i)
	fmt.Println(i)
}