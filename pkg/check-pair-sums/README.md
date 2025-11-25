# Check Pair Sums

`check-pair-sums` is a simple Go application that determines whether a pair of numbers in an integer slice sums up to a given target value.

## Description

The program defines a function `checkPairSums` that takes a slice of integers and a target integer. It efficiently finds if any two numbers in the slice add up to the target. If a pair is found, it returns the indices of the two numbers and `true`. Otherwise, it returns `false`.

The `main` function demonstrates how to use `checkPairSums` with a sample array and target value, and it prints the result.

## Algorithm

The `checkPairSums` function uses a map to store the numbers encountered so far and their indices. For each number in the input slice, it calculates the required `complement` to reach the target. It then checks if this complement already exists in the map.

- If the complement exists, a pair is found.
- If not, the current number and its index are added to the map.

This approach has a time complexity of O(n) because it only requires a single pass through the slice.

## How to Run

Navigate to the application's directory and run the `main.go` file.

```sh
go run main.go
```

### Example Output

```
Pair found: (4, 0): 6 + -1 = 5
```
