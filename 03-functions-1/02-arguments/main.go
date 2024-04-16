package main

import "fmt"

func main() {
	Greet("Alice")
	Greet("Bob")
}
func Greet(greet string){
	fmt.Println("Hello,", greet)
}