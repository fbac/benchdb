package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	// Initialize cmd flags
	csvFile := flag.String("csv-file", "os.Stdin", "path to query csv file")
	//maxThreads := flag.Int("max-threads", 1, "max threads to process csv")
	flag.Parse()

	// Retrieve all the csv records
	records, err := GetCSVData(*csvFile)
	if err != nil {
		log.Fatalf("main: error retrieving data from %s: %v", *csvFile, err)
	}

	// Temporal debug
	for _, v := range records {
		fmt.Println(v)
	}
}
