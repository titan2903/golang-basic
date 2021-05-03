package main

import "fmt"

func main() {
	var name1 = "titan"
	var name2 = "titan"
	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
