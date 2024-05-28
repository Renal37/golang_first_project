// рекурсивная функция 

package main

import "fmt"


func Recursive(number int) int{
	if number == 0{
        return 1
    }
    return number + Recursive(number-1)
}

func main() {
    fmt.Println(Recursive(4))
}