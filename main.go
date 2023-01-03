package main

import (
	"github.com/prynnekey/go-reggie/config"
	"github.com/prynnekey/go-reggie/router"
)

func main() {
	config.InitConfig()

	router.InitRouter()
}
