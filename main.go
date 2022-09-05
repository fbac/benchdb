package main

import (
	"flag"
	"log"
)

func main() {
	// Initialize cmd flags
	csvFile := flag.String("csv-file", "data-file.csv", "path to query csv file")
	maxThreads := flag.Int("max-threads", 1, "max threads to process csv")
	flag.Parse()

	// Initialize BenchApp
	benchdb := NewBenchApp(*maxThreads)

	// Retrieve all the csv records
	records, err := GetCSVData(*csvFile)
	if err != nil {
		log.Fatalf("main: error retrieving data from %s: %v", *csvFile, err)
	}

	// Create Job channel
	jobs := make(chan *Job, len(records))

	// Start workers
	for i := 0; i < *maxThreads; i++ {
		wg.Add(1)
		go doWork(i, jobs)
	}
	// Create work
	go createWork(records, benchdb, jobs)

	// Wait until all work is done
	wg.Wait()

	// Report benchmarking data
	benchdb.ReportData()
}
