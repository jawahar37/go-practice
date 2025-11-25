# Find Anagrams

`find-anagrams` is a simple Go application that identifies and groups anagrams from a given list of words.

## Description

The program defines a function `findAnagrams` that takes a slice of strings and groups them based on whether they are anagrams of each other. It returns a slice of string slices, where each inner slice contains a group of anagrams.

The `main` function provides an example of how to use `findAnagrams` with a sample list of words and prints the resulting anagram groups.

## Algorithm

The `findAnagrams` function employs an efficient method to group anagrams. It uses a map where keys are the sorted versions of words and values are slices of the original words that correspond to that sorted key.

For each word in the input slice:
1.  The characters of the word are sorted alphabetically. This sorted string serves as a unique identifier for a group of anagrams.
2.  The original word is appended to the list of words associated with its sorted version in the map.

After processing all words, the function iterates through the map. Any group containing more than one word is considered a valid group of anagrams and is added to the final result.

This approach has a time complexity of O(n * k log k), where `n` is the number of words and `k` is the maximum length of a word, making it efficient for grouping anagrams.

## How to Run

To run the application, navigate to the program's directory and execute the `main.go` file.

```sh
go run main.go
```

### Example Output

The output will show the groups of anagrams found in the sample list. The order of the groups and words within the groups may vary.

```
Anagrams found: [[cat tac act] [dog god]]
```
