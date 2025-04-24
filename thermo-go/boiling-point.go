// Desafio 1: Converta o ponto de ebulição da água em Kelvin (K) para Celsius (°C).

package main

import "fmt"

func main() {
    // Definir constante para o ponto de ebulição da água em Kelvin
    const boilingPointKelvin = 373.15

    // Converter para graus Celsius
    boilingPointCelsius := boilingPointKelvin - 273.15

    // Exibir o resultado
    fmt.Printf("O ponto de ebulição da água é %.2f°C.\n", boilingPointCelsius)
}
