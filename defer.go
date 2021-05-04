package main

import "fmt"

func logging() {
	fmt.Println("selesai memanggil function")
}

func runApp(value int) {
	defer logging()
	fmt.Println("run app")
	result := 10 / value
	fmt.Println("result: ", result)
}

func main() {
	runApp(10)
}
