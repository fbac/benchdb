package main

import (
	"log"
	"sync"
)

var wg sync.WaitGroup

type Job struct {
	Hostname string
	Func     func()
}

func createWork(records [][]string, b *benchApp, jobs chan<- *Job) {
	for k := range records {
		query := Query{records[k][0], records[k][1], records[k][2]}
		queryFunc := func() {
			if err := query.queryDB(*b); err != nil {
				log.Fatal(err)
			}
		}
		jobs <- &Job{records[k][0], queryFunc}
	}

	close(jobs)
}

func doWork(id int, jobs <-chan *Job) {
	defer wg.Done()
	for job := range jobs {
		log.Printf("Running %v %s\n", id, job.Hostname)
		job.Func()
	}
}
