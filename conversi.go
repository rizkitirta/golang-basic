package main

import "fmt"

func main() {
	// Coversi Integer
	nilai32 := 127
	nilai64 := int64(nilai32)
	nilai8 := int8(nilai32)

	// Conversi String
	name := "tirta"
	t := string(name[0])

	
	fmt.Println("conversi int",nilai64)
	fmt.Println("conversi int",nilai8)
	fmt.Println("conversi str",t)
}
