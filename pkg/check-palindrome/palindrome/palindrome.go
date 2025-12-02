package palindrome

func IsPalindrome(word string) bool {
	n := len(word)

	for i, j := 0, n-1; i <= j; {
		if word[i] != word[j] {
			return false
		}
		i += 1
		j -= 1
	}
	return true
}
