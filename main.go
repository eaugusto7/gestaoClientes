package main

import (
	"fmt"

	db "github.com/eaugusto7/gestaoClientes/database"
	"github.com/eaugusto7/gestaoClientes/routes"
)

func main() {
	db.ConectaBanco()
	fmt.Println("Iniciando Servidor...")
	routes.HandleRequest()
}
