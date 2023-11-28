package random

func Expectation(min, max, length int64) float64 {
	var result float64
	for i := min; i <= max; i++ {
		result += float64(i)
	}

	return result / float64(length)
}
