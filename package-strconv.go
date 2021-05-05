package main

import (
	"fmt"
	"strconv"
)


func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	}else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)

	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	value := strconv.FormatInt(1000000, 8)
	fmt.Println(value)

	valueInt, err := strconv.Atoi("200000") //! Cara paling mudah conversi ke int
	if err == nil {
		fmt.Println(valueInt)
	} else {
		fmt.Println(err.Error())
	}

}

/*
!	conversi data menjadi data yang kita inginkan
!    strconv.ParseInt(angkanya, base, bit)
*/