package main

import "fmt"

type HasName interface {
	getName() string
}

type User struct {
	name string
	age  int
}

// automatis kontrak ke HasName interface
func (user User) getName() string {
	return user.name
}

func sayHello(hasName HasName) {
	fmt.Println("Hallo", hasName.getName())
}

func main() {
	var user_1 User
	user_1.name = "tirta"
	user_1.age = 20

	sayHello(user_1)

	fmt.Println(user_1.getName())
}
