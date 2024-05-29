package main
// Поиск в map
import "fmt"

func main() {
	var a = make(map[string]int)
	a["test"] = 1
	a["qwee"] = 3

	fmt.Printf("%v\n", a)

	c := &a
	fmt.Printf("%v\n", (*c)["test"])
}
