package main

import (
	"fmt"
	"sort"
)

func main() {
	words := []string{"cat", "dog", "tac", "god", "act", "fan"}
	anagrams, exists := findAnagrams(words)
	if exists {
		fmt.Printf("Anagrams found: %v\n", anagrams)
	} else {
		fmt.Println("No anagrams found")
	}
}

// findAnagrams returns a 2d slice of strings of words that are grouped by anagram
func findAnagrams(words []string) ([][]string, bool) {

	// use sorted string of the word as key for the map
	anagrams := make(map[string][]string)
	for _, word := range words {
		sortedWord := sortString(word)
		anagrams[sortedWord] = append(anagrams[sortedWord], word)
	}

	var result [][]string
	for _, anagramGroup := range anagrams {
		if len(anagramGroup) > 1 {
			result = append(result, anagramGroup)
		}
	}

	return result, len(result) > 0
}

func sortString(s string) string {
	runes := []rune(s)

	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	return string(runes)
}
