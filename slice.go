package main

import "fmt"

func main() {
	var months = [...]string{
		"januari",
		"febuari",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}

	var slice = months[1:5]
	var app = append(slice,"bulan")

	fmt.Println(slice)
	fmt.Println(len(slice))
	fmt.Println(cap(slice))
	fmt.Println(app)

	var src = []string{
		"apple", "orange", "banana", "melon",
	}
	var srcAppend = append(src,"strawberry")
	
	var newSlice = make([]string,len(src))
	copy(newSlice,src)

	fmt.Println(newSlice)
	fmt.Println(srcAppend)
	
}