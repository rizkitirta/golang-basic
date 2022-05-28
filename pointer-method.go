package main

import "fmt"

type User struct {
	name    string
	married bool
}

func (user *User) getMaried() {
	if user.married {
		user.name = "Mr." + user.name
	}

	fmt.Println("di method", user)
}

func main() {
	user_1 := User{name: "Tirta", married: true}
	user_1.getMaried()

	fmt.Println(user_1)
}
