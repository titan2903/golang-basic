package main

import "fmt"

func getComplateName2() (firstName string, middleName string, lastName string) {
	firstName = "Eko"
	middleName = "kennedy"
	lastName = "kurniawan"

	return
}

func main() {
	a, b, c := getComplateName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}
