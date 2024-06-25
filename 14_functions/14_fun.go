package main
import "fmt"

func main() {
	fmt.Println(sum(2,3,4))
}

func sum (items ...int) int{
	result := 0
	for _,item := range items{
        result += item
    }
	return result
}