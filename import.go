package main

import (
	"belajar_golang_dasar/helper"
	"fmt"
)


func main() {
	helper.SayHallo("titan")
	fmt.Println(helper.Application)
	// fmt.Println(helper.version) Error unexported
}