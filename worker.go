package main

import (
	"log"
	"sync"
)

type queryData struct {
	data interface{}
}

type queryFunc func() queryData

type queryJob struct {
	query queryFunc
}

type worker struct {
	id       int
	workChan chan queryJob
	exitChan chan bool
}

type workerPool struct {
	mu         sync.RWMutex
	maxThreads int
	workers    []*worker
}

// func (w *worker) startWorker() {}

// func (w *worker) startWorker() {}

// func (wp *workerPool) ScheduleJob() {}

func NewWorker(id int) *worker {
	w := &worker{
		id:       id,
		workChan: make(chan queryJob),
		exitChan: make(chan bool),
	}
	log.Println("created worker", id)
	return w
}

func NewWorkerPool(maxThreads int) *workerPool {
	wp := &workerPool{
		maxThreads: maxThreads,
	}

	for i := 0; i < maxThreads; i++ {
		wp.workers = append(wp.workers, NewWorker(i))
	}

	return wp
}
