package scheduler

import (
	"PinGo/pkg/api"
	"fmt"
	"github.com/jasonlvhit/gocron"
)

type Scheduler struct {
	service   api.RequestService
	scheduler gocron.Scheduler
}

func (s *Scheduler) AddJob(schema api.RequestPostSchema) error {
	gs := gocron.NewScheduler()
	err := gs.Every(uint64(schema.RepeatTimeMs)/10).Seconds().Do(task, schema)
	return err
}

func task(schema api.RequestPostSchema) {
	fmt.Println(schema.Name)
}
