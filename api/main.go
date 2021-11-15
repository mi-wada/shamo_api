package main

import (
	"api/infra"
)

func main() {
	server := infra.NewServer()
	server.Run(":8080")
}
