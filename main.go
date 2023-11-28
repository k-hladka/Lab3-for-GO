package main

import (
	"fmt"
	"random/random"
)

func main() {
	params1 := map[string]int64{"a": 1103515245, "c": 12345, "m": 31, "min": 0, "max": 100, "length": 10000}
	randoms := random.Random(params1)
	interval := random.Interval(randoms)
	probability := random.Probability(17, randoms)
	expectation := random.Expectation(params1["min"], params1["max"], params1["length"])
	dispersion := random.Dispersion(params1["min"], params1["max"], params1["length"], expectation)
	deviation := random.Deviation(dispersion)

	fmt.Println("Псевдовипадкові значення = ", randoms)
	fmt.Println("Частота інтервалів	0 = ", interval[0], " 1 = ", interval[1])
	fmt.Println("Статистична імовірність появи числа 17 = ", probability)
	fmt.Printf("Математичне сподівання = %.3f\n", expectation)
	fmt.Printf("Дисперсія = %.3f\n", dispersion)
	fmt.Printf("Середньоквадратичне відхилення = %.3f\n", deviation)

	params2 := map[string]float64{"a": 1664525, "c": 1013904223, "m": 32, "min": 0, "max": 150, "length": 20000}
	randoms2 := random.RandomFloat(params2)
	fmt.Println("************\nПсевдовипадкові дійсні значення = ", randoms2)
}
