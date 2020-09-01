package main

import "fmt"

func main() {
	leagueTitles := make(map[string]int)
	leagueTitles["sunderland"] = 6
	leagueTitles["Newcastle"] = 4

	recentHead2Head := map[string]int{
		"sunderland": 6,
		"Newcastle":  4,
	}

	fmt.Println(leagueTitles, "\n", recentHead2Head)

	testMap := map[string]int{"A": 1, "B": 2, "C": 3, "D": 4}
	for key, value := range testMap {
		fmt.Println(key, ":", value)
	}

	delete(testMap, "D")
	fmt.Println(testMap)
}
