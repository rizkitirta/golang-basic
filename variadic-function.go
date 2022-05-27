package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}

	return total
}

func getFruit(fruit ...string) string {
	for _, value := range fruit {
		fmt.Println("Buah:", value)
	}

	return "ok"
}

func main() {
	total := sumAll(10, 20, 30, 40, 50)
	fmt.Println(total)

	// with slice
	slice := []int{10, 20, 30, 40, 50, 60}
	total2 := sumAll(slice...)

	// with slice
	slice2 := []string{"apel", "mangga", "jeruk", "melon"}
	res := getFruit(slice2...)

	fmt.Println(total2)
	fmt.Println(res)
}
