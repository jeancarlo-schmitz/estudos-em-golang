package lacoFor

import (
	"fmt"
)

func ExecutarForEach() {
	fmt.Println("\n\nExecutando o Laço de Repetição FOREACH")

	alunos := map[string]map[string]int{
		"João":   {"matematica": 85, "ingles": 58},
		"Maria":  {"matematica": 92},
		"Carlos": {"matematica": 78},
	}

	for nome, dadosAluno := range alunos {
		for materia, nota := range dadosAluno {
			fmt.Printf("Nome: %s, Materia: %s, Nota: %d\n", nome, materia, nota)
		}
	}
}
