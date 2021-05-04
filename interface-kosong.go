package main

import "fmt"

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "ups"
	}
}

func main() {
	var data interface{} = Ups(7)
	fmt.Println(data)
}

/*
? Interface kosong tidak memperdulikan tipe data apapun
? semua tipe data dapat di kembalikan oleh interface kosong
*/
