package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ndeloof/mobylette/api"
)

func main() {
	log.Printf("Server started")

	server, err := api.NewServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	log.Fatal(http.ListenAndServe(":2375", api.NewRouter(server)))
}

