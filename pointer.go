package main

import (
	"fmt"
)

type Address struct {
	City, Province, Country string
}

func ChangeCountryToIndonesia(address *Address) {
	address.Country = "Indonesia"
}

func main() {
	// address1 := Address{"Subang", "Jabar", "indonesia"}
	// address2 := address1

	// address2.City = "bandung"
	// fmt.Println(address1)
	// fmt.Println(address2)

	// address1 := Address{"Subang", "Jabar", "indonesia"}
	// address2 := &address1 //! mengubah value yang ada di address1 dan ini merupakan pointer
	// address2.City = "bandung"

	// fmt.Println(address1)
	// fmt.Println(address2)

	address1 := Address{"Subang", "Jabar", "indonesia"}
	address2 := &address1 //! mengubah value yang ada di address1 dan ini merupakan pointer
	address2.City = "bandung"

	address3 := &address1

	*address2 = Address{"Malang", "Jatim", "indonesia"} //! mengubah semua data atau value yang ada pada variable sebelumnya

	address4 := new(Address) //! menggunakan function new jika ingin menampilkan struct kosong atau data kosong
	address4.City = "bandung"

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println(address4)

	var alamat = Address{
		City:     "Subang",
		Province: "Jabar",
		Country:  "",
	}

	var alamatPointer *Address = &alamat
	ChangeCountryToIndonesia(alamatPointer)
	fmt.Println(alamatPointer)
}
