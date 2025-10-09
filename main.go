package main

import (
	"crypto/rand"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

const maxNum = (6 * 6 * 6 * 6 * 6) - 1

func main() {
	wordList := Words()

	password, entropy := GeneratePassword(wordList, 4)
	fmt.Printf(
		"Your password is \"%s\"\nThe length of your password is %d with %d bits of entropy.\n",
		password,
		len(password),
		entropy,
	)
}

func GeneratePassword(words WordList, n int) (string, int) {
	var buffer = strings.Builder{}
	for i := range 4 {
		var word = words.GetWord()
		if i%2 == 1 {
			word = strings.ToUpper(word)
		}
		buffer.WriteString(word)
	}
	return buffer.String(), CalculateEntropy(math.Pow(float64(maxNum+1), float64(n)))
}

func CalculateEntropy(permutations float64) (numBits int) {
	var log2 = math.Ceil(math.Log2(permutations))
	return int(log2)
}

func (w WordList) GetWord() string {
	return w[GetRand()]
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

type ParseState int

const (
	ParseNum ParseState = iota
	ParseWord
)

type WordList map[int]string

func Words() WordList {
	const path = "./wordlist.txt"
	var words = make(map[int]string, 8000)
	var state = ParseNum

	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var buffer = strings.Builder{}
	var num = 0
	var ch byte
	// we assume wordlist is ASCII
	for i := range len(bytes) {
		ch = bytes[i]

		switch state {
		case ParseNum:
			if ch == '\t' {
				state = ParseWord
			} else {
				digit, err := strconv.Atoi(string(ch))
				if err != nil {
					panic(err)
				}
				// Base 6 digits
				num = num*6 + (digit - 1)
			}
		case ParseWord:
			if ch == '\n' {
				words[num] = buffer.String()
				buffer.Reset()
				num = 0
				state = ParseNum
			} else {
				buffer.WriteByte(ch)
			}
		}
	}
	return words
}
