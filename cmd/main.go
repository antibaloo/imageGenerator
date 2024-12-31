package main

import (
	"flag"
	"log"

	"github.com/antibaloo/imageGenerator/configs"
	"github.com/antibaloo/imageGenerator/internal/server"
)

var confPath = flag.String("conf-path", "./configs/.env", "Path to config env.")

func main() {
	conf, err := configs.New(*confPath)
	if err != nil {
		log.Fatal(err)
	}
	server.Run(conf)
}
