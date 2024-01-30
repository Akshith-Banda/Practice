package main

import "strings"

// https://www.hackerrank.com/challenges/reduced-string/problem?isFullScreen=true
func superReducedString(s string) string { // baabaddfef
	// Write your code here
	for i := 0; i < len(s); i++ {
		if strings.Contains(s, string(s[i])+string(s[i])) {
			s = strings.ReplaceAll(s, string(s[i])+string(s[i]), "")
			i = -1
		}
	}
	if s == "" {
		return "Empty String"
	}
	return s
}
