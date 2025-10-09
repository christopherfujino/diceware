package diceware

import (
	"crypto/rand"
	"fmt"
	"strings"
)

func GeneratePassword(n int) string {
	var buffer = strings.Builder{}
	for i := range n {
		var word = getWord()
		if i%2 == 1 {
			word = strings.ToUpper(word)
		}
		buffer.WriteString(word)
	}
	return buffer.String()
}

func getWord() string {
	return wordList[GetRand()]
}

func GetRand() int {
	var buffer = make([]byte, 5)
	n, err := rand.Read(buffer)
	if err != nil {
		// This should be unreachable
		panic(err)
	}
	if n != 5 {
		panic(
			fmt.Sprintf(
				"crypto/rand.Read() generated the wrong number of bytes; expected 5, got %d",
				n,
			),
		)
	}
	var newRand int
	for i := range 5 {
		var integer = int(buffer[i]) % 6
		if integer < 0 || integer >= 6 {
			panic(fmt.Sprintf("Unreachable (%d)", integer))
		}
		newRand = newRand*6 + integer
	}
	return newRand
}
