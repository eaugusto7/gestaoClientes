package controllers

import (
	"net/http"

	"crypto/md5"
	"encoding/hex"

	database "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/models"
	"github.com/gin-gonic/gin"
)

// GetAllLogin godoc
// @Summary Todos Usuários
// @Description Obtem todos os logins vindos do banco de dados
// @Tags Logins
// @Produce json
// @Sucess 200 [object] model.Login
// @Router /api/v1/login [get]
func GetAllLogin(context *gin.Context) {
	var login []models.Login
	var loginHelper []models.Login

	database.Database.Find(&login)

	//Define como vazio todos os passwords, para listar no front end não é necessário enviar a senha
	for _, loginInterator := range login {
		loginHelper = append(loginHelper, models.Login{Id: loginInterator.Id, Username: loginInterator.Username, Password: ""})
	}
	context.JSON(200, loginHelper)
}

//Obtem o json de um determinado login, filtrado por username
func GetLoginById(context *gin.Context) {
	var login models.Login
	var loginJson models.Login

	if error := context.ShouldBindJSON(&loginJson); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": error.Error()})
		return
	}
	//Busca no banco de dados através do username
	database.Database.Where(models.Login{Username: loginJson.Username}).FirstOrInit(&login)

	//Transforma o password que veio do JSON em código md5
	hasher := md5.New()
	hasher.Write([]byte(loginJson.Password))
	loginJson.Password = hex.EncodeToString(hasher.Sum(nil))

	//Compara se a senha vinda do JSON e a do Banco são as mesmas (Ambas devem ser Md5)
	if login.Password != loginJson.Password {
		context.JSON(http.StatusNotFound, gin.H{
			"Message: ": "Login/Senha incorreto"})
		return
	}
	context.JSON(http.StatusOK, login)
}

//Cria um novo login no banco de dados
func InsertLogin(context *gin.Context) {
	var login models.Login

	if error := context.ShouldBindJSON(&login); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": error.Error()})
		return
	}
	if error := models.ValidaDadosLogin(&login); error != nil {
		context.JSON(http.StatusBadGateway, gin.H{
			"Message error: ": error.Error()})
		return
	}
	//Transforma a senha em md5 antes de gravar no banco de dados
	hasher := md5.New()
	hasher.Write([]byte(login.Password))
	login.Password = hex.EncodeToString(hasher.Sum(nil))

	database.Database.Create(&login)
	context.JSON(http.StatusOK, login)
}

//Atualiza as informações de um determinado login no banco de dados
func UpdateLogin(context *gin.Context) {
	var login models.Login
	id := context.Params.ByName("id")
	database.Database.First(&login, id)

	if error := context.ShouldBindJSON(&login); error != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"Message error: ": error.Error()})
		return
	}

	//Transforma a senha em md5 antes de gravar no banco de dados
	hasher := md5.New()
	hasher.Write([]byte(login.Password))
	login.Password = hex.EncodeToString(hasher.Sum(nil))

	database.Database.Save(&login)
	context.JSON(http.StatusOK, login)
}

//Remove o login indicado pelo id no banco de dados
func DeleteLogin(context *gin.Context) {
	var login models.Login
	id := context.Params.ByName("id")
	database.Database.Delete(&login, id)

	context.JSON(http.StatusOK, gin.H{
		"Message": "Login Deletado"})
}
