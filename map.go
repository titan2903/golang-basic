package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "titanio",
		"address": "subang",
	}

	person["title"] = "software engineer"

	fmt.Println(person)
	fmt.Println(len(person))
	fmt.Println(person["name"])
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	book := make(map[string]string)

	book["title"] = "belajar golang"
	book["author"] = "eko"
	book["uops"] = "salah"

	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "uops")
	fmt.Println(book)
	fmt.Println(len(book))

}
