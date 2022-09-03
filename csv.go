package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

// GetCSVData is the public function that returns all the csv records
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

		csvRecords, err := getCSVData(f)
		if err != nil {
			return nil, err
		}

		return csvRecords, nil

	} else {
		csvRecords, err := getCSVData(os.Stdin)
		if err != nil {
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

// getCSVData is the private function that returns the data
func getCSVData(f *os.File) ([][]string, error) {
	// Check if file has data
	if !fileHasData(f) {
		err := fmt.Errorf("getCSVData.fileHasData: File %v doesn't contain any data", f.Name())
		return nil, err
	}

	// Initialize reader and get the csv records
	csvReader, err := InitializeReader(f)
	if err != nil {
		err := fmt.Errorf("getCSVData.csvReader.ReadAll: error reading %v: %v", f.Name(), err)
		return nil, err
	}

	// ReadAll at once
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

// inputIsStdin checks if the input is os.Stdin or a file
func inputIsStdin(filename string) bool {
	return filename == "os.Stdin"
}
