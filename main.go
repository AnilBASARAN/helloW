package main

import (
	"fmt"
	"net/http"
	"myapp/handlers"
)

const portNumber = ":8080"




func main() {

	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about", handlers.About)
	

	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}

// if you use go run *go instead of go run main.go
// because if you try go main.go it will say ı dont know what Home is and what About is



