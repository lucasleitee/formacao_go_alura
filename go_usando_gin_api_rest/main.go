package main

import (
	"github.com/lucasleitee/go_gin_api_rest/database"
	"github.com/lucasleitee/go_gin_api_rest/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequest()
}
