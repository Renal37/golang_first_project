package main
import "fmt"

type Student struct {
	name string
}

func studentName(name string) string {
	s := Student{name}
	return s.name
}
func main() {
	println(studentName("Renal"))
	student := Student{"alex"}
	fmt.Println(student)
}
