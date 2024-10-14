# <h1 align="center">Laporan Praktikum 3 "Riview Struktur Kontrol" </h1>
<p align="center">Amelia Azmi - 2311102135</p>


# LATIHAN

## Latihan1

```go
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
```

### Output:

![Alt text](SoalLatihan1.png)

## Latihan2

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var pita string
	var bungaCount int

	for {
		fmt.Printf("Bunga %d: ", bungaCount+1)
		scanner.Scan()
		input := scanner.Text()

		if strings.ToLower(input) == "selesai" {
			break
		}

		if pita == "" {
			pita = input

		} else {
			pita += "[-]" + input
		}
		bungaCount++
	}
	fmt.Printf("Pita: %s\n", pita)
	fmt.Printf("Bunga: %s\n", bungaCount)
}

```

### Output:

![Alt text](SoalLatihan2.png)

# PERULANGAN

## Tugas1

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var totalBerat float64
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		scanner.Scan()
		input := scanner.Text()

		parts := strings.Split(input, " ")
		if len(parts) != 2 {
			fmt.Println("Input harus berupa dua buah bilangan real positif")
			continue
		}

		berat1, err1 := strconv.ParseFloat(parts[0], 64)
		berat2, err2 := strconv.ParseFloat(parts[1], 64)

		if err1 != nil || err2 != nil {
			fmt.Println("Input harus berupa bilangan real")
			continue
		}

		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalBerat += berat1 + berat2
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := abs(berat1 - berat2)
		fmt.Printf("Sepeda motor pak Andi akan oleng: %v\n", selisih >= 9)
	}
}

func abs(x float64) float64 {
	if x < 0 {
		return -x
	}
	return x
}

```

### Output:



## 

```go

```

### Output:


## 

```go

```

### Output:



## 

```go

```

### Output:



## 

```go

```

### Output:










