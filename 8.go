package main

import "sync"

type OldDog struct {
	name string
	die  sync.Once
}

func (d *OldDog) Die() {
	d.die.Do(func() { println("bye!") })
}
func main() {
	d := OldDog{name: "bob"}
	d.Die()
	d.Die()
	d.Die()
}
