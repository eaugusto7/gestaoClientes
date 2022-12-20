package main

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/eaugusto7/gestaoClientes/controllers"
	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/stretchr/testify/assert"
)

var ID_Atendimentos int

func CriaAtendimentoMock() {
	atendimento := models.Atendimento{Nome: "Atendimento de Teste",
		Horario:        14.00,
		Idservico:      2,
		Idatendente:    1,
		Idcliente:      100000,
		Status:         "Em Andamento",
		Statusfixo:     false,
		Formapagamento: "Não Informado",
	}
	database.Database.Create(&atendimento)
	ID_Atendimentos = int(atendimento.Id)
}

func CriaAtendimentoModel() models.Atendimento {
	atendimento := models.Atendimento{Nome: "Atendimento de Teste",
		Horario:        14.00,
		Idservico:      2,
		Idatendente:    1,
		Idcliente:      100000,
		Status:         "Em Andamento",
		Statusfixo:     false,
		Formapagamento: "Não Informado",
	}
	return atendimento
}

func DeletaAtendimentoMock() {
	var atendimento models.Atendimento
	database.Database.Delete(&atendimento, ID_Atendimentos)
}

func TestGetAllAtendimentos(t *testing.T) {
	database.ConectaBanco()
	CriaAtendimentoMock()
	defer DeletaAtendimentoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/atendimentos", controllers.GetAllAtendimentos)
	req, _ := http.NewRequest("GET", "/api/v1/atendimentos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetAtendimentoById(t *testing.T) {
	database.ConectaBanco()
	CriaAtendimentoMock()
	defer DeletaAtendimentoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/atendimentos/:id", controllers.GetAtendimentoById)
	req, _ := http.NewRequest("GET", "/api/v1/atendimentos/"+strconv.Itoa(ID_Atendimentos), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var atendimentoMock models.Atendimento
	json.Unmarshal(resposta.Body.Bytes(), &atendimentoMock)

	assert.Equal(t, "Atendimento de Teste", atendimentoMock.Nome, " - Deveriam ter nomes iguais")
	assert.Equal(t, 14.00, atendimentoMock.Horario)
	assert.Equal(t, 1, atendimentoMock.Idatendente)
	assert.Equal(t, 2, atendimentoMock.Idservico)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetAtendimentoByClienteId(t *testing.T) {
	database.ConectaBanco()
	CriaAtendimentoMock()
	defer DeletaAtendimentoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/atendimentos/clientes/:id", controllers.GetAtendimentoByClienteId)
	req, _ := http.NewRequest("GET", "/api/v1/atendimentos/clientes/"+strconv.Itoa(100000), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var atendimentoMock models.Atendimento
	json.Unmarshal(resposta.Body.Bytes(), &atendimentoMock)

	assert.Equal(t, "Atendimento de Teste", atendimentoMock.Nome, " - Deveriam ter nomes iguais")
	assert.Equal(t, 14.00, atendimentoMock.Horario)
	assert.Equal(t, 1, atendimentoMock.Idatendente)
	assert.Equal(t, 2, atendimentoMock.Idservico)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestInsertAtendimento(t *testing.T) {
	database.ConectaBanco()

	r := SetupDasRotasDeTeste()
	r.POST("/api/v1/atendimentos", controllers.InsertAtendimento)

	atendimentoModelo := CriaAtendimentoModel()

	jsonValue, _ := json.Marshal(atendimentoModelo)
	req, _ := http.NewRequest("POST", "/api/v1/atendimentos", bytes.NewBuffer(jsonValue))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	//Converte Response para Map
	bodyBytes, err := io.ReadAll(resposta.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)

	mapResponse := map[string]interface{}{}
	if err := json.Unmarshal([]byte(bodyString), &mapResponse); err != nil {
		panic(err)
	}

	//Deleta ClienteMock gerado
	database.Database.Delete(&atendimentoModelo, mapResponse["Id"])
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestUpdateAtendimento(t *testing.T) {
	database.ConectaBanco()
	CriaAtendimentoMock()
	defer DeletaAtendimentoMock()
	r := SetupDasRotasDeTeste()
	r.PATCH("/api/v1/atendimentos/"+strconv.Itoa(ID_Atendimentos), controllers.UpdateAtendimento)

	var modeloJson JsonEdicao
	modeloJson.Nome = "Teste de Edição do Nome"
	valorJson, _ := json.Marshal(modeloJson)

	req, _ := http.NewRequest("PATCH", "/api/v1/atendimentos/"+strconv.Itoa(ID_Atendimentos), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var atendimentoMockAtualizado models.Atendimento
	json.Unmarshal(resposta.Body.Bytes(), &atendimentoMockAtualizado)
}

func TestDeleteAtendimento(t *testing.T) {
	database.ConectaBanco()
	CriaAtendimentoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/api/v1/atendimentos/:id", controllers.DeleteAtendimento)
	req, _ := http.NewRequest("DELETE", "/api/v1/atendimentos/"+strconv.Itoa(ID_Atendimentos), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
