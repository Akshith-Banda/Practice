package main

import "fmt"

// https://www.hackerrank.com/challenges/larrys-array/problem?isFullScreen=true
// core logic
func larrysArray(A []int32) string { //1, 3, 4, 2
	// Write your code here

	var inversion int32 //3, 1, 2
	for i, n := range A {
		fmt.Println("n : ", n)
		if i == 0 {
			continue
		}
		for _, m := range A[:i] {
			fmt.Print("  m : ", m)
			if m > n {
				fmt.Printf("\tm > n : %t\n", m > n)
				inversion++
			}
			fmt.Println()
		}
	}
	if inversion%2 == 0 {
		return "YES"
	}
	return "NO"
}
