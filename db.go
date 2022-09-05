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
	"time"

	_ "github.com/lib/pq"
)

// Hardcoded datasource
// Insecure: refactor for production usage
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "postgres"
	dbname   = "homework"
)

const benchQueryTempl = `
	SELECT time_bucket('1 minutes', ts) AS one_min, 
	min(usage), 
	max(usage) 
	FROM cpu_usage 
	WHERE host = $1 AND ts >= $2 AND ts < $3 
	GROUP BY one_min 
	ORDER BY one_min ASC;
	`

// Query represents one instance of sql query
type Query struct {
	Hostname  string
	StartTime string
	EndTime   string
}

// benchDB represents one instance of benchDB
type benchDB struct {
	DB *sql.DB
}

// NewBenchDB returns a new instance of benchDB
func NewBenchDB() *benchDB {
	conn := handleDBConnection()
	db, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatalf("db not connected: %v", err)
	}
	log.Println("db connected!")

	return &benchDB{db}
}

// handleDBConnection returns a connection to a DB
// Insecure: refactor for production usage
func handleDBConnection() string {
	return fmt.Sprintf("user=%s password=%s dbname=%s port=%v sslmode=disable", user, password, dbname, port)
}

// queryDB creates a bench query
// It returns time.Duration as benchmarking data to be processed
func (p *Query) queryDB(b benchApp) error {
	// Increase number of queries processed
	b.incNumQueries()

	// Start benchmarking
	t0 := time.Now()

	// Query the database and discard data to avoid using memory
	data, err := b.db.DB.Query(benchQueryTempl, p.Hostname, p.StartTime, p.EndTime)
	if err != nil {
		return err
	}
	defer data.Close()

	// Stop benchmarking
	t1 := time.Now()
	tTotal := t1.Sub(t0)

	// Update statistics
	b.ProcessData(tTotal)

	// DEBUG
	// log.Printf("Querying %s from %v to %v took %v ms", p.Hostname, p.StartTime, p.EndTime, float64(tTotal)/float64(time.Millisecond))

	return nil
}
