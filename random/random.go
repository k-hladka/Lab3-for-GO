package random

import (
	"math"
	"time"
)

type paramMap map[string]int64

func Random(p paramMap) []int64 {
	// Лінійний конгруентний метод
	// X = (aX + c) mod m
	var (
		i       int64 = 0
		x             = p["min"]
		numbers       = make([]int64, p["length"])
	)

	p["m"] = int64(math.Pow(2, float64(p["m"])))

	for ; i < p["length"]; i++ {
		x = (p["a"]*x + p["c"]) % p["m"] % p["max"] % int64(time.Now().Second())

		numbers[i] = x
	}
	return numbers
}
