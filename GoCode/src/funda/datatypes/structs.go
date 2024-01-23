package main

import "fmt"

func main() {

	type courseMeta struct {
		Author string
		Level  string
		Rating float64
	}

	//var DockerDeepDive courseMeta
	//DockerDeepDive := new(courseMeta)
	DockerDeepDive := courseMeta{
		Author: "Akshith",
		Level:  "Beginner",
		Rating: 4.7,
	}

	fmt.Println(DockerDeepDive)
	fmt.Println("\nAuthor is ", DockerDeepDive.Author)
}
