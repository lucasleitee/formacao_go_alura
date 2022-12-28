package main

import (
	"fmt"

	"github.com/lucasleitee/formacao_go_alura/database"
	"github.com/lucasleitee/formacao_go_alura/models"
	"github.com/lucasleitee/formacao_go_alura/routes"
)

func main() {
	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBancoDeDados()
	fmt.Println("Iniciando Servidor")
	routes.HandleRequest()
}
