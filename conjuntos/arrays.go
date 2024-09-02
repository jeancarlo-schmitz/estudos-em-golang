package main

import (
	"fmt"
)

func main() {
	var amigos [5]string

	fmt.Println("Informe o nome de 5 dos seus amigos")
	for i := 0; i < len(amigos); i++ {
		fmt.Print("Informe o nome do ", (i + 1), " seu amigo: ")
		fmt.Scan(&amigos[i])
	}

	fmt.Println("Seus Amigos São")
	for _, amigo := range amigos {
		fmt.Printf(" - %s \n", amigo)
	}

	arrayInicializado := [3]int{3, 5, 6}
	fmt.Println(arrayInicializado)

	var matriz [3][3]int

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print("Digite um número: ")
			fmt.Scan(&matriz[i][j])
		}
	}

	fmt.Println(matriz)
}
