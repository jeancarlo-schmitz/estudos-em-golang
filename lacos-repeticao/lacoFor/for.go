package lacoFor

import (
	"fmt"
)

func ExecutaFor(numero int, qtdVezesMultiplicar int) {
	fmt.Println("Executando o Laço de Repetição FOR")
	fmt.Println("Multiplicando os valores por: ", numero)
	fmt.Println("Por: ", qtdVezesMultiplicar, " Vezes")

	for i := 0; i <= qtdVezesMultiplicar; i++ {
		fmt.Printf("%d X %d = %d \n", numero, i, numero*i)
	}
}
