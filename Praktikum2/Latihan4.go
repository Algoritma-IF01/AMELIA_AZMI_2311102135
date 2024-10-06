package main

import (
	"fmt"
)

func main() {
	var celcius, fahrenheit, kelvin, reamur float64

	fmt.Print("Temperatur celcius: ")
	fmt.Scanln(&celcius)
	reamur = celcius * 4 / 5
	fahrenheit = (celcius * 9 / 5) + 32
	kelvin = celcius + 273.15

	fmt.Println("Derajat reamur: ", reamur)
	fmt.Println("Derajat fahrenheit: ", fahrenheit)
	fmt.Println("Derajat kelvin: ", kelvin)

}
