package main

import (
	api2 "PinGo/pkg/api"
	"PinGo/pkg/app"
	pg "PinGo/pkg/repo/postgres"
	services "PinGo/pkg/service/postgres"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"os"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "this is the startup error: %s\n", err)
		os.Exit(1)
	}
}

func run() error {
	err := api2.NewEnvParser(".env").Parse()
	if err != nil {
		return err
	}
	router := gin.Default()
	router.Use(cors.Default())
	repo := pg.NewLogRepository(os.Getenv("DB_DSN"))
	logService := services.NewLogService(repo)
	requestService := services.NewRequestService(pg.NewRequestRepository("DB_DSN"))
	server := app.NewServer(router, logService, requestService)
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}
