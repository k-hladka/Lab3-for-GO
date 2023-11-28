package random

func Interval(randoms []int64) map[int8]int64 {
	var intervalValues = map[int8]int64{0: 0, 1: 0}
	for _, number := range randoms {
		if number == 0 {
			intervalValues[0]++
		}
		if number == 1 {
			intervalValues[1]++
		}
	}

	return intervalValues
}
