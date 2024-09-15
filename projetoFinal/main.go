package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const ARQUVIVO string = "agenda.txt"

type ConversivelParaString interface {
	toString() string
}

type Contato struct {
	Nome         string
	FormaContato string
	ValorContato string
}

func (contato *Contato) toString() string {
	return fmt.Sprintf("%s|%s|%s \n", contato.Nome, contato.FormaContato, contato.ValorContato)
}

type GerenciadorContatos struct{}

func (gerenciador *GerenciadorContatos) carregarContatos() ([]Contato, error) {
	contatos := make([]Contato, 0)

	if _, e := os.Stat(ARQUVIVO); !os.IsNotExist(e) {
		arquivo, err := os.Open(ARQUVIVO)
		if err != nil {
			return contatos, err
		}
		defer arquivo.Close()
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linhaContato := scanner.Text()
			infoContato := strings.Split(linhaContato, "|")
			contatos = append(contatos, Contato{Nome: infoContato[0], FormaContato: infoContato[1], ValorContato: infoContato[2]})
		}
	}

	return contatos, nil
}

func (gerenciador *GerenciadorContatos) salvarContato(contato ConversivelParaString) error {
	arquivo, err := os.OpenFile(ARQUVIVO, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		return err
	}
	defer arquivo.Close()

	if _, e := arquivo.WriteString(contato.toString()); e != nil {
		return e
	}

	return nil
}

func main() {
	gerenciador := GerenciadorContatos{}
	opcao := 0
	for opcao != 3 {
		fmt.Println("O que você deseja fazer?")
		fmt.Println(" - 1) Listar Contatos")
		fmt.Println(" - 2) Criar Novo Contato")
		fmt.Println(" - 3) Sair")
		fmt.Scanln(&opcao)

		switch opcao {
		case 1:
			listarContatos(&gerenciador)
		case 2:
			criarNovoContato(&gerenciador)
		}
	}

	fmt.Println(" Bye Bye")
}

func listarContatos(gerenciador *GerenciadorContatos) {
	contatos, err := gerenciador.carregarContatos()
	if err != nil {
		fmt.Printf("Não foi Possivel carregar os contatos")
	} else {
		fmt.Println("Lista de Contatos")
		for _, contato := range contatos {
			fmt.Printf(" - %s (%s): %s\n", contato.Nome, contato.FormaContato, contato.ValorContato)
		}
	}
}

func criarNovoContato(gerenciador *GerenciadorContatos) {
	novoContato := Contato{}

	fmt.Println("Nome do Contato: ")
	fmt.Scanln(&novoContato.Nome)
	fmt.Println("Forma do Contato: ")
	fmt.Scanln(&novoContato.FormaContato)
	fmt.Println("Contato: ")
	fmt.Scanln(&novoContato.ValorContato)

	err := gerenciador.salvarContato(&novoContato)
	if err != nil {
		fmt.Printf("Houve um erro: %s\n", err)
	}
}
