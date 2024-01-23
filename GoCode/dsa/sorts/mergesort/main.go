package main

import "fmt"

func main() {
	fmt.Println(mergesort([]int{8, 5, 7, 3, 0, 8}))
}

func mergesort(items []int) []int {
	if len(items) < 2 {
		return items
	}

	first := mergesort(items[:len(items)/2])
	second := mergesort(items[len(items)/2:])

	return merge(first, second)
}

func merge(first, second []int) []int {
	final := []int{}
	i := 0
	j := 0
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			final = append(final, first[i])
			i++
		} else {
			final = append(final, second[j])
			j++
		}
	}

	for ; i < len(first); i++ {
		final = append(final, first[i])
	}
	for ; j < len(second); j++ {
		final = append(final, second[j])
	}
	return final
}
