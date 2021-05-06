package main

import (
	"fmt"
	"sort"
)


type User struct {
	Name string
	Age  int
}

type UserSlice []User //! membuat contract interface function untuk sorting

func (value UserSlice) Len() int { //! function mengetahui jumlah data
	return len(value)
}

func (value UserSlice) Less(i, j int) bool { //! function membandingkan
	return value[i].Age > value[j].Age 
}

func (value UserSlice) Swap(i, j int) { //! function swap
	value[i], value[j] = value[j], value[i]
}

func main() {
	users := []User{
		{"eko", 30},
		{"andi", 20},
		{"rudi", 40},
		{"akbar", 31},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}