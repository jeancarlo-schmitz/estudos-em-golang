package main

import (
	"fmt"
)

func main() {
	// var saldo float64
	// var saldo *float64 = new(float64)
	var saldo = new(float64)
	fmt.Println("Bem Vindo!")
	fmt.Print("Digite o seu saldo: ")
	// fmt.Scanln(&saldo)
	fmt.Scanln(saldo)
	// calcularRendimentos(&saldo)
	calcularRendimentos(saldo)
	fmt.Printf("Seu Saldo com Rendimento é de R$ %.2f\n", *saldo)
}

func calcularRendimentos(saldo *float64) {
	*saldo += *saldo * 0.03
}
