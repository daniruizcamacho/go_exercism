package grains

import "errors"

func Square(number int) (uint64, error) {
	if number <= 0 || number > 64 {
		return 0, errors.New("No valid number")
	}

	var result uint64 = 1

	return result << (number - 1), nil
}

func Total() uint64 {
	var total uint64 = 0
	for i := 1; i <= 64; i++ {
		square, _ := Square(i)
		total += square
	}

	return total
}
