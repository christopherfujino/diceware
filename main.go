package main

import (
	"fmt"
	//"io"
	"os"
	"strings"
)

func main() {
	fmt.Println("Hello, world!")
	words := Words()

	for _, word := range words {
		fmt.Println(word)
	}
}

func Words() []string {
	const path = "./wordlist.txt"
	var words = make([]string, 0, 8000)
	bytes, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	var buffer = strings.Builder{}
	var ch byte
	// we assume wordlist is ASCII
	for i := range len(bytes) {
		ch = bytes[i]
		if ch == '\n' {
			words = append(words, buffer.String())
			buffer.Reset()
		} else {
			buffer.WriteByte(ch)
		}
	}
	return words
}
