package main

import (
	"fmt"
	"strings"
)

func main() {
	var word1 = []string{"abcd", "d", "defg"}
	var word2 = []string{"abcddefg"}

	fmt.Println(strings.Join(word1, "") == strings.Join(word2, ""))
}
