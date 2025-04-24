// Desafio extra: Conversor de temperatura entre Celsius, Fahrenheit e Kelvin

package main

import "fmt"

func main() {
	// Convert Celsius to Fahrenheit
	celsius := 25.0
	fahrenheit := (celsius * 9/5) + 32
	fmt.Printf("%.2f Celsius is %.2f Fahrenheit\n", celsius, fahrenheit)

	// Convert Fahrenheit to Celsius
	fahrenheit = 77.0
	celsius = (fahrenheit - 32) * 5/9
	fmt.Printf("%.2f Fahrenheit is %.2f Celsius\n", fahrenheit, celsius)
	// Convert Celsius to Kelvin
	celsius = 25.0
	kelvin := celsius + 273.15
	fmt.Printf("%.2f Celsius is %.2f Kelvin\n", celsius, kelvin)
}