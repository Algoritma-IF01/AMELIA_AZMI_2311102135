package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel
	var nama string

	// Minta input dari pengguna
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)

	// Format output
	hasil := "Hai, " + nama

	// Cetak hasil
	fmt.Println(hasil)
}
