package main

import (
	"fmt"
	"os"
	"tutorial/palindrome"
)

func main() {
	var word string
	if len(os.Args) == 2 {
		word = os.Args[1]
	} else {
		word = "racecar"
	}

	isPalindrome := palindrome.IsPalindrome(word)

	if isPalindrome {
		fmt.Printf("%s is a palindrome\n", word)
	} else {
		fmt.Printf("%s is not a palindrome\n", word)
	}

}
