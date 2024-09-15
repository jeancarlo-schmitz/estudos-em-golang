package main

import (
	"fmt"
	"time"
)

func main() {
	var limite int
	canal1 := make(chan int)
	canal2 := make(chan int)
	canal3 := make(chan int)
	canal4 := make(chan int)
	fmt.Println("Informe um Limite: ")
	fmt.Scanln(&limite)
	go conteAteVezes20(limite, canal1)
	go conteAteVezes10(limite, canal2)
	go conteAteVezes30(limite, canal3)
	go conteAteVezes5(limite, canal4)

	resultadoCanal1 := <-canal1
	resultadoCanal2 := <-canal2
	resultadoCanal3 := <-canal3
	resultadoCanal4 := <-canal4
	fmt.Printf(" - Resultado Canal 1: %d\nResultado Canal 2: %d\nResultado Canal 3: %d\nResultado Canal 4: %d\n", resultadoCanal1, resultadoCanal2, resultadoCanal3, resultadoCanal4)

}

func conteAteVezes20(limite int, canal chan int) {
	resultado := 0
	for i := 0; i <= limite*20; i++ {
		fmt.Printf(" - [conteAteVezes20] O número é de: %d\n", i)
		resultado = i * 20
		time.Sleep(2 * time.Second)
	}

	canal <- resultado

}

func conteAteVezes10(limite int, canal chan int) {
	resultado := 0
	for i := 0; i <= limite*10; i++ {
		fmt.Printf(" - [conteAteVezes10] O número é de: %d\n", i)
		resultado = i * 10
		time.Sleep(3 * time.Second)
	}

	canal <- resultado
}

func conteAteVezes30(limite int, canal chan int) {
	resultado := 0
	for i := 0; i <= limite*30; i++ {
		fmt.Printf(" - [conteAteVezes30] O número é de: %d\n", i)
		resultado = i * 30
		time.Sleep(1 * time.Second)
	}

	canal <- resultado
}

func conteAteVezes5(limite int, canal chan int) {
	resultado := 0
	for i := 0; i <= limite*5; i++ {
		fmt.Printf(" - [conteAteVezes5] O número é de: %d\n", i)
		resultado = i * 5
		time.Sleep(4 * time.Second)
	}

	canal <- resultado
}
