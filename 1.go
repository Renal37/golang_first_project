package main

import (
    "fmt"
    "sync"
    "time"
)

type Dog struct {
    name         string
    walkDuration time.Duration
}

func (d Dog) Walk() {
    fmt.Printf("%s is taking a walk\n", d.name)
    time.Sleep(d.walkDuration)
    fmt.Printf("%s is going home\n", d.name)
}

func walkTheDogs(dogs []Dog, wg *sync.WaitGroup) chan struct{} {
    done := make(chan struct{})
    wg.Add(len(dogs))

    for _, d := range dogs {
        go func(d Dog) {
            defer wg.Done()
            d.Walk()
            done <- struct{}{}
        }(d)
    }

    return done
}

func main() {
    dogs := []Dog{{"vasya", time.Second}, {"john", time.Second * 3}}
    var wg sync.WaitGroup
    done := walkTheDogs(dogs, &wg)

    go func() {
        wg.Wait()
        fmt.Println("everybody's home")
        close(done)
    }()

    select {
    case <-done:
        fmt.Println("One dog is back home")
    default:
        time.Sleep(time.Second * 5)
        fmt.Println("All dogs should be back home")
    }
}