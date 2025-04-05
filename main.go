package main

import "fmt"

func main() {
	fmt.Println("hello, world")

	var whatToSay string

	whatToSay = "Goodbye"

	fmt.Println(whatToSay)

	var number int
	number = 10

	fmt.Println("number is set to :",number)

	whatWasSaid := saySomething()

	fmt.Println(whatWasSaid)
}

func saySomething() string {
	return "something"
}