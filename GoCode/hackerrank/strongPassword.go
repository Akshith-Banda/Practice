package main

import "regexp"

// https://www.hackerrank.com/challenges/strong-password/problem?isFullScreen=true
func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var expList = []string{`([a-z]+)`, `([A-Z]+)`, `([0-9]+)`, `([!@#$%^&*()+-]+)`}

	var times int32
	for _, exp := range expList {
		re := regexp.MustCompile(exp)
		if !re.Match([]byte(password)) {
			times++
		}
	}

	if n+times < 6 {
		toAdd := 6 - (n + times)
		times += toAdd
	}
	return times
}
