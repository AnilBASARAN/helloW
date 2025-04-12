package main

import (
	"fmt"
	"net/http"
	"myapp/handlers"
)

const portNumber = ":8080"

//bugün de boş geçmesin diye bir commit atalım 


func main() {

	http.HandleFunc("/",handlers.Home)
	http.HandleFunc("/about", handlers.About)
	

	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}

// if you use go run *go instead of go run main.go
// because if you try go main.go it will say ı dont know what Home is and what About is
// instead of compiling just main.go, which is go run main.go does, you want to compile all of the files which 
// you can achieve with * (Asterix)



