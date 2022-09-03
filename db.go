package main

/*
The following query will be used to benchmark the database:

SELECT time_bucket('1 minutes', ts) AS one_min,
  min(usage),
  max(usage)
  FROM cpu_usage
  WHERE host = 'host_000008' AND ts >= '2017-01-01 08:59:22' AND ts < '2017-01-01 09:59:22'
  GROUP BY one_min
  ORDER BY one_min ASC;

Expected:

        one_min         |  min  |  max
------------------------+-------+-------
 2017-01-01 08:59:00+00 | 27.54 | 51.01

 [...]

*/

import (
	"database/sql"
	"fmt"
	"log"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

// Hardcoded constant
// Insecure: refactor for production usage
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "homework"
)

// BenchApp represents one instance of benchdb
type BenchApp struct {
	DB *sql.DB
	benchAppData
}

type benchAppData struct {
	mu         sync.Mutex
	NumQueries int
	TotalTime  int
}

// Query represents one instance of sql query
type Query struct {
	Hostname  string
	StartTime string
	EndTime   string
}

// Initialize is the public method to initialize the app
func (b *BenchApp) Initialize() {
	var err error
	conn := fmt.Sprintf("user=%s password=%s dbname=%s port=%v sslmode=disable", user, password, dbname, port)

	b.DB, err = sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("db not connected: %v", err)
	}

	log.Println("db connected!")
}

func (b *BenchApp) incNumQueries() {
	b.benchAppData.mu.Lock()
	b.NumQueries++
	b.benchAppData.mu.Unlock()
}

func (b *BenchApp) incTotalTime(t time.Duration) {
	b.benchAppData.mu.Lock()
	b.TotalTime += int(t)
	b.benchAppData.mu.Unlock()
}

func (b *BenchApp) reportData() {
	log.Printf("Total queries processed: %v\n", b.NumQueries)
	log.Printf("Total processing time: %v ms\n", float64(b.benchAppData.TotalTime)/float64(time.Millisecond))
}

// queryDB is the method to create a bench query
// It returns time.Duration as benchmarking data to be processed
func (p *Query) queryDB(b *BenchApp) (time.Duration, error) {
	// Increase number of queries processed
	b.incNumQueries()

	// Start benchmarking
	t0 := time.Now()

	// Query DB
	_, err := b.DB.Query(`SELECT time_bucket('1 minutes', ts) AS one_min, min(usage), max(usage) FROM cpu_usage WHERE host = $1 AND ts >= $2 AND ts < $3 GROUP BY one_min ORDER BY one_min ASC;`, p.Hostname, p.StartTime, p.EndTime)
	if err != nil {
		return time.Duration(-1), err // return -1 if hit error
	}

	// Stop benchmarking
	t1 := time.Now()
	tTotal := t1.Sub(t0)

	// Increase total duration
	b.incTotalTime(tTotal)

	return tTotal, nil
}
