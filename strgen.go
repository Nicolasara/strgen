package strgen

import (
	"math/rand"
	"strings"
)

var Numbers = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
var LowercaseLetters = []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
var UppercaseLetters = []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}

func Generate(strLen int) string {
	possibleChars := getPossibleChars(true, true, true)
	strChars := []string{}
	for i := 0; i < strLen; i++ {
		char := possibleChars[rand.Intn(len(possibleChars))]
		strChars = append(strChars, char)
	}
	return strings.Join(strChars, "")
}

func getPossibleChars(numbers bool, lowercase bool, uppercase bool) []string {
	characters := []string{}
	if numbers {
		characters = append(characters, Numbers...)
	}
	if lowercase {
		characters = append(characters, LowercaseLetters...)
	}
	if uppercase {
		characters = append(characters, UppercaseLetters...)
	}
	return characters
}
