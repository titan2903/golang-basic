package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	// months[5] = "Diubah"
	// fmt.Println(slice1)

	slice1[0] = "Mei lagi"
	fmt.Println(len(months))

	var slice2 = months[11:]

	var slice3 = append(slice2, "EKo") //! menambahkan value pada slice
	slice3[1] = "Bukan Desember"

	fmt.Println(slice2)
	fmt.Println(slice3)
	fmt.Println(months)

	newSlice := make([]string, 2, 5)

	newSlice[0] = "Titanio"
	newSlice[1] = "Yudista"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)
}
