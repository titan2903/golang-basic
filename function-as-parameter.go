package main

import "fmt"

type Filter func(string) string //! Jika function as parameter terlalu panjang bisa diubah ke bentuk type seperti ini

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Titan", spamFilter)
	sayHelloWithFilter("Anjing", spamFilter)
}
