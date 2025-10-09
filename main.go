package main

import (
	"fmt"
	"math"

	"github.com/christopherfujino/go-pass/diceware"
)

const maxNum = (6 * 6 * 6 * 6 * 6) - 1

func main() {
	const n = 4
	password := diceware.GeneratePassword(n)
	var entropy = diceware.CalculateEntropy(math.Pow(float64(maxNum + 1), float64(n)))
	fmt.Printf(
		"Your password is \"%s\"\nThe length of your password is %d with %d bits of entropy.\n",
		password,
		len(password),
		entropy,
	)
}

//type ParseState int
//
//const (
//	ParseNum ParseState = iota
//	ParseWord
//)
//
//func Words() WordList {
//	const path = "./wordlist.txt"
//	var words = make(map[int]string, 8000)
//	var state = ParseNum
//
//	bytes, err := os.ReadFile(path)
//	if err != nil {
//		panic(err)
//	}
//	var buffer = strings.Builder{}
//	var num = 0
//	var ch byte
//	// we assume wordlist is ASCII
//	for i := range len(bytes) {
//		ch = bytes[i]
//
//		switch state {
//		case ParseNum:
//			if ch == '\t' {
//				state = ParseWord
//			} else {
//				digit, err := strconv.Atoi(string(ch))
//				if err != nil {
//					panic(err)
//				}
//				// Base 6 digits
//				num = num*6 + (digit - 1)
//			}
//		case ParseWord:
//			if ch == '\n' {
//				words[num] = buffer.String()
//				buffer.Reset()
//				num = 0
//				state = ParseNum
//			} else {
//				buffer.WriteByte(ch)
//			}
//		}
//	}
//	return words
//}
