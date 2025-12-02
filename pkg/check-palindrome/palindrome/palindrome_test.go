package palindrome_test

import (
	"testing"
	"tutorial/palindrome"
)

func TestIsPalindrome(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		word string
		want bool
	}{
		{"empty string", "", true},
		{"single character", "a", true},
		{"two identical characters", "aa", true},
		{"two different characters", "ab", false},
		{"odd length palindrome", "racecar", true},
		{"even length palindrome", "madam", true},
		{"odd length non-palindrome", "hello", false},
		{"even length non-palindrome", "world", false},
		{"palindrome with spaces", "race car", false}, // Asserting spaces make it not a palindrome
		{"palindrome with mixed case", "Racecar", false},
		{"palindrome with numbers", "12321", true},
		{"non-palindrome with numbers", "12345", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := palindrome.IsPalindrome(tt.word)
			if got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
