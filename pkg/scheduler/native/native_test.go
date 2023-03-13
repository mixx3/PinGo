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
	s := NewScheduler(&in, &out)
	schema := &api.RequestPostSchema{Address: "https://stackoverflow.com/questions/52719015/using-gin-gonic-and-some-scheduler-in-golang",
		StatusExpected: 200,
		RepeatTimeMs:   100,
		Name:           "test",
	}
	id, err := s.AddTask(schema)
	if err != nil {
		t.Fatal(err)
	}
	err = s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	ch := out[id]
	if err != nil {
		t.Fatal(err)
	}
	for {
		select {
		case data := <-ch:
			if err != nil {
				t.Fatal(err)
			}
			fmt.Println(data.ResponseTimeMs, data.Name)
			s.Stop(id)
			break
		}
	}
}
