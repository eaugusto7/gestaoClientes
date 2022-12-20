package routes

import (
	"log"
	"net/http"

	"github.com/eaugusto7/gestaoClientes/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"

	docs "github.com/eaugusto7/gestaoClientes/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func HandleRequest() {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	r.Static("/assets", "./assets")

	//CRUD Clientes
	r.POST("/api/v1/clientes", controllers.InsertCliente)
	r.GET("/api/v1/clientes", controllers.GetAllClientes)
	r.GET("/api/v1/clientes/:id", controllers.GetClienteById)
	r.GET("/api/v1/clientes/nome/:nome", controllers.GetClienteByName)
	r.PATCH("/api/v1/clientes/:id", controllers.UpdateCliente)
	r.DELETE("/api/v1/clientes/:id", controllers.DeleteCliente)

	//CRUD - Servicos
	r.POST("/api/v1/servicos", controllers.InsertServico)
	r.GET("/api/v1/servicos", controllers.GetAllServicos)
	r.GET("/api/v1/servicos/:id", controllers.GetServicosById)
	r.PATCH("/api/v1/servicos/:id", controllers.UpdateServico)
	r.DELETE("/api/v1/servicos/:id", controllers.DeleteServico)

	//CRUD - Atendimentos
	r.POST("/api/v1/atendimentos", controllers.InsertAtendimento)
	r.GET("/api/v1/atendimentos", controllers.GetAllAtendimentos)
	r.GET("/api/v1/atendimentos/:id", controllers.GetAtendimentoById)
	r.GET("/api/v1/atendimentos/clientes/:idcliente", controllers.GetAtendimentoByClienteId)
	r.GET("/api/v1/atendimentos/atendentes/:idatendente", controllers.GetAtendimentoByAtendenteId)
	r.GET("/api/v1/atendimentos/servicos/:idservico", controllers.GetAtendimentoByServicoId)
	r.PATCH("/api/v1/atendimentos/:id", controllers.UpdateAtendimento)
	r.DELETE("/api/v1/atendimentos/:id", controllers.DeleteAtendimento)

	//CRUD - Atendente
	r.POST("/api/v1/atendentes", controllers.InsertAtendente)
	r.GET("/api/v1/atendentes", controllers.GetAllAtendentes)
	r.GET("/api/v1/atendentes/:id", controllers.GetAtendenteById)
	r.GET("/api/v1/atendentes/nome/:nome", controllers.GetAtendenteByName)
	r.PATCH("/api/v1/atendentes/:id", controllers.UpdateAtendente)
	r.DELETE("/api/v1/atendentes/:id", controllers.DeleteAtendente)

	//CRUD - Produtos
	r.POST("/api/v1/produtos", controllers.InsertProduto)
	r.GET("/api/v1/produtos", controllers.GetAllProdutos)
	r.GET("/api/v1/produtos/:id", controllers.GetProdutoById)
	r.PATCH("/api/v1/produtos/:id", controllers.UpdateProduto)
	r.DELETE("/api/v1/produtos/:id", controllers.DeleteProduto)

	//CRUD - Login
	r.POST("/api/v1/login", controllers.InsertLogin)
	r.GET("/api/v1/login", controllers.GetAllLogin)
	r.POST("/api/v1/login/", controllers.GetLoginById)
	r.PATCH("/api/v1/login/:id", controllers.UpdateLogin)
	r.DELETE("/api/v1/login/:id", controllers.DeleteLogin)

	//CRUD - Quadro de Horários
	r.POST("/api/v1/horarios", controllers.InsertQuadroHorario)
	r.GET("/api/v1/horarios", controllers.GetAllHorarios)
	r.GET("/api/v1/horarios/:id", controllers.GetHorarioById)
	r.GET("/api/v1/horarios/atendente/:idatendente", controllers.GetHorarioByAtendente)
	r.PATCH("/api/v1/horarios/:id", controllers.UpdateHorarios)
	r.DELETE("/api/v1/horarios/:id", controllers.DeleteHorario)

	//Define para a utilizar a porta 8090
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	log.Fatal(http.ListenAndServe(":8090", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
