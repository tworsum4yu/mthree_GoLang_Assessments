package code

import (
	"fmt"
	"strings"
)

func countWords(s string) map[string]int {

	list := strings.Fields(s)

	count := make(map[string]int)

	for _, word := range list {
		count[word]++
	}

	return count
}

// Do not change the code in the main function
func RunQ3() {
	s := "go is fun and go is fast"
	fmt.Println(countWords(s))
	t := "this is a test this is only a test"
	fmt.Println(countWords(t))
}
