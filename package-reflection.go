package main

import (
	"fmt"
	"reflect"
)


type Sample struct {
	Name string `required:"true" max:"10"`
}

type ContohLagi struct {
	Name string `required:"true"`
	Description string `required:"true"`
}

func IsValid(data interface{}) bool { //! Validation menggunakan validation
	t := reflect.TypeOf(data)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)

		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"titan"}

	var sempleType reflect.Type = reflect.TypeOf(sample)
	
	fmt.Println(sempleType.NumField())
	fmt.Println(sempleType.Field(0).Name)
	fmt.Println(sempleType.Field(0).Tag.Get("required"))
	fmt.Println(sempleType.Field(0).Tag.Get("max"))
	// sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"titan", "test"}
	fmt.Println(IsValid(contoh))
}

/*
? cocok di gunakan kuntuk validasi jika ingin membuat framework atau validation
*/