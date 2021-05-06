package main

import (
	"container/ring"
	"fmt"
	"strconv"
)


func main() {
	var data *ring.Ring = ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "DAta " + strconv.FormatInt(int64(i), 10) //! Proses convert datanya
		data = data.Next() //! proses memasukkan data ke selanjutnya
	}

	fmt.Println(*data) //! tidak memiliki function untuk di Prtinln

	data.Do(func(i interface{}) {
		fmt.Println(i)
	})

	/*
		? untuk kasus ini jangan melakukan perulangan menggunakan for loop karena ada data.Next(), oleh karena itu perulangan tidak pernah berhenti
	*/
}