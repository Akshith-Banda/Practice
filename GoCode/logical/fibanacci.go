package main

import (
	"fmt"
)

func main() {
	fibanacci(10)
	recursionFib(10)
	fmt.Println("*************RECFIBANACCI***********")
	fmt.Println(series)
	fmt.Println("********************************")
}

func fibanacci(n int) {
	list := []int{0, 1}

	for i := 3; i <= n; i++ {
		sum := list[len(list)-1] + list[len(list)-2]
		list = append(list, sum)
	}

	fmt.Println("*************FIBANACCI***********")
	fmt.Println(list)
	fmt.Println("********************************")
}

var series = make(map[int]int)

func recursionFib(n int) int {

	if n, ok := series[n]; ok {
		return n
	}

	if n == 0 || n == 1 {
		series[n] = n
		return n
	}

	sum := recursionFib(n-1) + recursionFib(n-2)

	series[n] = sum
	return sum
}
