package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    var input = readLine("Enter yes or no: ")

    switch strings.ToLower(input) {
    case "yes", "да":
        fmt.Println("you agreed")
    case "no", "нет":
        fmt.Println("you disagreed")
    default:
        fmt.Println("invalid input")
    }
}

func readLine(greeting string) string {
    fmt.Print(greeting)
    reader := bufio.NewReader(os.Stdin)
    text, _ := reader.ReadString('\n')
    return strings.TrimSpace(text)
}