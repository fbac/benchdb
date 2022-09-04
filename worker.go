package main

/*
type queryFunc func() error

type QueryJob struct {
	query queryFunc
}

type worker struct {
	queries chan QueryJob
}

func NewWorker() *worker {
	worker := &worker{
		queries: make(chan QueryJob),
	}

	//go worker.startWorker()

	return worker
}

//func (w *worker) startWorker() {
//	for q := range w.queries {
//q.query
//	}
//}

func (w *worker) Do(query QueryJob) {
	w.queries <- query
}

// Close is stopping worker.
func (w *worker) Close() {
	close(w.queries)
}

// Bugged code
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
