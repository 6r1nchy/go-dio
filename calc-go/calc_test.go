package main

import (
    "testing"
    "math"
)

// Função para arredondar para duas casas decimais
func roundTwoDecimals(value float64) float64 {
    return math.Round(value*100) / 100
}

func TestAdicaoPassa(t *testing.T) {
    resultado := roundTwoDecimals(Adicao(3, 4))
    esperado := roundTwoDecimals(7.00)
    if resultado != esperado {
        t.Errorf("Erro na Adição: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}

func TestAdicaoFalha(t *testing.T) {
    resultado := roundTwoDecimals(Adicao(3, 4))
    esperado := roundTwoDecimals(8.00) // Valor errado propositalmente
    if resultado != esperado {
        t.Errorf("Erro na Adição: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}

func TestSubtracaoPassa(t *testing.T) {
    resultado := roundTwoDecimals(Subtracao(10, 4))
    esperado := roundTwoDecimals(6.00)
    if resultado != esperado {
        t.Errorf("Erro na Subtração: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}

func TestSubtracaoFalha(t *testing.T) {
    resultado := roundTwoDecimals(Subtracao(10, 4))
    esperado := roundTwoDecimals(5.00) // Valor errado propositalmente
    if resultado != esperado {
        t.Errorf("Erro na Subtração: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}

func TestMultiplicacaoPassa(t *testing.T) {
    resultado := roundTwoDecimals(Multiplicacao(3, 3))
    esperado := roundTwoDecimals(9.00)
    if resultado != esperado {
        t.Errorf("Erro na Multiplicação: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}

func TestMultiplicacaoFalha(t *testing.T) {
    resultado := roundTwoDecimals(Multiplicacao(3, 3))
    esperado := roundTwoDecimals(10.00) // Valor errado propositalmente
    if resultado != esperado {
        t.Errorf("Erro na Multiplicação: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}

func TestDivisaoPassa(t *testing.T) {
    resultado := roundTwoDecimals(Divisao(10, 2))
    esperado := roundTwoDecimals(5.00)
    if resultado != esperado {
        t.Errorf("Erro na Divisão: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}

func TestDivisaoFalha(t *testing.T) {
    resultado := roundTwoDecimals(Divisao(10, 2))
    esperado := roundTwoDecimals(6.00) // Valor errado propositalmente
    if resultado != esperado {
        t.Errorf("Erro na Divisão: esperado %.2f, obtido %.2f", esperado, resultado)
    }
}
