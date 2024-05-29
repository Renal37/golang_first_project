package main

import "fmt"

var counter int

type User struct {
	UserId int
	loginData struct{
		Login string
		Password string
	}
	previos *User
	counter *int
}

type Signal struct {}

func NewUser(login string) *User {
	counter++
	return &User{
		loginData: struct{
			Login string
            Password string
		}{
			Login:login,
		},
		counter:&counter,
	}
}

func main(){
	user1 := NewUser("Petrov")
	user2 := User{/* 100,"Vasya","mypass",user1 */}
	fmt.Printf("%v %v\n",user1 , user2)
}