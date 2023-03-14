package main

import (
	"net/http"
	"time"
)

type Task struct {
	Address string
}

func getRequest(t *Task, outCh chan int, resCh chan int) {
	t1 := time.Now().Unix()
	res, _ := http.Get(t.Address)
	t2 := time.Now().Unix()
	delta := int(t2 - t1)
	outCh <- delta
	resCh <- res.StatusCode
}

type Scheduler

func main() {
	resCh := make(chan int)
	outCh := make(chan int)
	go getRequest()
}
