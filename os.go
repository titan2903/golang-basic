package main

import (
	"fmt"
	"os"
)


func main() {
	args := os.Args //! Mengambil data argument yang ada di golang

	fmt.Println(args) 
	fmt.Println(args[1])
	fmt.Println(args[2])

	name, err := os.Hostname() //! Mengambil nama os di Komputer atau laptop

	if err == nil {
		fmt.Println(name)
	} else {
		fmt.Println(err.Error())
	}

	// ! Setting env di aplikasinya
	username := os.Getenv("APP_USERNAME") 
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}