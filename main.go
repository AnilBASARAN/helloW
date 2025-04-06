package main

import (

	"log"
)
// also fmt can be used too instead of log, its a valid approach also when it comes to console.log
func main() {
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)

	log.Println("second time :",myString)

}

func changeUsingPointer(s *string){
	// this is a different way than usual,  ":=" this is declaring and equaling to the variable at the same time
	newValue := "Red"
	*s = newValue
}

// pointer is , it points to a spesific location in memory and gives you a means of getting that particular location in memory