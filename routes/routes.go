package routes

import (
	"net/http"

	"github.com/Gabriel-Newton-dev/controllers"
)

func CarregaRotas() {
	http.HandleFunc("/", controllers.Index)
}
