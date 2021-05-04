package main

import "fmt"

type Person struct {
	Name string
}

type Animal struct {
	Name string
}

type HasName interface {
	GetName() string
}

func (person Person) GetName() string {
	return person.Name
}

func (animal Animal) GetName() string {
	return animal.Name
}

func SayHello(hasName HasName) {
	fmt.Println("Hello", hasName.GetName())
}

func main() {
	var eko Person

	eko.Name = "eko"

	SayHello(eko)

	animal := Animal{
		Name: "Kucing",
	}
	SayHello(animal)
}
