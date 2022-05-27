package main

import "fmt"

type Blacklist func(string) bool

func checkUser(user string, blacklist Blacklist) {
	if blacklist(user) {
		fmt.Println("You are blocked", user)
	} else {
		fmt.Println("Welcome", user)
	}
}

func main() {
	isAdmin := func(user string) bool {
		return user == "admin"
	}

	checkUser("tirta", isAdmin)

	checkUser("admin", func(user string) bool {
		return user == "admin"
	})
}
