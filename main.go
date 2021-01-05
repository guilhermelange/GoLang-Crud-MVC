package main

import (
	"net/http"

	"gui.luizlange/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
