package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}

func dalam(cx, cy, r, x, y float64) bool {
	return jarak(cx, cy, x, y) <= r
}

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	fmt.Print("Pusat dan radius lingkaran 1 (cx cy r): ")
	fmt.Scan(&cx1, &cy1, &r1)

	fmt.Print("Pusat dan radius lingkaran 2 (cx cy r): ")
	fmt.Scan(&cx2, &cy2, &r2)

	fmt.Print("Masukkan koordinat titik (x y): ")
	fmt.Scan(&x, &y)

	dilingkaran1 := dalam(cx1, cy1, r1, x, y)
	dilingkaran2 := dalam(cx2, cy2, r2, x, y)

	if dilingkaran1 && dilingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dilingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dilingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
