package main

import (
	"fmt"
)

func main() {
	amigos := make([]string, 0)
	fmt.Println(len(amigos))
	nome := ""
	i := 0
	for nome != "q" {
		fmt.Println("Digite um nome, ou digite 'q' para sair: ")
		fmt.Scan(&nome)
		if nome != "q" {
			amigos = append(amigos, nome)
		}
		i++
	}

	fmt.Println(amigos)
}
