package main 

import "fmt"

func main() {
	strToNum := map[string]int{
		"zero": 0,
		"one": 1,
        "two": 2,
        "three": 3,
	}
	if value, ok := strToNum["zero"]; ok {
		fmt.Println("Zero is inside",value)
	}else{
		fmt.Println("Zero is not inside")
	}

}