package main

import (
	"fmt"

	"github.com/codecamp-prem/stringJoins/prose"
)

func main() {
	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a bull prize"}
	fmt.Println("A photo of", prose.JoinWithCommas(phrases))
}
