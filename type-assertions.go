package main

import "fmt"

func random(value interface{}) interface{} {
	return value
}

func main() {
	var data interface{} = random("Hallo gaes")

	// v1
	// var resultString string = data.(string)
	// fmt.Println(resultString)

	// v2
	switch value := data.(type) {
	case string:
		fmt.Println("string", value)
	case int:
		fmt.Println("int", value)
	}
}
