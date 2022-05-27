package main

import "fmt"

func getFullName() (firstName string, middleName string, lastName string) {
	firstName = "Rizki"
	middleName = "Tirta"
	lastName = "Buwono"

	return
}

func main() {
	first, middle, last := getFullName()

	fmt.Println(first)
	fmt.Println(middle)
	fmt.Println(last)
}
