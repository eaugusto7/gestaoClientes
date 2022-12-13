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
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

var ID_Cliente int

/*var ID_Atendimentos int
var ID_Servicos int*/

func SetupDasRotasDeTeste() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	rotas := gin.Default()
	return rotas
}

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

func CriaClienteModel() models.Cliente {
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
	return cliente
}

func DeletaClienteMock() {
	var cliente models.Cliente
	db.DB.Delete(&cliente, ID_Cliente)
}

/*func CriaServicoMock() {
	servicos := models.Servico{
		Nome:  "Servico Teste",
		Valor: 20,
		Tempo: 0.50,
	}
	db.DB.Create(&servicos)
	ID_Servicos = int(servicos.Id)
}
*/
func TestGetAllClientes(t *testing.T) {
	db.ConectaBanco()
	CriaClienteMock()

	defer DeletaClienteMock()

	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/clientes", controllers.GetAllClientes)
	req, _ := http.NewRequest("GET", "/api/v1/clientes", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetClienteById(t *testing.T) {
	db.ConectaBanco()
	CriaClienteMock()
	defer DeletaClienteMock()
	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/clientes/:id", controllers.GetByIdClientes)
	req, _ := http.NewRequest("GET", "/api/v1/clientes/"+strconv.Itoa(ID_Cliente), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var clienteMock models.Cliente
	json.Unmarshal(resposta.Body.Bytes(), &clienteMock)

	assert.Equal(t, "Nome de Teste", clienteMock.Nome, " - Deveriam ter nomes iguais")
	assert.Equal(t, "(00) 0 0000 0000", clienteMock.Celular)
	assert.Equal(t, "123.456.789-09", clienteMock.Cpf)
	assert.Equal(t, "01/01/2000", clienteMock.Datanascimento)
	assert.Equal(t, "emailteste@email.com", clienteMock.Email)

	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestInsertCliente(t *testing.T) {
	db.ConectaBanco()

	r := SetupDasRotasDeTeste()
	r.POST("/api/v1/clientes", controllers.InsertClient)

	clienteModelo := CriaClienteModel()

	jsonValue, _ := json.Marshal(clienteModelo)
	req, _ := http.NewRequest("POST", "/api/v1/clientes", bytes.NewBuffer(jsonValue))
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
	db.DB.Delete(&clienteModelo, mapResponse["Id"])
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestUpdateCliente(t *testing.T) {
	db.ConectaBanco()
	CriaClienteMock()
	defer DeletaClienteMock()
	r := SetupDasRotasDeTeste()

	r.PATCH("/api/v1/clientes/:id", controllers.UpdateClient)
	cliente := models.Cliente{Nome: "Teste de Edicao do Nome"}
	valorJson, _ := json.Marshal(cliente)

	req, _ := http.NewRequest("PATCH", "/api/v1/clientes/"+strconv.Itoa(ID_Cliente), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var clienteMockAtualizado models.Cliente
	json.Unmarshal(resposta.Body.Bytes(), &clienteMockAtualizado)
}

func TestDeleteCliente(t *testing.T) {
	db.ConectaBanco()
	CriaClienteMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/api/v1/clientes/:id", controllers.DeleteClient)
	req, _ := http.NewRequest("DELETE", "/api/v1/clientes/"+strconv.Itoa(ID_Cliente), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
