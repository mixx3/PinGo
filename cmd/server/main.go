package main

import (
	api2 "PinGo/pkg/api"
	"PinGo/pkg/app"
	postgres2 "PinGo/pkg/repo/postgres"
	postgres3 "PinGo/pkg/service/postgres"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
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
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DB_DSN"), PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return err
	}
	defer func() {
		dbInstance, _ := db.DB()
		_ = dbInstance.Close()
	}()
	repo := postgres2.NewLogRepository(db)
	logService := postgres3.NewLogService(repo)
	requestService := postgres3.NewRequestService(postgres2.NewRequestRepository(db))
	receiverService := postgres3.NewReceiverService(postgres2.NewReceiverRepository(db))
	server := app.NewServer(router, logService, requestService, receiverService)
	err = server.Run()
	if err != nil {
		return err
	}
	return nil
}
