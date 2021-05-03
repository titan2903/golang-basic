package main

import (
	"fmt"
)

func main() {
	name := "titan"
	switch name {
	case "titan":
		fmt.Println("titan")
		fmt.Println("titanio")
	case "joko":
		fmt.Println("joko")
		fmt.Println("jokooo")
	default:
		fmt.Println("!found")
		fmt.Println("not found")
	}

	// switch length := len(name); length > 5 {
	// case true:
	// 	fmt.Println("nama terlalu panjang")
	// case false:
	// 	fmt.Println("nama sudah benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("nama terlalu panjang")
	case length > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("nama sudah benar")
	}
}
