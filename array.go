package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Titanio"
	names[1] = "Yudista"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var values = [3]int{
		80,
		94,
		70,
	}

	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])
	fmt.Println(values)

	fmt.Println(len(names))
	fmt.Println(len(values))

}
