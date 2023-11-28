package random

import (
	"math"
	"time"
)

type paramMapFloat map[string]float64

func RandomFloat(p paramMapFloat) []float64 {
	// Лінійний конгруентний метод
	// X = (aX + c) mod m
	var (
		i       int64 = 0
		x             = p["min"]
		length        = int64(p["length"])
		numbers       = make([]float64, length)
	)

	p["m"] = math.Pow(2, p["m"])
	for ; i < length; i++ {
		if x >= math.MaxFloat64 {
			x = p["min"]
		}
		x = (p["a"]*x + p["c"]) / p["m"] / p["max"] * float64(time.Now().UnixMicro())
		numbers[i] = x
	}
	return numbers
}
