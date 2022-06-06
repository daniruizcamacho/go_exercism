package luhn

import (
	"strconv"
	"strings"
)

func Valid(id string) bool {
	id = strings.ReplaceAll(id, " ", "")
	if len(id) < 2 {
		return false
	}

	lenghtIsPair := isPair(len(id))

	sum := 0
	for index, number := range id {
		value, err := strconv.Atoi(string(number))
		if err != nil {
			return false
		}

		valueToSum := value
		if lenghtIsPair && isPair(index) {
			valueToSum = duplicateNumber(value)
		}

		if !lenghtIsPair && !isPair(index) {
			valueToSum = duplicateNumber(value)
		}

		sum += valueToSum
	}

	return sum%10 == 0
}

func isPair(value int) bool {
	return value%2 == 0
}

func duplicateNumber(value int) int {
	newValue := value * 2

	if newValue > 9 {
		return newValue - 9
	}

	return newValue
}
