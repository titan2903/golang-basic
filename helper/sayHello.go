package helper

import "fmt"

var Application = 3
var version = 1

func SayHallo(name string) { //! Jika ingin import atau export suatu func harus menggunakan CEMAL CASE
	fmt.Print("hallo ", name)
}


/*
?	Bisa di akses ke package lain karena di awali dengan huruf besar
? 	Jika di awali dengan huruf kecil maka tidak dapat di akses ke package lain 
*/