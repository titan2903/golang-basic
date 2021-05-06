package main

import (
	"container/list"
	"fmt"
)


func main() {
	data := list.New()

	data.PushBack("titan")
	data.PushBack("eko")
	data.PushBack("budi")
	data.PushFront("andi")

	// data.Front().Next().Next().Next() //! mengambil data selanjutnya
	// data.Front().Prev().Prev().Prev() //! mngambil data sebelumnya

	// fmt.Println(data.Front().Value) //! mengambil data paling depan
	// fmt.Println(data.Back().Value) //! mengambil data paling belakang

	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println("Dari belakang ke depan: ",element.Value)
	}
	
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println("Dari depan ke belakang: ", element.Value)
	}
}

/*
	tidak ada batas penyimpanan data
	data akhir pasti nil
	tidak bisa diakses menggnakan index
	Pushback memasukkan data di akhir
	pushfront memasukkan data di awal
*/