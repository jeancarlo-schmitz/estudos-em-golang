package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	mapaCursos := make(map[string]int)

	scanner := bufio.NewScanner(os.Stdin)
	curso := ""
	for curso != "q" {
		var cargaHorario int
		fmt.Print("Digite um curso ou 'q' para sair: ")
		scanner.Scan()
		curso = scanner.Text()
		if curso != "q" {
			fmt.Print("Digite a carga horario do curso: \n")
			fmt.Scan(&cargaHorario)
			mapaCursos[curso] = cargaHorario
		}
	}

	fmt.Println("Cursos Registrados")

	for nomeCurso, cargaHorario := range mapaCursos {
		fmt.Printf("Curso %s com a Carga Horaria: %dh", &nomeCurso, cargaHorario)
	}

}
