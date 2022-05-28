package main

import "strconv"

func newMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{"name": name}
	}
}

func main() {
	data := newMap("tirta")
	if data == nil {
		println("data is nil")
	} else {
		println("data is not nil & data is", strconv.Quote(data["name"]))
	}
}