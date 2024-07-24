package main

import (
	"log"
	"net/http"
	_ "net/http/pprof"
	"project1/config"
	"project1/routes"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")    
    }

    config.ConnectDatabase()

    e := echo.New()

    routes.Init(e)

    go func() { 
        log.Println(http.ListenAndServe("localhost:6060", nil))
    }()

    e.Logger.Fatal(e.Start(":8080"))
}
