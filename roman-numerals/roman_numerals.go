package romannumerals

import (
	"fmt"
)

var romanNum = map[int]string{
	1000: "M", 2000: "MM", 3000: "MMM",
	100: "C", 200: "CC", 300: "CCC", 400: "CD", 500: "D", 600: "DC", 700: "DCC", 800: "DCCC", 900: "CM",
	10: "X", 20: "XX", 30: "XXX", 40: "XL", 50: "L", 60: "LX", 70: "LXX", 80: "LXXX", 90: "XC",
	1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX",
}

// ToRomanNumeral converts an integer to roman numerals.
// The integer must be < 4000 and > 0
func ToRomanNumeral(num int) (roman string, err error) {
	// roman numerals can only express numbers >= 1
	if num <= 0 || num >= 4000 {
		err = fmt.Errorf("only numbers greater than 0 and less than 4000 can be converted into roman numerals")
		return
	}
	// thousands
	k := num / 1000
	roman += romanNum[k*1000]
	// hundreds
	c := (num / 100) % 10
	roman += romanNum[c*100]
	// tens
	d := (num / 10) % 10
	roman += romanNum[d*10]
	// units
	u := (num / 1) % 10
	roman += romanNum[u*1]
	return
}
