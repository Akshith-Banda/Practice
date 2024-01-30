package main

// https://www.hackerrank.com/challenges/lonely-integer/problem?isFullScreen=true
// core logic
func lonelyinteger(a []int32) int32 {
	// Write your code here
	var result int32
	for idx, num := range a {
		if idx == 0 {
			result = num
			continue
		}
		result = result ^ num //xor all the int's and will return which is not repeated
	}
	return result
}
