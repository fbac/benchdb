package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// GetCSVData is the public function that returns all the csv records
// Return all records at once: more expensive in mem, nicer on cpu
func GetCSVData(filename string) ([][]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	// Handle the file closure in a func, so the err can be handled
	defer func() {
		err := f.Close()
		if err != nil {
			log.Fatalf("csv: error closing %s: %v", f.Name(), err)
		}
	}()

	return getCSVData(f)
}

// initializeReader discards the first csv line (headers)
// And returns a csv.Reader pointing to the next offset
func initializeReader(f *os.File) (*csv.Reader, error) {
	r := csv.NewReader(f)

	_, err := r.Read()
	if err != nil {
		err := fmt.Errorf("DiscardFirstLine: error reading %v: %v", f.Name(), err)
		return nil, err
	}

	return r, nil
}

// getCSVData is the private function that returns the data
func getCSVData(f *os.File) ([][]string, error) {
	// Check if file has data
	if !fileHasData(f) {
		err := fmt.Errorf("getCSVData.fileHasData: File %v doesn't contain any data", f.Name())
		return nil, err
	}

	// Initialize reader and get the csv records
	csvReader, err := initializeReader(f)
	if err != nil {
		err := fmt.Errorf("getCSVData.csvReader.ReadAll: error reading %v: %v", f.Name(), err)
		return nil, err
	}

	// ReadAll at once
	// It's more expensive in memory, but nicer on cpu
	// Memory is less expensive than cpu nowadays
	csvRecords, err := csvReader.ReadAll()
	if err != nil {
		err := fmt.Errorf("getCSVData.csvReader.ReadAll: error reading %v: %v", f.Name(), err)
		return nil, err
	}

	return csvRecords, nil
}

// fileHasData checks if the input file contains data
func fileHasData(file *os.File) bool {
	f, err := file.Stat()

	if err != nil {
		return false
	}

	if f.Size() > 0 {
		return true
	}

	return false
}
