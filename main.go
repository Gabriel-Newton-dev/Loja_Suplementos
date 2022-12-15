package main

import (
	//"modulos/models"
	"net/http"

	"github.com/Gabriel-Newton-dev/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
