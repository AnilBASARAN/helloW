package main

import (
	"fmt"
	"net/http"
)

func Home(w http.ResponseWriter, r * http.Request){
		n, err := fmt.Fprintf(w, "This is the Home Page")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("Number of Bytes written: %d ", n ))
	
}
func About(w http.ResponseWriter, r * http.Request){
	sum := AddValues(2,2)

	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2 + 2 is %d", sum ))

}

func AddValues(x,y int ) int {

	return x + y
}

func main() {

	http.HandleFunc("/",Home)
	http.HandleFunc("/about",About)

	_ = http.ListenAndServe(":8080",nil)
}




