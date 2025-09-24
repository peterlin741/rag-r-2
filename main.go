package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// number between 1 and 3
	randomNumber := rand.Intn(3) + 1
	item := GoogleRagSimpsonsSpiderPig(randomNumber)
	fmt.Println(item)
}

// GoogleRagSimpsonsSpiderPig returns a vegetable based on the number provided
func GoogleRagSimpsonsSpiderPig(i int) string {
	if i == 0 {
		return "brocolli"
	} else if i == 1 {
		return "spinach"
	} else {
		return "cabbage"
	}
}
