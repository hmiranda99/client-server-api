package main

import (
	"log"
	"os"

	"github.com/hmiranda99/client-server-api/client"
	"github.com/hmiranda99/client-server-api/server"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "server" {
		server.StartServer()
		return
	}

	if len(os.Args) > 1 && os.Args[1] == "client" {
		client.StartClient()
		return
	}

	log.Fatal("Use 'server' ou 'client' como argumento para iniciar o servidor ou o cliente.")
}
