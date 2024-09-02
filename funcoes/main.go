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
	switch operacao {
	case "+":
		resultado := somar(numero1, numero2)
		fmt.Printf("%d + %d = %d\n", numero1, numero2, resultado)
	case "-":
		resultado := subtrair(numero1, numero2)
		fmt.Printf("%d - %d = %d\n", numero1, numero2, resultado)
	case "*":
		resultado := multiplicar(numero1, numero2)
		fmt.Printf("%d * %d = %d\n", numero1, numero2, resultado)
	case "/":
		if numero2 != 0 {
			resultado, resto := dividir(numero1, numero2)
			fmt.Printf("Quociente de: %d / %d = %d\n", numero1, numero2, resultado)
			fmt.Printf("Resto de: %d / %d = %d\n", numero1, numero2, resto)
		} else {
			fmt.Println("Erro: Divisão por zero não permitida!")
		}
	case "$":
		incremento, decremento := incrementoDecremento(numero1, numero2)
		fmt.Printf("Incremento = %d\nDecremento = %d\n", incremento, decremento)
	default:
		fmt.Println("Operação inválida!")
	}
}

func somar(numero1 int, numero2 int) int {
	return numero1 + numero2
}

func subtrair(numero1 int, numero2 int) int {
	return numero1 - numero2
}

func multiplicar(numero1 int, numero2 int) int {
	return numero1 * numero2
}

func dividir(numero1 int, numero2 int) (int, int) {
	quociente := numero1 / numero2
	resto := numero1 % numero2
	return quociente, resto
}

func incrementoDecremento(numero1 int, numero2 int) (int, int) {
	inc := numero1 + numero2
	dec := numero1 - numero2
	return inc, dec
}
