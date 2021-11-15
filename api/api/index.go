package handler

import (
	"api/infra"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := infra.NewServer()

	server.RunForServerless(w, r)
}
