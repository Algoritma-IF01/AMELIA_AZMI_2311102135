package main

import (
	"fmt"
	"math"
)

// Tipe bentukan untuk titik
type Point struct {
	x int
	y int
}

// Tipe bentukan untuk lingkaran
type Circle struct {
	center Point
	radius int
}

// Fungsi untuk menentukan apakah titik berada di dalam lingkaran
func isPointInCircle(point Point, circle Circle) bool {
	distance := math.Sqrt(float64((point.x-circle.center.x)*(point.x-circle.center.x) + (point.y-circle.center.y)*(point.y-circle.center.y)))
	return distance < float64(circle.radius)
}

func main() {
	var circle1, circle2 Circle
	var point Point

	// Input lingkaran 1
	fmt.Println("Masukkan koordinat pusat lingkaran 1 (cx cy) dan radius (r):")
	fmt.Scan(&circle1.center.x, &circle1.center.y, &circle1.radius)

	// Input lingkaran 2
	fmt.Println("Masukkan koordinat pusat lingkaran 2 (cx cy) dan radius (r):")
	fmt.Scan(&circle2.center.x, &circle2.center.y, &circle2.radius)

	// Input titik sembarang
	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&point.x, &point.y)

	// Cek posisi titik terhadap kedua lingkaran
	inCircle1 := isPointInCircle(point, circle1)
	inCircle2 := isPointInCircle(point, circle2)

	// Output hasil
	if inCircle1 && inCircle2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inCircle1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inCircle2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
