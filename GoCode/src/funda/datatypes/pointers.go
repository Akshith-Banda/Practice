package main

import "fmt"

func main() {
	module := 3.4
	ptr := &module

	fmt.Println("pointer to module is ", ptr)
	fmt.Println("value in pointer is ", *ptr)

	//args to funcs are pass by value. i.e copy is made and sent to func.
	fmt.Println("\nvalue in module is ", module)
	changeModule(module)
	fmt.Println("value in module is ", module)

	//pass by reference
	fmt.Println("\nvalue in module is ", module)
	changeModRef(&module)
	fmt.Println("value in module is ", module)
}
func changeModule(mod float64) float64 {
	mod = 3.2
	fmt.Println("value in module is ", mod)
	return mod
}
func changeModRef(mod *float64) float64 {
	*mod = 2.9
	fmt.Println("value in module is ", *mod)
	return *mod
}
