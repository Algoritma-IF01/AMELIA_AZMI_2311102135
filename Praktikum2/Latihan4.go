package main

import (
	"fmt"
)

func main() {
	var celcius, fahrenheit, kelvin, reamur float64

	fmt.Print("Masukkan suhu dalam celcius: ")
	fmt.Scanln(&celcius)
	reamur = celcius * 4 / 5
	fahrenheit = (celcius * 9 / 5) + 32
	kelvin = celcius + 273.15

	fmt.Println("Suhu dalam reamur: ", reamur)
	fmt.Println("Suhu dalam fahrenheit: ", fahrenheit)
	fmt.Println("Suhu dalam kelvin: ", kelvin)

}
