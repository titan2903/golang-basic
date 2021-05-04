package main

import (
	"fmt"
)

func endApp() { //! DEFER FUNCTION
	message := recover() //! MENGHENTIKAN PROSES PANIC DAN AKAN TETAP MELANJUTKAN PROGRAMNYA
	if message != nil {
		fmt.Println("error dengan message", message)
	}
	fmt.Println("aplikasi selesai")
}

//! RECOVER DI PASTIKAN KHARUS DI GUNAKAN DI DEFER FUNCTION

func runApp(error bool) {
	defer endApp()
	if error {
		panic("aplikasi error")
	}

	fmt.Println("aplikasi berjalan")
}

func main() {
	runApp(true)
	fmt.Println("masuk")
}
