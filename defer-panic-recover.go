package main

import "fmt"

func logging() {
	// recover menangkap pesan error
	err := recover()
	if err != nil {
		fmt.Println("Error With Message :", err)
	}
	fmt.Println("Aplikasi Selesai")
}

func runApp(err bool) {
	// defer dipanggil diakhir function
	defer logging()
	if err {
		// throw error with panic
		panic("Aplikasi Error!")
	}

	fmt.Println("Aplikasi Berjalan")
}

func main() {
	runApp(false)
}
