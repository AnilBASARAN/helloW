package main

import "log"


var s = "seven"
// also fmt can be used too instead of log, its a valid approach also when it comes to console.log
func main() {
var s2 = "six"

log.Println("m is",s)
log.Println("s2 is",s2)
log.Println(saySomething("xxx"))
}

func saySomething(s string) (string,string) {
	return s, "World"
}

// pointer is , it points to a spesific location in memory and gives you a means of getting that particular location in memory