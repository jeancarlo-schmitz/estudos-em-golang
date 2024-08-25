package main

import (
	"fmt"
	"lacos-repeticao/lacoFor"
)

func main() {
	fmt.Println("Inicializando Laços de Repetição")

	var numero, qtdVezesMultiplicar int
	fmt.Println("Digite um número para multiplicar: ")
	fmt.Scanln(&numero)

	fmt.Println("Por quantas vezes quer multiplicar esse numero? ")
	fmt.Scanln(&qtdVezesMultiplicar)

	lacoFor.ExecutaFor(numero, qtdVezesMultiplicar)
	lacoFor.ExecutarForEach()
}
