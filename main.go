package main

import (
	"anohana-service/config"
	"anohana-service/service"
	"flag"
)

func main() {
	flag.Parse()
	if err := config.Init(); err != nil {
		panic(err)
	}

	if err := service.Run(); err != nil {
		panic(err)
	}

	return
}
