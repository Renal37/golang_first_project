package main

import (
	"fmt"
	"os"
)

func main() {
	var env []string
	env = os.Environ()  // слайс (!) 
	fmt.Println(env[0]) // NEWVAR=val
	var newvar string
	newvar, ok := os.LookupEnv("NEWVAR")
	fmt.Printf(newvar,ok)          // val
	os.Setenv("NEWVAR", "val2") // установить
	os.Unsetenv("NEWVAR")
	// удалить
	fmt.Printf(os.ExpandEnv("$USER have a ${NEWVAR}")) // "шаблонизация"
}
