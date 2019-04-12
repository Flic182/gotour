package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	const diffTolerance = 0.000001
	z := 1.0
	previous := 0.0

	// for i := 0; i < 10; i++ {
	for diff := 1.0; diff > diffTolerance; diff = math.Abs(z - previous) {
		previous = z
		z -= (z * z - x) / (2 * z)
		fmt.Println(z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(9))
	fmt.Println(math.Sqrt(9))
}
