package main

import (
	"log"
	"time"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main() {

user := User {
	FirstName: "Max",
	LastName: "Payne",
}

log.Println(user.FirstName)
// use user field FirstName by using dot notation
	
}
// go doesnt have private, and public, if starts with capital letter, its accessible, else its just visible on that package 
func whatever(){
//this is just available on this package
}


