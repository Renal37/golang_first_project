package main

import "fmt"

func Get() (int, int) {
	return 3, 6
}

func Square(z int , c int) (s int, d int) {
	x, y := z+c, z-c
	s = x * x
	d = y * y
	return
}

func main() {
	Hi := func() {
		fmt.Println("HI")
	}
	Hi()
	f := func() int {
		return 100
	}
	fmt.Println(f())

	a, b := Get()
	fmt.Println(a, b)
	z, c := Square(10, 1)
	fmt.Println(z,c)
}
