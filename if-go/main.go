package main

import (
	"fmt"
)

func main() {
	var idade int
	fmt.Println("Digite sua Idade")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Você é Maior de Idade")
	} else {
		fmt.Println("Você é Menor de Idade")
	}
}
