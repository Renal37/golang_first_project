package main

type Employee struct {
	name,surname string
}

func FullName(e Employee) string {
	return e.name + " " + e.surname

}

func (e Employee) FullName() string {
	return e.name + " " + e.surname
}
func main() {
	println(Employee{"alex","davy"}.FullName())
	println(FullName(Employee{"alex","davy"}))
}