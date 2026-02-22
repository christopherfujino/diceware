package main

import (
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/christopherfujino/diceware/diceware"
)

func getAverageGuessCountString(entropy int) string {
	var entropyFloat = (math.Pow(2, float64(entropy))) / 2
	var rec func(int) string = nil
	rec = func(i int) string {
		if i < 1000 {
			return fmt.Sprintf("%03d", i)
		}
		var prefix = i / 1000
		var suffix = i - (prefix * 1000)
		var prefixString = ""
		if prefix < 1000 {
			prefixString = fmt.Sprintf("%d", prefix)
		} else {
			prefixString = rec(prefix)
		}
		return fmt.Sprintf("%s,%s", prefixString, rec(suffix))
	}
	return rec(int(entropyFloat))
}

func main() {
	var n = getN(os.Args)
	password := diceware.GeneratePassword(n)
	var entropy = diceware.CalculateEntropy(n)
	fmt.Printf(
		"%s\n\nThe length of your password is %d bytes with %d bits of entropy (on average %s guesses).\n",
		password,
		len(password),
		entropy,
		getAverageGuessCountString(entropy),
	)
}

func getN(args []string) int {
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
