package main

import "fmt"

func main() {
	var arr [3]string

	arr[0] = "tirta"
	arr[1] = "hanin"
	arr[2] = "arif"

	var arr2 = [2]string{
		"hana",
		"hani",
	}

	var countArr1 = len(arr)
	var countArr2 = len(arr2)

	fmt.Println(arr)
	fmt.Println(arr2)
	
	fmt.Println(countArr1)
	fmt.Println(countArr2)
}
