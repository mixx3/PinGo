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
	AddTask(tsk *api.RequestPostSchema) (uuid.UUID, error)
	StopAll() error
}

type scheduler struct {
	tasks   map[uuid.UUID]Task
	inChs   map[uuid.UUID]chan *api.RequestPostSchema
	outChs  map[uuid.UUID]chan *api.LogPostSchema
	doneChs map[uuid.UUID]chan bool
}

func NewScheduler(
	inChs *map[uuid.UUID]chan *api.RequestPostSchema,
	outChs *map[uuid.UUID]chan *api.LogPostSchema,
	doneChs map[uuid.UUID]chan bool,
) Scheduler {
	return &scheduler{
		tasks:   make(map[uuid.UUID]Task, 0),
		inChs:   *inChs,
		outChs:  *outChs,
		doneChs: doneChs,
	}
}

func (s *scheduler) TasksActive() []uuid.UUID {
	var res []uuid.UUID
	for k, _ := range s.tasks {
		res = append(res, k)
	}
	return res
}

func (s *scheduler) Start(id uuid.UUID) error {
	dc := s.tasks[id].Start()
	s.doneChs[id] = dc
	return nil
}

func (s *scheduler) AddTask(tsk *api.RequestPostSchema) (uuid.UUID, error) {
	uid := uuid.New()
	s.outChs[uid] = make(chan *api.LogPostSchema)
	nt := NewTask(tsk, s.outChs[uid])
	s.tasks[uid] = nt
	return uid, nil
}

func (s *scheduler) Stop(id uuid.UUID) error {
	defer close(s.doneChs[id])
	defer close(s.inChs[id])
	defer close(s.outChs[id])
	defer close(s.doneChs[id])
	return nil
}

func (s *scheduler) StartAll() error {
	for uid, tsk := range s.tasks {
		dc := tsk.Start()
		s.doneChs[uid] = dc
	}
	return nil
}

func (s *scheduler) StopAll() error {
	for uid, _ := range s.tasks {
		err := s.Stop(uid)
		if err != nil {
			return err
		}
	}
	return nil
}
