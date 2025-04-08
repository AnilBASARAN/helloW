package main

import (
	"log"
	"time"
)


var s = "seven"
// also fmt can be used too instead of log, its a valid approach also when it comes to console.log
func main() {


	type User struct {
		FirstName string
		LastName string
		PhoneNumber string
		Age int
		BirthDate time.Time
	}
}

func saySomething(s string) (string,string) {
	return s, "World"
}

// pointer is , it points to a spesific location in memory and gives you a means of getting that particular location in memory