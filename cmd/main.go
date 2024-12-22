package main

import (
	"github.com/Sinubio/mycalcservice/pkg/server"
)

func main() {
	app := server.New()
	app.RunServer()
}
