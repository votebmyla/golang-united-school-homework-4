package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
// For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {
	input = strings.ReplaceAll(input, " ", "")
	if len(input) == 0 {
		return "", fmt.Errorf("%w", errorEmptyInput)
	}

	var num1, num2 string

	lastPlus := strings.LastIndex(input, "+")
	lastMinus := strings.LastIndex(input, "-")

	if lastPlus > lastMinus {
		num1 = input[:lastPlus]
		num2 = input[lastPlus:]

		if strings.Contains(num1[1:], "+") || strings.Contains(num1[1:], "-") {
			return "", fmt.Errorf("%w", errorNotTwoOperands)
		}
	}

	if lastMinus > lastPlus {
		num1 = input[:lastMinus]
		num2 = input[lastMinus:]
	}

	if len(num1) == 0 {
		return "", fmt.Errorf("%w", errorNotTwoOperands)
	}

	x, errx := strconv.Atoi(num1)
	if errx != nil {
		return "", fmt.Errorf("%w", errx)
	}

	y, erry := strconv.Atoi(num2)
	if erry != nil {
		return "", fmt.Errorf("%w", erry)
	}

	output = strconv.Itoa(x + y)
	return output, nil
}
