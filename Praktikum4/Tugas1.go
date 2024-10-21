package main

import (
	"fmt"
)

// Fungsi yang ada
func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

// Fungsi baru
func k(x int) int {
	return x * 3
}

// Fungsi komposisi baru
func kofg(c int) int {
	return k(g(c))
}

// Fungsi komposisi yang ada
func fogoh(a int) int {
	return f(g(h(a)))
}

func gohof(b int) int {
	return g(h(f(b)))
}

func hofog(c int) int {
	return h(f(g(c)))
}

func main() {
	var a, b, c int

	fmt.Println("Masukkan nilai a, b, dan c:")
	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Println("fogoh(a):", fogoh(a))
	fmt.Println("gohof(b):", gohof(b))
	fmt.Println("hofog(c):", hofog(c))
	fmt.Println("kofg(c):", kofg(c)) // Menampilkan hasil dari fungsi baru
}
