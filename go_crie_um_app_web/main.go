package main

import (
	"net/http"

	"github.com/lucasleitee/go_crie_um_app_web/routes"

	_ "github.com/lib/pq"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
