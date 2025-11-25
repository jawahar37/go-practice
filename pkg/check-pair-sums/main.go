package main

import "fmt"

func main() {
	arr := []int{-1, 3, 4, 5, 6, 7, -7, -11}
	target := 5

	x, y, exists := checkPairSums(arr, target)
	if exists {
		fmt.Printf("Pair found: (%d, %d): %d + %d = %d\n", x, y, arr[x], arr[y], target)
	} else {
		fmt.Println("Pair not found")
	}
}

func checkPairSums(arr []int, target int) (int, int, bool) {
	complements := make(map[int]int)

	for i, num := range arr {
		complement := target - num
		if j, ok := complements[complement]; ok {
			return i, j, true
		}
		complements[num] = i
	}

	return 0, 0, false
}
