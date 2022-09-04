package main

import (
	"fmt"
	"log"
	"sync"
)

type queryData struct {
	data interface{}
}

type queryFunc func() queryData

type QueryJob struct {
	query    queryFunc
	dataChan chan queryData
}

type worker struct {
	id       int
	workChan chan QueryJob
	exitChan chan bool
}

type workerPool struct {
	mu             sync.RWMutex
	maxThreads     int
	maxConcurrency int
	workers        []*worker
}

func (w *worker) scheduleWork(jobs <-chan QueryJob, results chan<- queryData) {
	for job := range jobs {
		fmt.Println("worker", w.id, "started job", job)
	}
}

func NewWorker(id int) *worker {
	w := &worker{
		id:       id,
		workChan: make(chan QueryJob),
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

// Bugged code
/*
		jobs := make(chan Query, benchdb.wp.maxThreads)
		results := make(chan time.Duration, benchdb.wp.maxThreads)

		for thread := 0; thread < benchdb.wp.maxThreads; thread++ {
			go tWork(thread, benchdb, jobs, results)
		}

		for k := range records {
			query := Query{records[k][0], records[k][1], records[k][2]}
			jobs <- query
		}
		close(jobs)

		for thread := 0; thread < benchdb.wp.maxThreads; thread++ {
			<-results
		}

		func tWork(id int, b *benchApp, query <-chan Query, results chan<- time.Duration) {
	var q Query

	for {
		select {
		case q = <-query:
			log.Println("worker", id, "started query to", q.Hostname, "from", q.StartTime, "to", q.EndTime)

			tTime, err := q.queryDB(*b)
			if err != nil {
				log.Fatal(err)
			}

			results <- tTime
		}
	}
}
*/
