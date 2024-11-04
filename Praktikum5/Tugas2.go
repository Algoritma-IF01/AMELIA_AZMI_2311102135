package main

import (
	"fmt"
	"math"
)

// Fungsi untuk menampilkan keseluruhan isi array
func displayArray(arr []int) {
	fmt.Println("Isi array:", arr)
}

// Fungsi untuk menampilkan elemen dengan indeks ganjil
func displayOddIndices(arr []int) {
	fmt.Print("Elemen dengan indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks genap
func displayEvenIndices(arr []int) {
	fmt.Print("Elemen dengan indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}
	fmt.Println()
}

// Fungsi untuk menampilkan elemen dengan indeks kelipatan x
func displayMultiples(arr []int, x int) {
	fmt.Print("Elemen dengan indeks kelipatan ", x, ": ")
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Print(arr[i], " ")
		}
	}
	fmt.Println()
}

// Fungsi untuk menghapus elemen pada indeks tertentu
func deleteElement(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

// Fungsi untuk menghitung rata-rata
func calculateAverage(arr []int) float64 {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return float64(sum) / float64(len(arr))
}

// Fungsi untuk menghitung standar deviasi
func calculateStdDev(arr []int, avg float64) float64 {
	sum := 0.0
	for _, value := range arr {
		sum += math.Pow(float64(value)-avg, 2)
	}
	return math.Sqrt(sum / float64(len(arr)))
}

// Fungsi untuk menghitung frekuensi dari bilangan tertentu
func calculateFrequency(arr []int, target int) int {
	count := 0
	for _, value := range arr {
		if value == target {
			count++
		}
	}
	return count
}

func main() {
	var N int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&N)

	if N <= 0 {
		fmt.Println("Jumlah elemen harus lebih dari 0.")
		return
	}

	arr := make([]int, N)

	// Mengisi array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < N; i++ {
		fmt.Print("Elemen ke-", i, ": ")
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	displayArray(arr)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil
	displayOddIndices(arr)

	// c. Menampilkan elemen-elemen array dengan indeks genap
	displayEvenIndices(arr)

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x
	var x int
	fmt.Print("Masukkan bilangan x untuk kelipatan: ")
	fmt.Scan(&x)
	displayMultiples(arr, x)

	// e. Menghapus elemen array pada indeks tertentu
	var deleteIndex int
	fmt.Print("Masukkan indeks elemen yang ingin dihapus: ")
	fmt.Scan(&deleteIndex)
	if deleteIndex >= 0 && deleteIndex < len(arr) {
		arr = deleteElement(arr, deleteIndex)
		displayArray(arr)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array
	average := calculateAverage(arr)
	fmt.Printf("Rata-rata: %.2f\n", average)

	// g. Menampilkan standar deviasi dari bilangan yang ada di dalam array
	stdDev := calculateStdDev(arr, average)
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array
	var target int
	fmt.Print("Masukkan bilangan untuk menghitung frekuensi: ")
	fmt.Scan(&target)
	frequency := calculateFrequency(arr, target)
	fmt.Printf("Frekuensi dari %d: %d\n", target, frequency)
}
