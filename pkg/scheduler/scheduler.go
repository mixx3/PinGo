package scheduler

import (
	"PinGo/pkg/api"
	"fmt"
	"github.com/jasonlvhit/gocron"
)

type scheduler struct {
	service     api.RequestService
	gcScheduler gocron.Scheduler
}

func NewScheduler(service api.RequestService, gcScheduler gocron.Scheduler) api.Scheduler {
	return &scheduler{service, gcScheduler}
}

func (s *scheduler) AddJob(schema *api.RequestPostSchema) error {
	gs := gocron.NewScheduler()
	err := gs.Every(uint64(schema.RepeatTimeMs)/10).Seconds().Do(task, schema)
	return err
}

func task(schema *api.RequestPostSchema) {
	fmt.Println(schema.Name)
}
