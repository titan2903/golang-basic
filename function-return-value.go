package main

import "fmt"

func getHello(name string) string {
	result := "hello"
	if name == "" {
		return "hello js"
	} else {
		return result + name
	}
}

func main() {
	result := getHello("Titan")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
