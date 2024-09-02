package main

import (
	"fmt"
)

func main() {
	fmt.Println("Estou abrindo uma conexão com o banco de dados")
	defer fecharConexaoBancoDeDados()
	executarConsulta()
}

func executarConsulta() {
	fmt.Println("Estou Executando a consulta com o banco de dados")
}

func fecharConexaoBancoDeDados() {
	fmt.Println("Fechando Conexão com Banco de Dados")
}
