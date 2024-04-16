package main

var users []string

func main() {
	AddUser("Alice")
	AddUser("Bob")
}
func AddUser(x string){
	users = append(users, x)
}