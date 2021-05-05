package main

import (
	"belajar_golang_dasar/database"
	"fmt"
)


func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}

/*
	_ blank idintifier jika tidak ingin di gunakan
*/