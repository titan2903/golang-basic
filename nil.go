package main

import "fmt"

func NewMap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}

func main() {
	person := NewMap("eko")

	if person == nil {
		fmt.Println("data kkosong")
	} else {
		fmt.Println(person)
	}
}

/*
? type data nil bisa di gunakan hanya di beberapa tipe data seperti interface, function, pointer, map, slice dan channel
*/
