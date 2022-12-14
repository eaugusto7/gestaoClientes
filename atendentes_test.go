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
	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/stretchr/testify/assert"
)

var ID_Atendente int

func CriaAtendenteMock() {
	atendente := models.Atendente{Nome: "Atendente de Teste",
		Telefone: "(00) 0 0000 0000",
	}
	db.DB.Create(&atendente)
	ID_Atendente = int(atendente.Id)
}

func CriaAtendenteModel() models.Atendente {
	atendente := models.Atendente{Nome: "Atendente de Teste",
		Telefone: "(00) 0 0000 0000",
	}
	return atendente
}

func DeletaAtendenteMock() {
	var atendente models.Atendente
	db.DB.Delete(&atendente, ID_Atendente)
}

func TestGetAllAtendente(t *testing.T) {
	db.ConectaBanco()
	CriaAtendenteMock()
	defer DeletaAtendenteMock()

	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/atendentes", controllers.GetAllAtendente)
	req, _ := http.NewRequest("GET", "/api/v1/atendentes", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetAtendenteById(t *testing.T) {
	db.ConectaBanco()
	CriaAtendenteMock()
	defer DeletaAtendenteMock()
	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/atendentes/:id", controllers.GetAtendenteById)
	req, _ := http.NewRequest("GET", "/api/v1/atendentes/"+strconv.Itoa(ID_Atendente), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var atendenteMock models.Atendente
	json.Unmarshal(resposta.Body.Bytes(), &atendenteMock)

	assert.Equal(t, "Atendente de Teste", atendenteMock.Nome, " - Deveriam ter nomes iguais")
	assert.Equal(t, "(00) 0 0000 0000", atendenteMock.Telefone)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestInsertAtendente(t *testing.T) {
	db.ConectaBanco()

	r := SetupDasRotasDeTeste()
	r.POST("/api/v1/atendentes", controllers.InsertAtendente)

	atendenteModelo := CriaAtendenteModel()

	jsonValue, _ := json.Marshal(atendenteModelo)
	req, _ := http.NewRequest("POST", "/api/v1/atendentes", bytes.NewBuffer(jsonValue))
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
	db.DB.Delete(&atendenteModelo, mapResponse["Id"])
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestUpdateAtendente(t *testing.T) {
	db.ConectaBanco()
	CriaAtendenteMock()
	defer DeletaAtendenteMock()
	r := SetupDasRotasDeTeste()

	r.PATCH("/api/v1/atendentes/"+strconv.Itoa(ID_Atendente), controllers.UpdateAtendente)
	atendente := models.Atendente{Nome: "Teste de Edicao do Nome - Atendente"}
	valorJson, _ := json.Marshal(atendente)

	req, _ := http.NewRequest("PATCH", "/api/v1/atendentes/"+strconv.Itoa(ID_Atendente), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var atendenteMockAtualizado models.Atendente
	json.Unmarshal(resposta.Body.Bytes(), &atendenteMockAtualizado)
}

func TestDeleteAtendente(t *testing.T) {
	db.ConectaBanco()
	CriaAtendenteMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/api/v1/atendentes/:id", controllers.DeleteAtendente)
	req, _ := http.NewRequest("DELETE", "/api/v1/atendentes/"+strconv.Itoa(ID_Atendente), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
