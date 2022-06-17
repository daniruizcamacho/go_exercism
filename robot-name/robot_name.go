package robotname

import (
	"errors"
	"math/rand"
	"strconv"
	"time"
)

// Define the Robot type here.
type Robot struct {
	name string
}

var letters = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
var usedNames map[string]bool

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}

	if usedNames == nil {
		usedNames = make(map[string]bool)
	}

	if len(usedNames) == 26*26*10*10*10 {
		return "", errors.New("No possible names")
	}

	newName := ""
	for newName == "" {
		generatedName := generateRandomLetter() + generateRandomLetter() + generateRandomNumber() + generateRandomNumber() + generateRandomNumber()
		if usedNames[generatedName] == false {
			newName = generatedName
		}
	}

	usedNames[newName] = true
	r.name = newName

	return newName, nil
}

func generateRandomLetter() string {
	rand.Seed(time.Now().UnixNano())

	return string(letters[rand.Intn(len(letters))])
}

func generateRandomNumber() string {
	rand.Seed(time.Now().UnixNano())

	return strconv.Itoa(rand.Intn(9))
}

func (r *Robot) Reset() {
	r.name = ""
}
