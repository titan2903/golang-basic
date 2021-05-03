package main

import "fmt"

func main() {
	var a = 10
	var b = 20
	c := a + b

	fmt.Println(c)

	var i = 10
	i += 10
	fmt.Println(i)

	// i++
	// fmt.Println(i)

	i--
	fmt.Println(i)
}
