package main

import (
	"fmt"
)

func main() {
	var numero int

	fmt.Println("Bem Vindo!")
	defer destruct()
	fmt.Print("Digite Um numero maior do que 10: ")
	fmt.Scanln(&numero) // Usa Scanln para garantir que o input seja capturado corretamente

	if numero <= 10 {
		panic("Número Inválido!")
	} else {
		fmt.Printf("Você Digitou um bom número! : %d\n", numero)
	}
}

func destruct() {
	fmt.Println("Obrigado por utilizar nosso software")
}
