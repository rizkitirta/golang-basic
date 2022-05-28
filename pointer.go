package main

import "fmt"

type User struct {
	name string
	age  int
	city string
}

func main() {
	user_1 := User{
		name: "Tirta",
		age:  20,
		city: "Lampung",
	}

	// Pass by value atau duplikasi data
	user_2 := &user_1
	user_2.city = "Jakarta"

	// Pass by reference (dengan tanda &) atau mereferensi ke data / memori yang sama
	var user_3 *User = &user_1
	user_3.city = "Depok"

	*user_3 = User{
		name: "arif",
		age:  18,
		city: "kuningan",
	}

	fmt.Println(user_1)
	fmt.Println(user_2)
	fmt.Println(user_3)

	user_4 := new(User)
	user_4.name = "hanin"

	fmt.Println(user_4)
}
