package main

import (
	"fmt"
	"math"
)

func main() {
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	fk := calculateFK(float64(k))
	fmt.Printf("Nilai f(k) = %.10f\n", fk)

	sqrt2 := calculateSqrt2(k)
	fmt.Printf("Nilai akar 2 = %.10f\n", sqrt2)
}

func calculateFK(k float64) float64 {
	return math.Pow((4*k+2), 2) / ((4*k + 1) * (4*k + 3))
}

func calculateSqrt2(k int) float64 {
	var sqrt2 float64
	for i := 0; i <= k; i++ {
		sqrt2 *= math.Pow((4*float64(i)+2), 2) / ((4*float64(i) + 1) * (4*float64(i) + 3))
	}
	return sqrt2
}
