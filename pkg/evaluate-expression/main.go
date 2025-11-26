package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	str := "nineonepluseighttwominusthreeplusfour"
	fmt.Printf("Original: %s\n", str)

	exprTokens := convertStringToExpression(str)
	fmt.Printf("Expression Tokens: %v\n", exprTokens)

	result := evaluateExpression(exprTokens)
	fmt.Printf("Result: %d\n", result)
}

// convertStringToExpression returns the numbers and operators as a slice of string tokens
// "[91 + 82 - 3 + 4]" is returned for "nineonepluseighttwominusthreeplusfour"
func convertStringToExpression(str string) []string {
	tokens := map[string]string{
		"zero":  "0",
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
		"plus":  "+",
		"minus": "-",
	}

	isOperator := func(s string) bool {
		return s == "+" || s == "-"
	}

	var result []string
	var currentNumber strings.Builder
	i := 0

	for i < len(str) {
		found := false
		for word, symbol := range tokens {
			if strings.HasPrefix(str[i:], word) {
				if !isOperator(symbol) {
					currentNumber.WriteString(symbol)
				} else {
					if currentNumber.Len() > 0 { // Add any pending number before the operator
						result = append(result, currentNumber.String())
						currentNumber.Reset()
					}
					result = append(result, symbol)
				}

				i += len(word)
				found = true
				break
			}
		}
		if !found {
			// Add any pending number first
			if currentNumber.Len() > 0 {
				result = append(result, currentNumber.String())
				currentNumber.Reset()
			}
			i++
			fmt.Printf("Invalid character skipped: %c\n", str[i-1])
		}
	}

	// Add any remaining numbers
	if currentNumber.Len() > 0 {
		result = append(result, currentNumber.String())
	}

	return result
}

// evaluateExpression returns the computed value of the expression
// 174 is returned for [91 + 82 - 3 + 4]
func evaluateExpression(tokens []string) int {
	if len(tokens) == 0 {
		return 0
	}

	// The first token must be a number.
	result, err := strconv.Atoi(tokens[0])
	if err != nil {
		fmt.Printf("Invalid number: %s\n", tokens[0])
		return 0
	}

	// Loop through the rest of the tokens as operator-number pairs.
	for i := 1; i < len(tokens); i += 2 {
		operator := tokens[i]
		operandStr := tokens[i+1]

		operand, err := strconv.Atoi(operandStr)
		if err != nil {
			fmt.Printf("Invalid number: %s\n", operandStr)
			return 0
		}

		if operator == "+" {
			result += operand
		} else if operator == "-" {
			result -= operand
		}
	}
	return result
}
