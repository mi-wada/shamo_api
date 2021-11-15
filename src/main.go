package main

import (
	"api/infrastructure"
)

func main() {
	server := infrastructure.NewServer()
	server.Run(":8080")
}
