package main

import (
	api "PinGo/pkg/api"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type task struct {
	Name     string
	mtx      sync.Mutex
	schema   api.RequestPostSchema
	tickChan time.Ticker
}

func NewTask(name string,
	mtx sync.Mutex,
	schema api.RequestPostSchema,
	tc time.Ticker) Task {
	return &task{
		Name:     name,
		mtx:      mtx,
		schema:   schema,
		tickChan: tc,
	}
}

type Task interface {
	Process()
	Stop() error
}

func (t *task) Process() {
	tc := time.NewTicker(time.Duration(t.schema.RepeatTimeMs) * time.Second)
	defer tc.Stop()
	var forever chan struct{}
	select {
	case <-tc.C:
		res, _ := http.Get(t.schema.Address)
		fmt.Println(res.StatusCode)
	}
	<-forever
}

func (t *task) Stop() error {
	return nil
}

func main() {

}
