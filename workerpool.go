package main

/*
import (
	"log"
	"sync"
)

type workerPool struct {
	mu         sync.Mutex
	workers    []*worker
	queries    chan QueryJob
	maxThreads int
	rateLimit  int
}

func (wp *workerPool) scheduleWorker(q QueryJob) (worker, uint64) {
	wp.mu.Lock()
	defer wp.mu.Unlock()

	return wp.workers[workerNum], workerNum
}

func (wp *workerPool) startPool() {
	for task := range wp.queries {
		worker, num := wp.scheduleWorker(task)

		if err := worker.Do(task); err != nil {
			p.logger.Errorf(err, `worker #%d doing task "%s"`, num, task.GetID())
		}
	}
}

func (wp *workerPool) NewWorker(id int) *worker {
	wp.mu.Lock()
	defer wp.mu.Unlock()
	log.Println("created worker", id)
	wp.workers = append(wp.workers)
	return w
}

func NewWorkerPool(maxThreads int) *workerPool {
	wp := &workerPool{
		queries:    make(chan QueryJob),
		maxThreads: maxThreads,
	}

	go wp.startPool()

	return wp
}
*/
