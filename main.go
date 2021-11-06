package main

import (
	"fmt"
	"log"

	"example.com/go/api"
	"example.com/go/util"
)

func main() {
	config, err := util.LoadConfig(".")

	if err != nil {
		log.Fatal("Cannot load config", err)
	}

	fmt.Println("hello")

	server, _ := api.NewServer()
	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("Cannot start server", err)
	}

}
