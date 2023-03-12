package native

import (
	"PinGo/pkg/api"
	"log"
	"net/http"
	"sync"
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
	done   chan bool
	mutex  sync.Mutex
}

func NewTask(data *api.RequestPostSchema, outCh chan *api.LogPostSchema) Task {
	t := &task{
		data:   data,
		outCh:  outCh,
		ticker: time.NewTicker(time.Duration(data.RepeatTimeMs) * time.Second),
	}
	return t
}

func (t *task) Start() error {
	go func() {
		for {
			select {
			case <-t.ticker.C:
				go func() {
					t.mutex.Lock()
					defer t.mutex.Unlock()
					t1 := time.Now().Unix()
					res, err := http.Get(t.data.Address)
					t2 := time.Now().Unix()
					if err != nil {
						log.Fatalln(err)
					}
					if res.StatusCode != t.data.StatusExpected {
						t.outCh <- &api.LogPostSchema{
							StatusCode:     res.StatusCode,
							Address:        t.data.Address,
							Name:           t.data.Name,
							ReceiverID:     t.data.ReceiverID,
							ResponseTimeMs: int(t2 - t1),
						}
					}
				}()
			case <-t.done:
				return
			}
		}
	}()
	return nil
}

func (t *task) Stop() error {
	<-t.done
	return nil
}
