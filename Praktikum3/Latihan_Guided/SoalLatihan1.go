package main

import (
	"fmt"
)

func main() {
	// Definisikan wrana yang benar
	urutanBenar := []string{"merah", "kuning", "hijau", "ungu"}
	hasil := true

	// lakukan 5 percobaan
	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("percobaan %d\n", i)
		fmt.Print("masukkan warna pertama: ")
		fmt.Scanln(&warna1)
		fmt.Print("masukkan warna kedua: ")
		fmt.Scanln(&warna2)
		fmt.Print("masukkan warna ketiga: ")
		fmt.Scanln(&warna3)
		fmt.Print("masukkan warna keempat: ")
		fmt.Scanln(&warna4)

		/// Periksa apakah urutan warna sesuai
		if warna1 != urutanBenar[0] || warna2 != urutanBenar[1] ||
			warna3 != urutanBenar[2] || warna4 != urutanBenar[3] {

		}

	}
	// Tampilkan hasil
	fmt.Println("BERHASIL:", hasil)

}
