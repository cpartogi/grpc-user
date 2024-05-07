package main

import (
	"log"
	"os"
	"user-service/cmd"
	"user-service/config"
)

func main() {
	cfg, err := config.Setup()
	if err != nil {
		log.Fatal("Cannot load config ", err.Error())
	}
	go func() {}()
	if cmd.Cli(cfg).Run(os.Args); err != nil {
		log.Fatal(err)
	}
	log.Println("Run")
}
