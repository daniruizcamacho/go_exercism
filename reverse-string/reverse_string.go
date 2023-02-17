package reverse

func Reverse(input string) string {
	result := ""
	inputRune := []rune(input)
	for i := len(inputRune) - 1; i >= 0; i-- {
		result += string(inputRune[i])
	}

	return result
}
