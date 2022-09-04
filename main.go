package main

import (
	"flag"
	"log"
)

func main() {
	// Initialize cmd flags
	csvFile := flag.String("csv-file", "foo/bar/data-file.csv", "path to query csv file")
	maxThreads := flag.Int("max-threads", 1, "max threads to process csv")
	flag.Parse()

	// Initialize BenchApp
	benchdb := NewBenchApp(*maxThreads)

	// Retrieve all the csv records
	records, err := GetCSVData(*csvFile)
	if err != nil {
		log.Fatalf("main: error retrieving data from %s: %v", *csvFile, err)
	}

	// DEBUG: Temporal boilerplate code
	for k := range records {
		query := Query{records[k][0], records[k][1], records[k][2]}
		tTime, err := query.queryDB(*benchdb)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("query %v took %v\n", k, tTime)
	}

	benchdb.reportData()
}
