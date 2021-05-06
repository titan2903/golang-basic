package main

import (
	"fmt"
	"regexp"
)


func main() {
	var regex *regexp.Regexp = regexp.MustCompilePOSIX("e([a-z])o")

	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	search := regex.FindAllString("eko eka edo eki eyo", -1) //! semua di ambil sesuai dengan regex nya
	fmt.Println(search)
}