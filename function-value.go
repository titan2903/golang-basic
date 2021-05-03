package main

import "fmt"

func getGoodBye(name string) string {
	return "good bye " + name
}

func main() {
	sayGoodBye := getGoodBye
	fmt.Println(sayGoodBye("Titan"))
}
