package native

import (
	"PinGo/pkg/api"
	"time"
)

type Task interface {
	Start() error
	Stop() error
}

type task struct {
	data   *api.RequestPostSchema
	outCh  chan *api.LogPostSchema
	ticker *time.Ticker
}

func NewTask(data *api.RequestPostSchema, outCh chan *api.LogPostSchema) Task {
	return &task{data: data, outCh: outCh}
}

func (t *task) Start() error {
	return nil
}

func (t *task) Stop() error {
	return nil
}
