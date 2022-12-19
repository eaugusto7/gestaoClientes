package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
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

var ID_Login int

type JsonEdicaoLogin struct {
	Username string
}

func CriaLoginMock() {
	login := models.Login{Username: "usertest", Password: "senhateste"}
	hasher := md5.New()
	hasher.Write([]byte(login.Password))
	login.Password = hex.EncodeToString(hasher.Sum(nil))
	database.Database.Create(&login)
	ID_Login = int(login.Id)
}

func CriaLoginModel() models.Login {
	login := models.Login{Username: "usertest", Password: "senhateste"}
	return login
}

func DeletaLoginMock() {
	var login models.Login
	database.Database.Delete(&login, ID_Login)
}

func TestGetAllLogin(t *testing.T) {
	database.ConectaBanco()
	CriaLoginMock()
	defer DeletaLoginMock()

	r := SetupDasRotasDeTeste()
	r.GET("/api/v1/login", controllers.GetAllLogin)
	req, _ := http.NewRequest("GET", "/api/v1/login", nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestGetLoginById(t *testing.T) {
	database.ConectaBanco()
	CriaLoginMock()
	defer DeletaLoginMock()

	modeloLogin := CriaLoginModel()
	modelJson, _ := json.Marshal(modeloLogin)

	r := SetupDasRotasDeTeste()
	r.POST("/api/v1/login/", controllers.GetLoginById)
	req, _ := http.NewRequest("POST", "/api/v1/login/", bytes.NewBuffer(modelJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var loginMock models.Login
	json.Unmarshal(resposta.Body.Bytes(), &loginMock)
	assert.Equal(t, http.StatusOK, resposta.Code)

}

func TestInsertLogin(t *testing.T) {
	database.ConectaBanco()

	r := SetupDasRotasDeTeste()
	r.POST("/api/v1/login", controllers.InsertLogin)

	loginModelo := CriaLoginModel()

	jsonValue, _ := json.Marshal(loginModelo)
	req, _ := http.NewRequest("POST", "/api/v1/login", bytes.NewBuffer(jsonValue))
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

	//Deleta Login Mock Gerado
	database.Database.Delete(&loginModelo, mapResponse["Id"])
	assert.Equal(t, http.StatusOK, resposta.Code)
}

func TestUpdateLogin(t *testing.T) {
	database.ConectaBanco()
	CriaLoginMock()
	defer DeletaLoginMock()
	r := SetupDasRotasDeTeste()

	r.PATCH("/api/v1/login/"+strconv.Itoa(ID_Login), controllers.UpdateLogin)

	var modeloJson JsonEdicaoLogin
	modeloJson.Username = "Teste de Edição do User"
	valorJson, _ := json.Marshal(modeloJson)

	req, _ := http.NewRequest("PATCH", "/api/v1/login/"+strconv.Itoa(ID_Login), bytes.NewBuffer(valorJson))
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)

	var loginMockAtualizado models.Login
	json.Unmarshal(resposta.Body.Bytes(), &loginMockAtualizado)
}

func TestDeleteLogin(t *testing.T) {
	database.ConectaBanco()
	CriaLoginMock()
	r := SetupDasRotasDeTeste()
	r.DELETE("/api/v1/login/:id", controllers.DeleteLogin)
	req, _ := http.NewRequest("DELETE", "/api/v1/login/"+strconv.Itoa(ID_Cliente), nil)
	resposta := httptest.NewRecorder()
	r.ServeHTTP(resposta, req)
	assert.Equal(t, http.StatusOK, resposta.Code)
}
