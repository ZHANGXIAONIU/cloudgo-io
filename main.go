package main

import (
	"github.com/ajian/cloudgo-io/server"
)

func main() {
	port := ":8080"
	server := server.NewServer()
	server.Run(port)
}
