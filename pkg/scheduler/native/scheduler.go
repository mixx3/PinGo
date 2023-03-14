package native

import (
	"PinGo/pkg/api"
	"github.com/google/uuid"
	"sync"
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
	tasks  map[uuid.UUID]Task
	inChs  map[uuid.UUID]chan *api.RequestPostSchema
	outChs map[uuid.UUID]chan *api.LogPostSchema
	Wg     *sync.WaitGroup
}

func NewScheduler(
	inChs *map[uuid.UUID]chan *api.RequestPostSchema,
	outChs *map[uuid.UUID]chan *api.LogPostSchema,
	wg *sync.WaitGroup,
) Scheduler {
	return &scheduler{
		tasks:  make(map[uuid.UUID]Task, 0),
		inChs:  *inChs,
		outChs: *outChs,
		Wg:     wg,
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
	s.Wg.Add(1)
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
	defer close(s.inChs[id])
	defer close(s.outChs[id])
	return nil
}

func (s *scheduler) StartAll() error {
	for _, tsk := range s.tasks {
		tsk.Start()
		s.Wg.Add(1)
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
