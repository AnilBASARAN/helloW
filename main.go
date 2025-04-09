package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r * http.Request){
		n, err := fmt.Fprintf(w, "This is the Home Page")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of Bytes written: %d ", n ))
	
}
func About(w http.ResponseWriter, r * http.Request){
	sum := addValues(2,2)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum ))

}

func addValues(x,y int ) int {

	return x + y
}

func main() {

	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)

	fmt.Println(fmt.Sprintf("Starting application on port %s",portNumber))
	_ = http.ListenAndServe(portNumber,nil)
}




