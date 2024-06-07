/* go:generate ./command.sh */
//go:generate echo "Hello world!"
package main

import (
	"fmt"
)

func main() {
	fmt.Println("run any unix command in go:generate")
}

//go:generate ls -l
//go:generate -command bye echo "Goodbye!"
//go:generate bye		
//go:generate go run 1.go