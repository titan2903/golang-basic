package main

import "fmt"

func main() {
	name := "titan"
	counter := 0

	increment := func() {
		name := "rudi"

		fmt.Println("increment")
		counter++
		fmt.Println(name)

	}

	increment()

	fmt.Println(counter)
	fmt.Println(name)
}
