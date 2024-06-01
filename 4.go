package main

import (
  "fmt"
)

type Hound interface {
     Hunt()
}

type Poodle interface {
     Bark() string
}

type Decorative interface {
  Sleep(hours int)
}

type Bloodhound struct{name string}
func (b Bloodhound) Hunt() { fmt.Printf("%s is hunting\n", b.name) }

type TeacupPoodle struct{name string}
func (p TeacupPoodle) Bark() string { return fmt.Sprintf("%s is barking...\n", p.name) }

type ShihTzu struct{name string}
func (p ShihTzu) Sleep(horts int) { fmt.Printf("%s is sleeping...\n", p.name) }


func zoo(dog interface{}) {
   switch dog.(type) {  
  case Hound:
    h := dog.(Hound)
    h.Hunt()
  case Poodle:
    p := dog.(Poodle)
    p.Bark()
  case Decorative:
    d := dog.(Decorative)
    d.Sleep(1)
}
}

func main() {
  zoo(Bloodhound{"misha"})
  zoo(TeacupPoodle{"vasya"})
  zoo(ShihTzu{"shihTzu"})
}
// 1.26.44