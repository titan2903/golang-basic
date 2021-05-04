package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
	// Married bool
}

func (customer Customer) sayHai(name string) {
	fmt.Println("Hallo", name, "My name is", customer.Name)
}

func (a Customer) sayHuu() {
	fmt.Println("Huu from ", a.Name)
}

func main() {
	// ! tiga cara memasukkan atau input value di Struct
	var eko Customer
	eko.Name = "eko"
	eko.Address = "indonesia"
	eko.Age = 30

	eko.sayHai("Joko")
	eko.sayHuu()
	// fmt.Println(eko.Name)
	// fmt.Println(eko.Address)
	// fmt.Println(eko.Age)
	// fmt.Println(eko)

	// joko := Customer{
	// 	Name:    "joko",
	// 	Address: "Cirebon",
	// 	Age:     35,
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Jakrata", 45}
	// fmt.Println(budi)

}
