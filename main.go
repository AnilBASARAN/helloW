package main

import (

	"log"
)

func main() {
	var myString string
	myString = "Green"
	log.Println("myString is set to", myString)

	changeUsingPointer(&myString)

	log.Println("second time :",myString)

}

func changeUsingPointer(s *string){
	newValue := "Red"
	*s = newValue
}