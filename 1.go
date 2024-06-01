package main

import "fmt"

type Adult interface {
    IsAdult() bool
    fmt.Stringer
}

type Person struct {
    age int
    name string
}

func (p Person) String() string { 
    return fmt.Sprintf("%s is %d age",p.name,p.age)
}

func (p Person) IsAdult() bool {
    return p.age >= 18
}

func adultFiler(people []Adult) []Adult {
    adults := make([]Adult, 0)
    alreadyPrinted := false
    for _, p := range people {
        if p.IsAdult() {
            adults = append(adults, p)
        }
        if !p.IsAdult() && !alreadyPrinted {
            fmt.Println("Не имеется человек вышел вашего запроса")
            alreadyPrinted = true
        }
    }
    return adults
}

func main() {
    people := []Adult{Person{19,"Renal"}, Person{18,"Mars"}, Person{45,"Mary"}}
    fmt.Println(adultFiler(people))
}