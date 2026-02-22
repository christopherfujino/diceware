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
