package main
import "fmt"

type Age int 

func (i *Age) Print(){
	*i++
	fmt.Println(*i)
}
	

func main() {
    i := Age(20)
    i.Print()
	i.Print()
	i.Print()
}