package main

import (
	"fmt"
	"math/rand"
	"time"
)

// if-statement

func main() {
	// if-statement
	if first, second := 38, 98; first < second {
		fmt.Println("first is higher")
	} else if first > second {
		fmt.Println("second is higher")
	} else {
		fmt.Println("both are equal")
	}
	//switch
	switch "docker" {
	case "linux":
		fmt.Println("its linux!")
	case "docker":
		fmt.Println("its docker")
	case "windows":
		fmt.Println("its windows")
	default:
		fmt.Println("none matches")
	}
	//switch <simple-statement>; expression
	switch tmpNumber := random(); tmpNumber {
	case 0, 2, 4, 6, 8:
		fmt.Println("even")
	case 1, 3, 5, 7, 9:
		fmt.Println("odd")
	}
}
func random() int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(10)
}
