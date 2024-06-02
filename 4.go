// Отправка в закрытый канал и паника 
package main

func main() {
	var ch = make(chan int)
	close(ch)
	ch <- 3
}
