package main

import "fmt"

func main() {
	name := "yudista"

	if name == "test" {
		fmt.Println("hello test")
	} else if name == "titan" {
		fmt.Println("hello titan")
	} else {
		fmt.Println("masuk sini")
	}

	// var length = len(name)
	if length := len(name); length > 5 {
		fmt.Print("terlalu panjang")
	} else {
		fmt.Println("nama sudah benar")
	}
}
