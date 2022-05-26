package main

import (
	"fmt"
)

func main() {
	user := map[string]string{
		"name":  "tirta",
		"email": "tirta@gmail.com",
	}
	
	fmt.Println(user)
	fmt.Println(user["name"])
	fmt.Println(user["email"])

	// change value
	user["email"] = "tirtadev01@gmail.com"

	fmt.Println(user)
	fmt.Println("============")

	// map versi 2
	user2 := make(map[string]string)
	user2["name"] = "joko"
	user2["email"] = "joko@mail.com"

	fmt.Println(user2)

	// delete value map
	delete(user2,"email")
	fmt.Println(user2)
}