package main

import (
	"alta/project/config"
	"alta/project/routes"
	"os"

	"github.com/labstack/echo" //dont use v4
)

func main() {
	//deploy tugas
	e := echo.New()
	config.InitDb()
	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = ":80"
	}
	routes.New(e)
	e.Logger.Fatal(e.Start(port))

}
