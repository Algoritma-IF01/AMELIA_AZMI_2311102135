package main

import (
	"fmt"
)

func main() {
	var left, right float64
	var totalWeight float64

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scan(&left, &right)
		if err != nil {
			fmt.Println("Invalid input. Please enter two numbers.")
			continue
		}

		if left < 0 || right < 0 {
			fmt.Println("Berat tidak boleh negatif.")
			continue
		}

		totalWeight = left + right
		if totalWeight > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		diff := abs(left - right)
		if diff >= 9 {
			fmt.Printf("Sepeda motor pak Andi akan oleng: true\n")
		} else {
			fmt.Printf("Sepeda motor pak Andi akan oleng: false\n")
		}
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}
