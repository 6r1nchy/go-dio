package main

import "fmt"

func main() {
	// Definindo valores iniciais
    a, b := 10.0, 5.0

    fmt.Println("Adição:", Adicao(a, b))
    fmt.Println("Subtração:", Subtracao(a, b))
    fmt.Println("Multiplicação:", Multiplicacao(a, b))
    fmt.Println("Divisão:", Divisao(a, b))
}

// Funções para operações matemáticas básicas
func Adicao(a, b float64) float64 {
    return a + b
}

func Subtracao(a, b float64) float64 {
    return a - b
}

func Multiplicacao(a, b float64) float64 {
    return a * b
}

func Divisao(a, b float64) float64 {
    if b == 0 {
        fmt.Println("Erro: Divisão por zero não permitida.")
        return 0
    }
    return a / b
}
