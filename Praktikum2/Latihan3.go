package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("Jari-jari: ")
	var radius int
	fmt.Scanln(&radius)

	volume := (4.0 / 3.0) * math.Pi * math.Pow(float64(radius), 3)
	surfaceArea := 4 * math.Pi * math.Pow(float64(radius), 2)

	fmt.Printf("Bola dengan jari-jari %d memiliki volume %.4f dan luas kulit %.4f\n", radius, volume, surfaceArea)
}
