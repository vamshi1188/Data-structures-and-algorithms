package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s := "0P"

	example1(s)

	fmt.Println(palindromecheck(s, 0))
}
func palindromecheck(s string, i int) bool {

	if i >= len(s)/2 {
		return true
	}

	if s[i] != s[len(s)-i-1] {
		return false
	}

	return palindromecheck(s, i+1)
}

func example1(s string) {
	var reversedoutput []rune

	var output []rune
	for _, name := range s {

		if unicode.IsLetter(name) || unicode.IsDigit(name) {
			output = append(output, name)

		}

	}

	// outputrune := []rune(output)

	for i := len(output) - 1; i >= 0; i-- {

		reversedoutput = append(reversedoutput, output[i])
	}

	output1 := string(output)
	reversedoutput2 := string(reversedoutput)

	output1 = strings.ToLower(output1)
	reversedoutput2 = strings.ToLower(reversedoutput2)
	ans := output1 == reversedoutput2

	fmt.Println(ans)
}
