package main

import (
	"testing"

	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
)

var ID_Cliente int
var ID_Atendimentos int
var ID_Servicos int

func CriaClienteMock() {
	cliente := models.Cliente{Nome: "Nome de Teste",
		Cpf:            "123.456.789-09",
		Rg:             "00.000.000",
		Email:          "emailteste@email.com",
		Telefone:       "(00) 0 0000 0000",
		Celular:        "(00) 0 0000 0000",
		Datanascimento: "01/01/2000",
		Sexo:           "Feminino",
		Profissao:      "Vendedora",
	}
	db.DB.Create(&cliente)
	ID_Cliente = int(cliente.Id)
}

func CriaAtendimentoMock() {
	atendimento := models.Atendimento{
		Nome:        "Atendimento Teste",
		Horario:     14.30,
		Idservico:   2,
		Idatendente: 1,
	}
	db.DB.Create(&atendimento)
	ID_Atendimentos = int(atendimento.Id)
}

func CriaServicoMock() {
	servicos := models.Servico{
		Nome:  "Servico Teste",
		Valor: 20,
		Tempo: 0.50,
	}
	db.DB.Create(&servicos)
	ID_Servicos = int(servicos.Id)
}

func TestGetAllClientes(t *testing.T) {

}
