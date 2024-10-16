package main

import "fmt"

func main() {
	var number int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&number)

	fmt.Print("Faktor: ")
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	isPrime := true
	if number <= 1 {
		isPrime = false
	} else {
		for i := 2; i*i <= number; i++ {
			if number%i == 0 {
				isPrime = false
				break
			}
		}
	}

	fmt.Println("Prima:", isPrime)
}
