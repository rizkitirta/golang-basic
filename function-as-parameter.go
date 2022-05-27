package main

import "fmt"

// versi 1 function as parameter
func sayHello(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

// versi 1 function as parameter dengan type declaration
type Filter func(string) string

func sayHello2(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println(nameFiltered)
}

func spamFilter(name string) string {
	if name != "anjing" {
		return name
	} else {
		return "***"
	}
}

func main() {
	sayHello("anjing", spamFilter)
	sayHello("melon", spamFilter)
}
