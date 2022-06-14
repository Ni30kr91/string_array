package main

import "fmt"

func main() {
	var word1 = []string{"abcd", "d", "defg"}
	var word2 = []string{"abcddefg"}

	Equal(word1, word2)
}
func Equal(word1, word2 []string) {
	var a string
	var b string
	for i := 0; i < len(word1); i++ {
		a = a + word1[i]
		fmt.Println(a)
	}
	for j := 0; j < len(word2); j++ {
		b = b + word2[j]
		fmt.Println(b)
	}
	fmt.Println(a == b)

}
