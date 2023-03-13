package native

import (
	"PinGo/pkg/api"
	"fmt"
	"github.com/google/uuid"
	"testing"
)

func TestScheduler(t *testing.T) {
	out := make(map[uuid.UUID]chan *api.LogPostSchema)
	in := make(map[uuid.UUID]chan *api.RequestPostSchema)
	done := make(map[uuid.UUID]chan bool)
	s := NewScheduler(&in, &out, done)
	schema := &api.RequestPostSchema{Address: "http://localhost:8080/v1/status",
		StatusExpected: 200,
		RepeatTimeMs:   1000,
		Name:           "test",
	}
	id, err := s.AddTask(schema)
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		err := s.Start(id)
		if err != nil {

		}
	}()
	if err != nil {
		t.Fatal(err)
	}
	ch := out[id]
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		for {
			select {
			case data := <-ch:
				if err != nil {
					t.Fatal(err)
				}
				fmt.Println(data.ResponseTimeMs, data.Name, data.StatusCode)
			}
		}
	}()
	go s.Stop(id)
}
