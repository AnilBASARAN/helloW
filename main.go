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


