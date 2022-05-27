package main

import "fmt"

func fruit() (string, int, string) {
	return "apple", 12000, "buah apel"
}

func main() {
	buah, harga, _ := fruit()

	fmt.Println(buah)
	fmt.Println(harga)
}
