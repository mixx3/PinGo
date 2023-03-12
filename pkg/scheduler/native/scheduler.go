package native

import (
	"PinGo/pkg/api"
	"github.com/google/uuid"
)

type Scheduler interface {
	TasksActive() []uuid.UUID
	Start(id uuid.UUID) error
	StartAll() error
	Stop(id uuid.UUID) error
	NewTask(tsk *api.RequestPostSchema) (uuid.UUID, error)
	StopAll() error
}

type scheduler struct {
	tasks  map[uuid.UUID]Task
	inChs  map[uuid.UUID]chan *api.RequestPostSchema
	outChs map[uuid.UUID]chan *api.LogPostSchema
}

func (s *scheduler) TasksActive() []uuid.UUID {
	res := []uuid.UUID{}
	for k, _ := range s.tasks {
		res = append(res, k)
	}
	return res
}

func (s *scheduler) Start(id uuid.UUID) error {
	err := s.tasks[id].Start()
	return err
}

func (s *scheduler) NewTask(tsk *api.RequestPostSchema) (uuid.UUID, error) {
	uid := uuid.New()
	s.outChs[uid] = make(chan *api.LogPostSchema)
	nt := NewTask(tsk, s.outChs[uid])
	s.tasks[uid] = nt
	s.inChs[uid] <- tsk
	return uid, nil
}

func (s *scheduler) Stop(id uuid.UUID) error {
	err := s.tasks[id].Stop()
	return err
}

func (s *scheduler) StartAll() error {
	for _, tsk := range s.tasks {
		err := tsk.Start()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *scheduler) StopAll() error {
	for _, tsk := range s.tasks {
		err := tsk.Stop()
		if err != nil {
			return err
		}
	}
	return nil
}
