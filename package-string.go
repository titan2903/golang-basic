package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("titanio yudista", "titanio"))
	fmt.Println(strings.Contains("titanio yudista", "test"))
	fmt.Println(strings.Split("titanio yudista", " "))
	fmt.Println(strings.ToLower("TITANIO YUDISTA"))
	fmt.Println(strings.ToUpper("Titani0 Yudista"))
	fmt.Println(strings.Trim("   Titanio Yudista      ", " "))
	fmt.Println(strings.ReplaceAll("joko titan titan", "titan", "eko"))
}