package lacowhile

import "fmt"

func ExecutaLacoWhile(numero int, qtdVezesMultiplicar int) {
	fmt.Println("\n\n Executando Laço While")
	i := 0
	for i <= 10 {
		fmt.Printf("%d X %d = %d \n", numero, i, numero*i)
		i++
	}
}
