package main

import "fmt"

func sumAll(numbers ...int) int { //! hanya bisa di tempatkan di final parameter atau di belakang
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func main() {
	total := sumAll(10, 30, 90, 10)
	fmt.Println(total)

	slice := []int{10, 10, 10, 40}
	total = sumAll(slice...)
	fmt.Println(total)

}
