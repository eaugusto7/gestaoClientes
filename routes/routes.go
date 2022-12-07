package routes

import (
	"log"
	"net/http"

	"github.com/eaugusto7/gestaoClientes/controllers"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func HandleRequest() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.Static("/assets", "./assets")

	//r := mux.NewRouter()
	//r.HandleFunc("/", controllers.Home)

	r.GET("/api/v1/clientes/getAll", controllers.GetAll)

	//CRUD - Clientes
	/*r.HandleFunc("/api/v1/clientes/insert", controllers.InsertClient).Methods("Post")
	r.HandleFunc("/api/v1/clientes/getAll", controllers.GetAll).Methods("Get")
	r.HandleFunc("/api/v1/clientes/{id}", controllers.GetById).Methods("Get")
	r.HandleFunc("/api/v1/clientes/{id}", controllers.UpdateClient).Methods("Put")
	r.HandleFunc("/api/v1/clientes/{id}", controllers.DeleteClient).Methods("Delete")

	//CRUD - Servicos
	r.HandleFunc("/api/v1/servicos/insert", controllers.InsertServicos).Methods("Post")
	r.HandleFunc("/api/v1/servicos/getAll", controllers.GetAllServicos).Methods("Get")
	r.HandleFunc("/api/v1/servicos/{id}", controllers.GetByIdServicos).Methods("Get")
	r.HandleFunc("/api/v1/servicos/{id}", controllers.UpdateServicos).Methods("Put")
	r.HandleFunc("/api/v1/servicos/{id}", controllers.DeleteServicos).Methods("Delete")

	//CRUD - Atendimentos
	r.HandleFunc("/api/v1/atendimentos/insert", controllers.InsertAtendimentos).Methods("Post")
	r.HandleFunc("/api/v1/atendimentos/getAll", controllers.GetAllAtendimentos).Methods("Get")
	r.HandleFunc("/api/v1/atendimentos/{id}", controllers.GetByIdAtendimentos).Methods("Get")
	r.HandleFunc("/api/v1/atendimentos/{id}", controllers.UpdateAtendimentos).Methods("Put")
	r.HandleFunc("/api/v1/atendimentos/{id}", controllers.DeleteAtendimento).Methods("Delete")*/

	//r.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
