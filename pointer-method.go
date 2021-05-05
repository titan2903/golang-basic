package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println(man.Name)
}

func main() {
	eko := Man{"eko"}

	eko.Married()
	fmt.Println(eko.Name)
}

/*
 ?jika ingin membuat struct function atau struct method di sarankan untuk menggunakan pointer di bandingkan value biasa
 ?lebih hemat memory karena tidak menduplikat datanya
*/
