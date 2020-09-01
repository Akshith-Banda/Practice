package main

import (
	"fmt"
)

func main() {

	coursesInProg := []string{"docker", "Go", "kubernetes"}
	coursesCompleted := []string{"Go", "python", "kube"}
breakpoint:
	for _, i := range coursesInProg {
		for _, j := range coursesCompleted {
			if i == j {
				fmt.Println(i, "both match")
				break breakpoint
			}
		}
	}
	fmt.Println("out of both loops")

	bestFinish := bestLeagueFinish(10, 12, 7, 15, 11)
	fmt.Println("\nbest finish ", bestFinish)
}

//variadic functions
func bestLeagueFinish(finish ...int) int {
	best := finish[0]
	for _, i := range finish {
		if i < best {
			best = i
		}
	}
	return best
}
