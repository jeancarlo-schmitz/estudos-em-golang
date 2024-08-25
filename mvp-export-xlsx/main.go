package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/xuri/excelize/v2"
)

type Publicacao struct {
	IDPublicacao         string `json:"idPublicacao"`
	NumeroProcesso       string `json:"numeroProcesso"`
	NomeAdvogado         string `json:"nomeAdvogado"`
	OrigemIntimacao      string `json:"origemIntimacao"`
	VariacaoEncontrada   string `json:"variacaoEncontrada"`
	DataCadastro         string `json:"dataCadastro"`
	DataDisponibilizacao string `json:"dataDisponibilizacao"`
	DataPublicacao       string `json:"dataPublicacao"`
	InicioPrazo          string `json:"inicioPrazo"`
	FinalPrazo           string `json:"finalPrazo"`
	Orgao                string `json:"orgao"`
	UF                   string `json:"uf"`
	Vara                 string `json:"vara"`
	Cidade               string `json:"cidade"`
	Jornal               string `json:"jornal"`
	ProcessoEletronico   string `json:"processoEletronico"`
	ExpedidaConfirmada   string `json:"expedidaConfirmada"`
	TipoIntimacao        string `json:"tipoIntimacao"`
	ListaProcEletronico  string `json:"listaProcEletronico"`
	PerfilAdvogado       string `json:"perfilAdvogado"`
	Conteudo             string `json:"conteudo"`
	Anexos               string `json:"anexos"`
}

func generateExcelReport(w http.ResponseWriter, r *http.Request) {
	var listaPublicacoes []Publicacao
	fmt.Println("hehe")
	err := json.NewDecoder(r.Body).Decode(&listaPublicacoes)
	if err != nil {
		fmt.Println("ERRO")
		fmt.Println(err)

		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}
	fmt.Println(listaPublicacoes)

	// Create a new Excel file
	f := excelize.NewFile()

	index, err := f.NewSheet("Relatório de intimações")
	if err != nil {
		fmt.Errorf("Failed to create sheet: %v", err)
	}
	// Set column names
	headers := []string{
		"Seq.", "Código", "Nº do Processo", "Termo de pesquisa", "Origem da Intimação", "Termo localizado na publicação",
		"Data de Cadastro", "Data de Disponibilização", "Data de Publicação", "Prazo Inicial", "Prazo Final", "Órgão",
		"UF", "Vara", "Cidade", "Jornal", "Sistema", "Expedida/Confirmada", "Tipo Intimação", "Lista Capturada do sistema proc. ele.",
		"Perfil de Proc. eletrônico", "Publicação", "Anexo - Links",
	}

	for i, header := range headers {
		cell := fmt.Sprintf("%s1", string('A'+i))
		if err := f.SetCellValue("Relatório de intimações", cell, header); err != nil {
			http.Error(w, "Error setting header", http.StatusInternalServerError)
			return
		}
	}

	// Populate the Excel sheet with data
	for rowIndex, pub := range listaPublicacoes {
		row := rowIndex + 2 // Start from row 2 (row 1 is for headers)
		data := []interface{}{
			rowIndex + 1,
			pub.IDPublicacao,
			pub.NumeroProcesso,
			pub.NomeAdvogado,
			pub.OrigemIntimacao,
			pub.VariacaoEncontrada,
			pub.DataCadastro,
			pub.DataDisponibilizacao,
			pub.DataPublicacao,
			pub.InicioPrazo,
			pub.FinalPrazo,
			pub.Orgao,
			pub.UF,
			pub.Vara,
			pub.Cidade,
			pub.Jornal,
			pub.ProcessoEletronico,
			pub.ExpedidaConfirmada,  // Assuming fixed value
			pub.TipoIntimacao,       // Assuming fixed value
			pub.ListaProcEletronico, // Assuming fixed value
			pub.PerfilAdvogado,      // Assuming fixed value
			pub.Conteudo,
			pub.Anexos,
		}
		for colIndex, cellValue := range data {
			cell := fmt.Sprintf("%s%d", string('A'+colIndex), row)
			if err := f.SetCellValue("Relatório de intimações", cell, cellValue); err != nil {
				http.Error(w, "Error setting cell value", http.StatusInternalServerError)
				return
			}
		}
	}

	// Set active sheet
	f.SetActiveSheet(index)

	// Generate file name
	now := time.Now()
	fileName := fmt.Sprintf("relatorio_intimacoes_%s.xlsx", now.Format("20060102_150405"))

	// Save the file
	if err := f.SaveAs(fileName); err != nil {
		http.Error(w, "Error saving file", http.StatusInternalServerError)
		return
	}

	// Respond with file path
	response := map[string]string{
		"success":       "true",
		"pathExcel":     fileName,
		"tipoImpressao": "salvarExcel",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func main() {
	http.HandleFunc("/generate-report", generateExcelReport)
	fmt.Println("Server started at :3100")
	http.ListenAndServe(":3100", nil)
}
