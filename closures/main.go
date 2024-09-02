package main

import (
	"fmt"
)

func main() {
	var numero1, numero2 int
	var operacao string

	fmt.Print("Digite o primeiro número: ")
	fmt.Scanln(&numero1) // Usa Scanln para garantir que o input seja capturado corretamente

	fmt.Print("Digite a operação (+, -, *, /, $): ")
	fmt.Scanln(&operacao) // Usa Scanln para garantir que o input seja capturado corretamente

	fmt.Print("Digite o segundo número: ")
	fmt.Scanln(&numero2) // Usa Scanln para garantir que o input seja capturado corretamente

	handleOperacao(numero1, numero2, operacao)
}

func handleOperacao(numero1 int, numero2 int, operacao string) {
	var metodoOperacao func(n1 int, n2 int) (int, int)
	switch operacao {
	case "+":
		metodoOperacao = func(n1 int, n2 int) (int, int) {
			return n1 + n2, 0
		}
	case "-":
		metodoOperacao = func(n1 int, n2 int) (int, int) {
			return n1 - n2, 0
		}
	case "*":
		metodoOperacao = func(n1 int, n2 int) (int, int) {
			return n1 * n2, 0
		}
	case "/":
		if numero2 != 0 {
			metodoOperacao = func(n1 int, n2 int) (int, int) {
				return n1 / n2, n1 % n2
			}
		} else {
			fmt.Println("Erro: Divisão por zero não permitida!")
		}
	case "$":
		metodoOperacao = func(n1 int, n2 int) (int, int) {
			return n1 + n2, n1 - n2
		}
	default:
		fmt.Println("Operação inválida!")
	}

	resultado1, resultado2 := metodoOperacao(numero1, numero2)
	fmt.Printf("%d %s %d = %d, %d\n", numero1, operacao, numero2, resultado1, resultado2)
}
