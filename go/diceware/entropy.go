package diceware

import (
	"math"
)

const maxNum = (6 * 6 * 6 * 6 * 6) - 1

func CalculateEntropy(wordCount int) (numBits int) {
	var permutations = math.Pow(maxNum + 1, float64(wordCount))
	var log2 = math.Ceil(math.Log2(permutations))
	return int(log2)
}
