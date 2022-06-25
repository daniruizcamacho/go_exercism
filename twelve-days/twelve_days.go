package twelve

import (
	"fmt"
	"strings"
)

var days = map[int]string{
	1:  "first",
	2:  "second",
	3:  "third",
	4:  "fourth",
	5:  "fifth",
	6:  "sixth",
	7:  "seventh",
	8:  "eighth",
	9:  "ninth",
	10: "tenth",
	11: "eleventh",
	12: "twelfth",
}

var items = map[int]string{
	1:  "a Partridge in a Pear Tree",
	2:  "two Turtle Doves",
	3:  "three French Hens",
	4:  "four Calling Birds",
	5:  "five Gold Rings",
	6:  "six Geese-a-Laying",
	7:  "seven Swans-a-Swimming",
	8:  "eight Maids-a-Milking",
	9:  "nine Ladies Dancing",
	10: "ten Lords-a-Leaping",
	11: "eleven Pipers Piping",
	12: "twelve Drummers Drumming",
}

func Verse(i int) string {
	return fmt.Sprintf("On the %s day of Christmas my true love gave to me: %s.", days[i], itemsList(i))
}

func itemsList(i int) string {
	if i == 1 {
		return items[1]
	}

	list := ""
	for j := i; j > 1; j-- {
		list += items[j] + ", "
	}

	list += "and " + items[1]

	return list
}

func Song() string {
	verses := []string{}
	for i := 1; i <= 12; i++ {
		verses = append(verses, Verse(i))
	}

	return strings.Join(verses, "\n")
}
