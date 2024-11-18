# <h1 align="center">Laporan Praktikum 6 "PENCARIAN NILAI EKSTRIM PADA HIMPUNAN DATA" </h1>
<p align="center">Amelia Azmi - 2311102135</p>


# LATIHAN

## Latihan1

```go
package main

import (
	"fmt"
)

type arrInt [2023]int

// Fungsi untuk mencari indeks dari nilai terkecil
func terkecil_2(tabInt arrInt, n int) int {
	var idx int = 0 // indeks data pertama
	var j int = 1   // pencarian dimulai dari data kedua
	for j < n {
		if tabInt[idx] > tabInt[j] { // cek apakah tabInt[j] lebih kecil dari tabInt[idx]
			idx = j // update idx ke indeks baru yang nilainya lebih kecil
		}
		j = j + 1
	}
	return idx // mengembalikan indeks dari nilai terkecil
}

func main() {
	var n int
	var data arrInt

	// Input jumlah elemen N
	fmt.Print("Masukkan jumlah elemen (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi N agar tidak melebihi kapasitas array
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah elemen harus antara 1 dan 2023")
		return
	}

	// Input elemen-elemen array
	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&data[i])
	}

	// Panggil fungsi untuk mencari indeks nilai terkecil
	idxTerkecil := terkecil_2(data, n)
	fmt.Printf("Indeks nilai terkecil: %d\n", idxTerkecil)
	fmt.Printf("Nilai terkecil: %d\n", data[idxTerkecil])
}
```

### Output:

![Alt text](Latihan1.png)

## Latihan2

```go
package main

import (
	"fmt"
)

// Mendefinisikan tipe data mahasiswa
type mahasiswa struct {
	nama, nim, kelas, jurusan string
	ipk                       float64
}

// Mendefinisikan array mahasiswa dengan kapasitas 2023
type arrMhs [2023]mahasiswa

// Fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
func IPK_2(T arrMhs, n int) int {
	// idx menyimpan indeks mahasiswa dengan IPK tertinggi sementara
	var idx int = 0
	var j int = 1
	for j < n {
		if T[idx].ipk < T[j].ipk {
			idx = j
		}
		j = j + 1
	}
	return idx
}

func main() {
	var n int
	var data arrMhs

	// Input jumlah mahasiswa
	fmt.Print("Masukkan jumlah mahasiswa (N <= 2023): ")
	fmt.Scan(&n)

	// Validasi jumlah mahasiswa
	if n <= 0 || n > 2023 {
		fmt.Println("Jumlah mahasiswa harus antara 1 dan 2023")
		return
	}

	// Input data mahasiswa
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data mahasiswa ke-%d\n", i+1)
		fmt.Print("Nama: ")
		fmt.Scan(&data[i].nama)
		fmt.Print("NIM: ")
		fmt.Scan(&data[i].nim)
		fmt.Print("Kelas: ")
		fmt.Scan(&data[i].kelas)
		fmt.Print("Jurusan: ")
		fmt.Scan(&data[i].jurusan)
		fmt.Print("IPK: ")
		fmt.Scan(&data[i].ipk)
	}

	// Panggil fungsi untuk mencari indeks mahasiswa dengan IPK tertinggi
	idxTertinggi := IPK_2(data, n)

	// Tampilkan data mahasiswa dengan IPK tertinggi
	fmt.Println("\nMahasiswa dengan IPK tertinggi:")
	fmt.Printf("Nama    : %s\n", data[idxTertinggi].nama)
	fmt.Printf("NIM     : %s\n", data[idxTertinggi].nim)
	fmt.Printf("Kelas   : %s\n", data[idxTertinggi].kelas)
	fmt.Printf("Jurusan : %s\n", data[idxTertinggi].jurusan)
	fmt.Printf("IPK     : %.2f\n", data[idxTertinggi].ipk)
}
```

### Output:

![Alt text](Latihan2.png)

# TUGAS

## Tugas1

```go
package main

import "fmt"

func main() {
	var N int
	fmt.Println("Masukkan jumlah anak kelinci :")
	fmt.Scan(&N)

	if N <= 0 || N > 1000{
		fmt.Println("Jumlah anak kelinci harus antara 1 dan 1000")
		return
	}
	weight := make([]float64, N)
	fmt.Println("Masukkan berat anak kelinci : ")
	for i := 0; i < N; i++ {
		fmt.Scan(&weight[i])
	}

	minWeight, maxWeight := weight[0], weight[0]

	for _, weight := range weight [1:] {
		if weight < minWeight {
			minWeight = weight
		}
		if weight > maxWeight {
			maxWeight = weight
		}
	}

	fmt.Printf("Berat kelinci terkecil : %.2f kg\n", minWeight)
	fmt.Printf("Berat kelinci terbesar : %.2f kg\n", maxWeight)
}
```

### Output:

![Alt text](Tugas1.png)

## Tugas2

```go
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
```

### Output:

![Alt text](Tugas2.png)

## Tugas3

```go
package main

import (
	"fmt"
)

// Definisi tipe array dengan kapasitas maksimum 100 elemen
type arrBalita [100]float64

// Fungsi untuk menghitung nilai minimum dan maksimum dalam array
func hitungMinMax(arrBerat arrBalita, bMin, bMax *float64, n int) {
	// Inisialisasi nilai awal bMin dan bMax dengan elemen pertama array
	*bMin = arrBerat[0]
	*bMax = arrBerat[0]

	// Iterasi untuk mencari nilai minimum dan maksimum dari elemen array
	for i := 0; i < n; i++ {
		if arrBerat[i] < *bMin {
			*bMin = arrBerat[i] // Jika elemen lebih kecil dari bMin, perbarui bMin
		}
		if arrBerat[i] > *bMax {
			*bMax = arrBerat[i] // Jika elemen lebih besar dari bMax, perbarui bMax
		}
	}
}

// Fungsi untuk menghitung rata-rata berat balita dalam array
func rerata(arrBerat arrBalita, n int) float64 {
	total := 0.0 // Variabel untuk menyimpan jumlah total berat balita

	// Iterasi untuk menjumlahkan semua elemen array
	for i := 0; i < n; i++ {
		total += arrBerat[i] // Tambahkan berat balita ke dalam total
	}

	// Mengembalikan rata-rata (total berat dibagi jumlah elemen)
	return total / float64(n)
}

func main() {
	var n int              // Variabel untuk jumlah data berat balita
	var arrBerat arrBalita // Array untuk menyimpan berat balita
	var bMin, bMax float64 // Variabel untuk menyimpan nilai minimum dan maksimum

	// Meminta input jumlah data berat balita
	fmt.Print("Masukkan banyak data berat balita: ")
	fmt.Scan(&n)

	// Validasi jumlah data (harus antara 1 dan 100)
	if n <= 0 || n > 100 {
		fmt.Println("Jumlah data harus antara 1 hingga 100.")
		return // Jika input tidak valid, hentikan program
	}

	// Input data berat balita
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan berat balita ke-%d: ", i+1)
		fmt.Scan(&arrBerat[i]) // Menyimpan input ke dalam array
	}

	// Memanggil fungsi untuk menghitung nilai minimum dan maksimum
	hitungMinMax(arrBerat, &bMin, &bMax, n)

	// Memanggil fungsi untuk menghitung rata-rata berat balita
	rataRata := rerata(arrBerat, n)

	// Menampilkan hasil
	fmt.Printf("\nBerat balita minimum: %.2f kg\n", bMin)  // Menampilkan berat minimum
	fmt.Printf("Berat balita maksimum: %.2f kg\n", bMax)   // Menampilkan berat maksimum
	fmt.Printf("Rerata berat balita: %.2f kg\n", rataRata) // Menampilkan rata-rata berat
}
```

### Output:

![Alt text](Tugas3.png)