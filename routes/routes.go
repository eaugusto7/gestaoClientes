package routes

import (
	"log"
	"net/http"

	"github.com/eaugusto7/gestaoClientes/controllers"
	"github.com/eaugusto7/gestaoClientes/middleware"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.Home)

	//CRUD - Clientes
	r.HandleFunc("/api/v1/clientes/insert", controllers.InsertClient).Methods("Post")
	r.HandleFunc("/api/v1/clientes/getAll", controllers.GetAll).Methods("Get")
	r.HandleFunc("/api/v1/clientes/{id}", controllers.GetById).Methods("Get")
	r.HandleFunc("/api/v1/clientes/{id}", controllers.UpdateClient).Methods("Put")
	r.HandleFunc("/api/v1/clientes/{id}", controllers.DeleteClient).Methods("Delete")

	//CRUD - Servicos
	r.HandleFunc("/api/v1/servico/insert", controllers.InsertServicos).Methods("Post")
	r.HandleFunc("/api/v1/servico/getAll", controllers.GetAllServicos).Methods("Get")
	r.HandleFunc("/api/v1/servico/{id}", controllers.GetByIdServicos).Methods("Get")
	r.HandleFunc("/api/v1/servico/{id}", controllers.UpdateServicos).Methods("Put")
	r.HandleFunc("/api/v1/servico/{id}", controllers.DeleteServicos).Methods("Delete")

	r.Use(middleware.ContentTypeMiddleware)
	log.Fatal(http.ListenAndServe(":8000", handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
