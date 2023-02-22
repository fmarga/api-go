package main

import (
	"fmt"

	"github.com/fmarga/api-go/database"
	"github.com/fmarga/api-go/models"
	"github.com/fmarga/api-go/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "Name 1", History: "History 1"},
		{Name: "Name 2", History: "History 2"},
	}

	database.ConnDatabase()
	fmt.Println("iniciando projeto")
	routes.HandleRequest()
}
