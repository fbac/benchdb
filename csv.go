package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// GetCSVData returns all the csv records
// TODO: Handle file and os.Stdin in separated functions
func GetCSVData(filename string) ([][]string, error) {
	if !inputIsStdin(filename) {
		f, err := os.Open(filename)
		if err != nil {
			err := fmt.Errorf("GetCSVData.Open(%s)", f.Name())
			return nil, err
		}

		// Handle the file closure in a func, so the err can be handled
		defer func() {
			err := f.Close()
			if err != nil {
				log.Fatalf("csv: error closing %s: %v", f.Name(), err)
			}
		}()

		// Check if file has data
		if !fileHasData(f) {
			err := fmt.Errorf("csv.fileHasData: File %v doesn't contain any data", f.Name())
			return nil, err
		}

		// Initialize reader and get the csv records
		csvReader, err := InitializeReader(f)
		if err != nil {
			err := fmt.Errorf("csvReader.ReadAll: error reading %v: %v", f.Name(), err)
			return nil, err
		}

		// ReadAll at once
		csvRecords, err := csvReader.ReadAll()
		if err != nil {
			err := fmt.Errorf("csvReader.ReadAll: error reading %v: %v", f.Name(), err)
			return nil, err
		}

		return csvRecords, nil

	} else {
		// Check if file has data
		if !fileHasData(os.Stdin) {
			err := fmt.Errorf("csv.fileHasData: os.Stdin doesn't contain any data")
			return nil, err
		}

		// Initialize reader and get the csv records
		csvReader, err := InitializeReader(os.Stdin)
		if err != nil {
			err := fmt.Errorf("csvReader.ReadAll: error reading os.Stdin: %v", err)
			return nil, err
		}

		// ReadAll at once
		csvRecords, err := csvReader.ReadAll()
		if err != nil {
			err := fmt.Errorf("csvReader.ReadAll: error reading os.Stdin: %v", err)
			return nil, err
		}

		return csvRecords, nil
	}
}

// InitializeReader discards the first csv line (headers)
// And returns a csv.Reader pointing to the next offset
func InitializeReader(f *os.File) (*csv.Reader, error) {
	r := csv.NewReader(f)

	_, err := r.Read()
	if err != nil {
		err := fmt.Errorf("DiscardFirstLine: error reading %v: %v", f.Name(), err)
		return nil, err
	}

	return r, nil
}

// fileHasData checks the input file size is greater than 0B
func fileHasData(file *os.File) bool {
	f, err := file.Stat()
	fmt.Printf("fileHasData: %v\n", file.Name())
	if err != nil {
		return false
	}

	if f.Size() > 0 {
		return true
	}

	return false
}

func inputIsStdin(filename string) bool {
	return filename == "os.Stdin"
}
