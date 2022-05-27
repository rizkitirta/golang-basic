package main

import "fmt"

type Customer struct {
	name string
	city string
	age  int
}

// struct method versi 1
func sayHello(customer Customer, name string) {
	fmt.Println("Hallo", name, "saya", customer.name)
}

// struct method versi 2
func (customer Customer) sayHello2(name string) {
	fmt.Println("Kenalin", name, "nama saya", customer.name)
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
	user_3 := Customer{"arif", "kuningan", 18}
	fmt.Println(user_3)

	// call struct method v1
	sayHello(user_1, user_2.name)
	sayHello(user_2, user_3.name)

	// call struct method v2
	user_1.sayHello2(user_2.name)
	user_3.sayHello2(user_1.name)
}
