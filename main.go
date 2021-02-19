package main

import (
	"fmt"
	"github.com/ndeloof/mobylette/api"
	"log"
	"os"
)

func main() {
	log.Printf("server started")

	server, err := api.NewServer()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = server.Serve()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
