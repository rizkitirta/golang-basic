package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println("Hello", i)
	}

	slice := []string{
		"apple", "orange", "watermelon", "strawberry", "blueberry",
	}

	for i := 0; i < len(slice); i++ {
		fmt.Println("fruit :", slice[i])
	}

	// for range
	for index, value := range slice {
		fmt.Println("index ke", index, "buah", value)
	}

	// for range with map
	user := make(map[string]string)
	user["name1"] = "tirta"
	user["name2"] = "joko"
	user["name3"] = "arif"
	user["name4"] = "hanin"

	for key, value := range user {
		fmt.Println(key, "=", value)
	}
}
