package main

import (
	"PinGo/pkg/api"
	pg1 "PinGo/pkg/repo/postgres"
	"PinGo/pkg/scheduler"
	pgs "PinGo/pkg/service/postgres"
	"github.com/jasonlvhit/gocron"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

func run() error {
	err := api.NewEnvParser(".env").Parse()
	if err != nil {
		return err
	}
	db, err := gorm.Open(postgres.New(postgres.Config{DSN: os.Getenv("DB_DSN"), PreferSimpleProtocol: true}), &gorm.Config{})
	if err != nil {
		return err
	}
	repo := pg1.NewRequestRepository(db)
	srvc := pgs.NewRequestService(repo)
	sched := gocron.Scheduler{}
	s := scheduler.NewScheduler(srvc, sched)
	err = s.AddJob(&api.RequestPostSchema{Name: "fff"})
	if err != nil {
		return err
	}
	return nil
}
