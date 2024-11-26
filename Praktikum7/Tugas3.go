package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Fungsi untuk mengurutkan array menggunakan Insertion Sort
func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		key := arr[i]
		j := i - 1
		// Geser elemen-elemen yang lebih besar dari key ke kanan
		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j--
		}
		// Tempatkan key pada posisi yang sesuai
		arr[j+1] = key
	}
}

// Fungsi untuk memeriksa apakah jarak antar elemen tetap
func penghitunganjarak(arr []int) string {
	// Jika panjang array kurang dari 2, tidak bisa dibandingkan jaraknya
	if len(arr) < 2 {
		return "Data berjarak tidak tetap"
	}

	// Menghitung selisih antara elemen pertama dan kedua
	diff := arr[1] - arr[0]

	// Memeriksa selisih antar elemen lainnya
	for i := 2; i < len(arr); i++ {
		if arr[i]-arr[i-1] != diff {
			return "Data berjarak tidak tetap"
		}
	}

	// Jika selisih antar elemen sama
	return fmt.Sprintf("Data berjarak %d", diff)
}

func main() {
	// Membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Masukkan sekumpulan bilangan bulat (akhiri dengan bilangan negatif):")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Memisahkan input yang dimasukkan menjadi elemen-elemen string
	inputs := strings.Fields(input)

	var numbers []int
	// Mengonversi string menjadi integer dan menambahkan elemen ke slice
	for _, val := range inputs {
		num, _ := strconv.Atoi(val)
		// Jika bilangan negatif ditemukan, berhenti membaca input
		if num < 0 {
			break
		}
		numbers = append(numbers, num)
	}

	// Mengurutkan angka dengan Insertion Sort
	insertionSort(numbers)

	// Mengecek jarak antar elemen setelah pengurutan
	status := penghitunganjarak(numbers)

	// Menampilkan hasil pengurutan
	fmt.Println("Hasil setelah pengurutan:")
	for _, num := range numbers {
		fmt.Printf("%d ", num)
	}
	fmt.Println()
	fmt.Println(status)
}
