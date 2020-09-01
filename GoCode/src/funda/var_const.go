package main

import (
	"fmt"
	"reflect"
)

//package level decleration of variables.
//go can type infer.
//If declaring inside a func use := for type inference.
var (
	name, course, module = "Akshith", "Docker", 3.2
)

func main() {
	fmt.Println("type of name is ", reflect.TypeOf(name))

	a := 10.869
	b := 3
	c := int(a) + b // can't add float and int, so convert into int.
	fmt.Println("value of c is ", c)

}
