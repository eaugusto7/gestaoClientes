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

var ID_Produto int

func CriaProdutoMock() {
	produto := models.Produtos{Nome: "Produto de Teste",
		Quantidade: 43,
		Descricao:  "Azul",
		Fabricante: "Castell",
		Valorcusto: 1.50,
		Valorvenda: 2.50,
	}
	db.DB.Create(&produto)
	ID_Produto = int(produto.Id)
}

func CriaProdutoModel() models.Produtos {
	produto := models.Produtos{Nome: "Produto de Teste",
		Quantidade: 43,
		Descricao:  "Azul",
		Fabricante: "Castell",
		Valorcusto: 1.50,
		Valorvenda: 2.50,
	}
	return produto
}

func DeletaProdutoMock() {
	var produto models.Produtos
	db.DB.Delete(&produto, ID_Produto)
}

func TestGetAllProduto(t *testing.T) {
	db.ConectaBanco()
	CriaProdutoMock()
	defer DeletaProdutoMock()

	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/produtos", controllers.GetAllProdutos)
	req, _ := http.NewRequest("GET", "/api/v1/produtos", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetProdutosById(t *testing.T) {
	db.ConectaBanco()
	CriaProdutoMock()
	defer DeletaProdutoMock()
	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/produtos/:id", controllers.GetProdutoById)
	req, _ := http.NewRequest("GET", "/api/v1/produtos/"+strconv.Itoa(ID_Produto), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var produtoMock models.Produtos
	json.Unmarshal(resposta.Body.Bytes(), &produtoMock)

	assert.Equal(t, "Produto de Teste", produtoMock.Nome, " - Deveriam ter nomes iguais")
	assert.Equal(t, 43, produtoMock.Quantidade)
	assert.Equal(t, "Azul", produtoMock.Descricao)
	assert.Equal(t, "Castell", produtoMock.Fabricante)
	assert.Equal(t, 1.50, produtoMock.Valorcusto)
	assert.Equal(t, 2.50, produtoMock.Valorvenda)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestInsertProduto(t *testing.T) {
	db.ConectaBanco()

	r := SetupDasRotasDeTeste()
	r.POST("/api/v1/produtos", controllers.InsertProduto)

	produtoModelo := CriaProdutoModel()

	jsonValue, _ := json.Marshal(produtoModelo)
	req, _ := http.NewRequest("POST", "/api/v1/produtos", bytes.NewBuffer(jsonValue))
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
	db.DB.Delete(&produtoModelo, mapResponse["Id"])
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestUpdateProduto(t *testing.T) {
	db.ConectaBanco()
	CriaProdutoMock()
	defer DeletaProdutoMock()
	r := SetupDasRotasDeTeste()

	r.PATCH("/api/v1/produtos/:id", controllers.UpdateProduto)
	produto := models.Atendimento{Nome: "Teste de Edicao do Nome - Chave de Fenda"}
	valorJson, _ := json.Marshal(produto)
	pathEdit := "/api/v1/produtos/" + strconv.Itoa(ID_Produto)

	req, _ := http.NewRequest("PATCH", pathEdit, bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var produtoMockAtualizado models.Produtos
	json.Unmarshal(resposta.Body.Bytes(), &produtoMockAtualizado)
	assert.Equal(t, "Teste de Edicao do Nome - Chave de Fenda", produtoMockAtualizado.Nome)

}

func TestDeleteProduto(t *testing.T) {
	db.ConectaBanco()
	CriaProdutoMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/api/v1/produtos/:id", controllers.DeleteProduto)
	req, _ := http.NewRequest("DELETE", "/api/v1/produtos/"+strconv.Itoa(ID_Produto), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
