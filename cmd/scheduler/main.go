package main

import (
	"PinGo/pkg/api"
	"PinGo/pkg/repo/postgres"
	"PinGo/pkg/scheduler"
	pgs "PinGo/pkg/service/postgres"
	"github.com/jasonlvhit/gocron"
	"os"
)

func run() error {
	err := api.NewEnvParser(".env").Parse()
	if err != nil {
		return err
	}
	repo := postgres.NewRequestRepository(os.Getenv("DB_DSN"))
	srvc := pgs.NewRequestService(repo)
	sched := gocron.Scheduler{}
	s := scheduler.NewScheduler(srvc, sched)
	return nil
}
