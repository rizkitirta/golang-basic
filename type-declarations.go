package main

import "fmt"

func main() {
	type KTP string
	type MARRIED bool
	
	var noKtp KTP = "1889264213445"
	var statusMarried MARRIED = false

	println(noKtp)
	fmt.Println(statusMarried)
}