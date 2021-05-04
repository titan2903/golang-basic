package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	// Married bool
}

func main() {
	// ! tiga cara memasukkan atau input value di Struct
	var eko Customer
	eko.Name = "eko"
	eko.Address = "indonesia"
	eko.Age = 30
	fmt.Println(eko.Name)
	fmt.Println(eko.Address)
	fmt.Println(eko.Age)
	fmt.Println(eko)

	joko := Customer{
		Name:    "joko",
		Address: "Cirebon",
		Age:     35,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakrata", 45}
	fmt.Println(budi)

}
