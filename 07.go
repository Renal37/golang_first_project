package main

import (
    "encoding/json"
    "fmt"
)

type Person struct {
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Fathername string `json:"fathername,omitempty"`
	Age     int    `json:"age"`
}

func main() {
	p := Person{
        Name:    "Renal",
        Surname: "Mustafin",
        // Fathername: "Ranisovith",
        Age:     18,
    }

    js, err := json.Marshal(p)
    if err!= nil {
        panic(err)
    }

    fmt.Println(string(js))
}