package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/christopherfujino/go-pass/diceware"
)

func main() {
	var n = getN(os.Args)
	password := diceware.GeneratePassword(n)
	var entropy = diceware.CalculateEntropy(n)
	fmt.Printf(
		"%s\n\nThe length of your password is %d with %d bits of entropy.\n",
		password,
		len(password),
		entropy,
	)
}

func getN(args []string) (int) {
	if len(args) > 1 {
		var nString = args[1]
		n, err := strconv.Atoi(nString)
		if err != nil {
			panic(err)
		}
		if n > 0 && n < 64 {
			return n
		}
		panic(fmt.Sprintf("Invalid length %d", n))
	}
	return 4
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
