package main

import (
	"fmt"

	db "github.com/eaugusto7/gestaoClientes/database"
)

func main() {
	fmt.Println("Deu certo")

	db.ConectaBanco()

}
