package main

import (
	"fmt"

	"github.com/fmarga/api-go/database"
	"github.com/fmarga/api-go/routes"
)

func main() {
	database.ConnDatabase()
	fmt.Println("iniciando projeto")
	routes.HandleRequest()
}
