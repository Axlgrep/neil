package main

import (
	"fmt"
	"unicode/utf8"
)

func IsPalindrome(word string) bool {
	if utf8.RuneCountInString(word) <= 1 {
		return true
	}
	first, sizeOfFirst := utf8.DecodeRuneInString(word)
	last, sizeofLast := utf8.DecodeLastRuneInString(word)
	if first != last {
		return false
	}
	return IsPalindrome(word[sizeOfFirst : len(word)-sizeofLast])
}

func main() {
	fmt.Println("XxxxaxxxX is palindrome: ", IsPalindrome("XxxxaxxxX"))
}
