package random

import (
	"math"
)

func Dispersion(min, max, length int64, expectationResult float64) float64 {
	var result float64
	for i := min; i <= max; i++ {
		result += math.Pow(math.Abs(float64(i)-expectationResult), 2)
	}

	return result / float64(length)
}
