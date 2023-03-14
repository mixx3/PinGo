package native

import (
	"PinGo/pkg/api"
	"bytes"
	"errors"
	"log"
	"net/http"
	"sync"
	"time"
)

type Task interface {
	Start() chan bool
	Stop() error
}

type task struct {
	data   *api.RequestPostSchema
	outCh  chan *api.LogPostSchema
	ticker *time.Ticker
	done   chan bool
	mutex  sync.Mutex
}

func NewTask(data *api.RequestPostSchema, outCh chan *api.LogPostSchema) Task {
	t := &task{
		data:  data,
		outCh: outCh,
	}
	return t
}

func (t *task) Start() chan bool {
	t.done = make(chan bool, 1)
	go func() {
		t.ticker = time.NewTicker(time.Duration(t.data.RepeatTimeMs) * time.Millisecond)
		defer t.ticker.Stop()
		for {
			select {
			case <-t.done:
				return
			case <-t.ticker.C:
				go t.process()
			}
		}
	}()
	return t.done
}

func (t *task) process() {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	delta := 0
	res, err := &http.Response{}, errors.New("No method")
	switch t.data.Method {
	case "GET":
		t1 := time.Now().Unix()
		res, err = http.Get(t.data.Address)
		t2 := time.Now().Unix()
		delta = int(t2 - t1)
	case "POST":
		t1 := time.Now().Unix()
		data := []byte(t.data.Body)
		r := bytes.NewReader(data)
		res, err = http.Post(
			t.data.Address,
			"application/json",
			r,
		)
		t2 := time.Now().Unix()
		delta = int(t2 - t1)
	}
	if err != nil {
		log.Fatalln(err)
	}
	if res.StatusCode != t.data.StatusExpected {
		t.outCh <- &api.LogPostSchema{
			StatusCode:     res.StatusCode,
			Address:        t.data.Address,
			Name:           t.data.Name,
			ReceiverID:     t.data.ReceiverID,
			ResponseTimeMs: delta,
		}
	}
}

func (t *task) Stop() error {
	t.ticker.Stop()
	<-t.done
	return nil
}
