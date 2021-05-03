package main

import "fmt"

func sayHallTo(firstName string, lastName string) {
	fmt.Println("hello", firstName, lastName)
}

func main() {
	FisrtName := "titanio"
	sayHallTo(FisrtName, "yudista")
	sayHallTo("Budi", "Nugraha")
}
