package main

import "fmt"

// https://www.hackerrank.com/challenges/crush/problem?isFullScreen=true
// core logic
func arrayManipulation(n int32, queries [][]int32) int64 {
	// Write your code here
	var arr = make([]int64, n+2)
	var max int64

	for _, query := range queries {

		a := query[0]
		b := query[1]
		k := query[2]

		arr[a] += int64(k)
		arr[b+1] -= int64(k)
		fmt.Println(arr)
	}
	var sum int64
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		if sum > max {
			max = sum
		}
	}
	return int64(max)
}
