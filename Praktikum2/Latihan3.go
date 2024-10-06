package main

import (
	"fmt"
	"math"
)

const pi = 3.1415926535

func main() {
	var radius int
	fmt.Print("Jarijari=")
	fmt.Scan(&radius)

	volume = 4.0 / 3.0 * pi * math.Pow(float64(radius), 3)
	surfaceArea = 4 * pi * math.Pow(float64(radius), 2)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", radius, volume, surfaceArea)
}
