package handler

import (
	"api/infrastructure"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	server := infrastructure.NewServer()

	server.RunForServerless(w, r)
}
