// Замыкающая функция спросить

package main
import "fmt"

func logger(prefix string) func(s string) {
	return func(s string) {
        fmt.Printf("[%s] %s\n",prefix, s)
    }
}
func main() {
	warn := logger("WARNING")
	err := logger("ERROR")
	warn("TEST")
	err("TEST!!!!")
}