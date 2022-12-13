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

var ID_Servicos int

func CriaServicoMock() {
	servico := models.Servico{Nome: "Servico de Teste",
		Valor: 28.00,
		Tempo: 30,
	}
	db.DB.Create(&servico)
	ID_Servicos = int(servico.Id)
}

func CriaServicoModel() models.Servico {
	servico := models.Servico{Nome: "Servico de Teste",
		Valor: 28.00,
		Tempo: 30,
	}
	return servico
}

func DeletaServicoMock() {
	var servico models.Servico
	db.DB.Delete(&servico, ID_Servicos)
}

func TestGetAllServicos(t *testing.T) {
	db.ConectaBanco()
	CriaServicoMock()
	defer DeletaServicoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/servicos", controllers.GetAllServicos)
	req, _ := http.NewRequest("GET", "/api/v1/servicos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetServicoById(t *testing.T) {
	db.ConectaBanco()
	CriaServicoMock()
	defer DeletaServicoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/servicos/:id", controllers.GetByIdServicos)
	req, _ := http.NewRequest("GET", "/api/v1/servicos/"+strconv.Itoa(ID_Servicos), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var servicoMock models.Servico
	json.Unmarshal(resposta.Body.Bytes(), &servicoMock)

	assert.Equal(t, "Servico de Teste", servicoMock.Nome, " - Deveriam ter nomes iguais")
	assert.Equal(t, 30.00, servicoMock.Tempo)
	assert.Equal(t, 28.00, servicoMock.Valor)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestInsertServico(t *testing.T) {
	db.ConectaBanco()

	r := SetupDasRotasDeTeste()
	r.POST("/api/v1/servicos", controllers.InsertServicos)

	servicoModelo := CriaServicoModel()

	jsonValue, _ := json.Marshal(servicoModelo)
	req, _ := http.NewRequest("POST", "/api/v1/servicos", bytes.NewBuffer(jsonValue))
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
	db.DB.Delete(&servicoModelo, mapResponse["Id"])
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestUpdateServico(t *testing.T) {
	db.ConectaBanco()
	CriaServicoMock()
	defer DeletaServicoMock()
	r := SetupDasRotasDeTeste()

	r.PATCH("/api/v1/servicos/"+strconv.Itoa(ID_Servicos), controllers.UpdateServicos)
	servico := models.Servico{Nome: "Teste de Edicao do Nome - Servico"}
	valorJson, _ := json.Marshal(servico)

	req, _ := http.NewRequest("PATCH", "/api/v1/servicos/"+strconv.Itoa(ID_Servicos), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var servicoMockAtualizado models.Servico
	json.Unmarshal(resposta.Body.Bytes(), &servicoMockAtualizado)
}

func TestDeleteServico(t *testing.T) {
	db.ConectaBanco()
	CriaServicoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/api/v1/Servicos/:id", controllers.DeleteServicos)
	req, _ := http.NewRequest("DELETE", "/api/v1/Servicos/"+strconv.Itoa(ID_Servicos), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
