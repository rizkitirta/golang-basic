package main

import "fmt"

type Customer struct {
	name string
	city string
	age  int
}

func main() {
	// versi 1
	var user_1 Customer
	user_1.name = "tirta"
	user_1.city = "lampung"
	user_1.age = 20

	fmt.Println(user_1)

	// versi 2
	user_2 := Customer{
		name: "joko",
		city: "klaten",
		age:  19,
	}

	fmt.Println(user_2)

	// versi 3
	user_3 := Customer{"arif","kuningan",18}
	fmt.Println(user_3)
}
