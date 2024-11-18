package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menghitung total berat setiap wadah dan rata-rata berat per wadah
func calculateContainerWeightsAndAverage(x, y int, weights []float64) ([]float64, float64) {
	numContainers := int(math.Ceil(float64(x) / float64(y)))
	containerWeights := make([]float64, numContainers)
	total := 0.0

	// Hitung total berat per wadah dan rata-rata sekaligus
	for i, weight := range weights {
		containerWeights[i/y] += weight
		total += weight
	}

	average := total / float64(numContainers)
	return containerWeights, average
}

// Fungsi untuk validasi input
func validateInput(x, y int, weights []float64) error {
	if x <= 0 || y <= 0 || x > 1000 {
		return fmt.Errorf("jumlah ikan harus antara 1 dan 1000, dan kapasitas wadah harus lebih dari 0")
	}
	if len(weights) != x {
		return fmt.Errorf("jumlah berat ikan yang dimasukkan harus sebanyak %d", x)
	}
	for _, weight := range weights {
		if weight < 0 {
			return fmt.Errorf("berat ikan tidak boleh negatif")
		}
	}
	return nil
}

func main() {
	var x, y int
	fmt.Print("Masukkan banyak ikan yang akan dijual: ")
	fmt.Scan(&x)
	fmt.Print("Masukkan banyak ikan yang akan dimasukkan ke dalam wadah: ")
	fmt.Scan(&y)

	weights := make([]float64, x)
	fmt.Printf("Masukkan berat %d ikan (dipisahkan dengan spasi): ", x)
	for i := 0; i < x; i++ {
		if _, err := fmt.Scan(&weights[i]); err != nil {
			fmt.Println("Input berat ikan tidak valid.")
			return
		}
	}

	// Validasi input
	if err := validateInput(x, y, weights); err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Memanggil fungsi untuk menghitung total berat dan rata-rata
	containerWeights, totalAverage := calculateContainerWeightsAndAverage(x, y, weights)

	// Cetak hasil
	fmt.Println("\nTotal berat di setiap wadah:")
	for i, weight := range containerWeights {
		fmt.Printf("Wadah %d: %.2f\n", i+1, weight)
	}

	fmt.Printf("\nRata-rata berat per wadah: %.2f\n", totalAverage)
}
