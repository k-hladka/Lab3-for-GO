package random

func Probability(number int64, randoms []int64) float64 {
	var count int64
	for _, random := range randoms {
		if random == number {
			count++
		}
	}
	return float64(count) / float64(len(randoms))
}
