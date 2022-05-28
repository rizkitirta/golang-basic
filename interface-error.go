package main

import (
	"errors"
	"fmt"
)

func Pembagi(number int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi tidak boleh 0")
	} else {
		return number / pembagi, nil
	}
}

func main() {
	value, err := Pembagi(6, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Result", value)
	}
}
