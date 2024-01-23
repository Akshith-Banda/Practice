package main

func palindrome(text string) bool {
	isPalindrome := true
	var limit int
	if len(text)%2 != 0 {
		limit = (len(text) - 1) / 2
	} else {
		limit = len(text) / 2
	}
	for i := 0; i < limit; i++ {
		if text[i] != text[len(text)-(i+1)] {
			isPalindrome = false
			break
		}
	}
	return isPalindrome
}
