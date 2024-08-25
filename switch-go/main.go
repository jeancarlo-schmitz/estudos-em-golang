package main

import (
	"fmt"
)

func main() {
	var mes int
	fmt.Println("Digite o número do mês")
	fmt.Scanf("%d", &mes)

	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	case 4:
		fmt.Println("Abril")
	case 5:
		fmt.Println("Maio")
	case 6:
		fmt.Println("Junho")
	case 7:
		fmt.Println("Julho")
	case 8:
		fmt.Println("Agosto")
	case 9:
		fmt.Println("Setembro")
	case 10:
		fmt.Println("Outubro")
	case 11:
		fmt.Println("Novembro")
	case 12:
		fmt.Println("Dezembro")
	default:
		fmt.Println("Mês Inválido!")
	}

	switch mes {
	case 1, 2, 3:
		fmt.Println("Primeiro Trimestre")
	case 4, 5, 6:
		fmt.Println("Segundo Trimestre")
	case 7, 8, 9:
		fmt.Println("Terceiro Trimestre")
	case 10, 11, 12:
		fmt.Println("Quarto Trimestre")
	}

	switch true {
	case mes >= 1 && mes <= 6:
		fmt.Println("Primeiro Semestre")
	case mes >= 7 && mes <= 12:
		fmt.Println("Segundo Semetre")
	}
}
