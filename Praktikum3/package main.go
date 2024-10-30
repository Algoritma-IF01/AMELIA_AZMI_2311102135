package main

import (
	"fmt"
	"strconv"
)

func isValidVoucher(voucher string) bool {
	length := len(voucher)

	// Cek panjang nomor seri
	if length != 5 && length != 6 {
		return false
	}

	// Ambil digit pertama dan kedua
	firstDigit, _ := strconv.Atoi(string(voucher[0]))
	secondDigit, _ := strconv.Atoi(string(voucher[1]))

	// Ambil digit terakhir
	lastDigit, _ := strconv.Atoi(string(voucher[length-1]))
	secondLastDigit, _ := strconv.Atoi(string(voucher[length-2]))

	// Cek apakah perkalian digit dua digit pertama dan dua digit terakhir sama
	if (firstDigit * secondDigit) != (lastDigit * secondLastDigit) {
		return false
	}

	// Cek digit tengah
	var middleDigit int
	if length == 5 {
		middleDigit, _ = strconv.Atoi(string(voucher[2]))
	} else { // length == 6
		middleDigit1, _ := strconv.Atoi(string(voucher[2]))
		middleDigit2, _ := strconv.Atoi(string(voucher[3]))
		middleDigit3, _ := strconv.Atoi(string(voucher[4]))
		middleDigit = middleDigit1 + middleDigit2 + middleDigit3 // Menggunakan jumlah untuk validasi genap
	}

	// Cek apakah digit tengah genap
	if length == 5 {
		return middleDigit%2 == 0
	} else { // length == 6
		return middleDigit%2 == 0
	}
}

func main() {
	var N int
	fmt.Println("Masukkan jumlah mahasiswa:")
	fmt.Scan(&N)

	validCount := 0
	invalidCount := 0

	for i := 0; i < N; i++ {
		var voucher string
		fmt.Printf("Masukkan nomor seri voucher mahasiswa ke-%d: ", i+1)
		fmt.Scan(&voucher)

		if isValidVoucher(voucher) {
			validCount++
		} else {
			invalidCount++
		}
	}

	fmt.Printf("Jumlah voucher valid: %d\n", validCount)
	fmt.Printf("Jumlah voucher tidak valid: %d\n", invalidCount)
}