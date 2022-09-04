package main

import (
	"log"
	"sort"
	"sync"
	"time"
)

// BenchApp represents one instance of benchdb
type benchApp struct {
	db   *benchDB
	wp   *workerPool
	data *benchAppData
}

// Constants to help with calculations
const (
	MaxUint = ^uint(0)
	MinUint = 0
	MaxInt  = int(MaxUint >> 1)
	MinInt  = -MaxInt - 1
)

// benchAppData represent the data to report
type benchAppData struct {
	mu             sync.Mutex
	NumQueries     int
	TotalTime      int
	MinTime        int
	MaxTime        int
	AvgTime        int
	TimeCollection []int
}

// Initialize is the public method to initialize the app
func NewBenchApp(maxThreads int) *benchApp {
	return &benchApp{
		db:   NewBenchDB(),
		wp:   NewWorkerPool(maxThreads),
		data: newBenchAppData(),
	}
}

// Initialize app data store
// Avoid copylocks: By passing data as pointer we avoid copying mutexes
// When passing benchApp through different funcions
func newBenchAppData() *benchAppData {
	return &benchAppData{
		mu:             sync.Mutex{},
		NumQueries:     0,
		TotalTime:      0,
		MinTime:        MaxInt,
		MaxTime:        MinInt,
		TimeCollection: make([]int, 0),
	}
}

// ProcessData is a public method to update data based on time.Duration
func (b *benchApp) ProcessData(t time.Duration) {
	b.incTotalTime(t)
	b.addDuration(t)
	b.updateMin(t)
	b.updateMax(t)
}

// increase NumQueries atomically
func (b *benchApp) incNumQueries() {
	b.data.mu.Lock()
	b.data.NumQueries++
	b.data.mu.Unlock()
}

// increase TotalTime atomically
func (b *benchApp) incTotalTime(t time.Duration) {
	b.data.mu.Lock()
	b.data.TotalTime += int(t)
	b.data.mu.Unlock()
}

// add a time.Duration to collection
func (b *benchApp) addDuration(t time.Duration) {
	b.data.mu.Lock()
	b.data.TimeCollection = append(b.data.TimeCollection, int(t))
	b.data.mu.Unlock()
}

// calculateMedian returns the 0.5 percentile
func (b *benchApp) calculateMedian() int {
	b.data.mu.Lock()
	sort.Ints(b.data.TimeCollection)
	median := b.data.TimeCollection[len(b.data.TimeCollection)/2]
	b.data.mu.Unlock()
	return median
}

// calculateAvg returns the average query processing time
func (b *benchApp) calculateAvg() int {
	var t int
	for _, v := range b.data.TimeCollection {
		t += v
	}
	return t / len(b.data.TimeCollection)
}

// updateMin updates the min query processing time
func (b *benchApp) updateMin(t time.Duration) {
	b.data.mu.Lock()
	if int(t) < b.data.MinTime {
		b.data.MinTime = int(t)
	}
	b.data.mu.Unlock()
}

// updateMax updates the max query processing time
func (b *benchApp) updateMax(t time.Duration) {
	b.data.mu.Lock()
	if int(t) > b.data.MaxTime {
		b.data.MaxTime = int(t)
	}
	b.data.mu.Unlock()
}

// Report BenchApp Data
func (b *benchApp) ReportData() {
	log.Printf("Total queries processed:\t%v\n", b.data.NumQueries)
	log.Printf("Min processing time:\t%v\tms\n", float64(b.data.MinTime)/float64(time.Millisecond))
	log.Printf("Max processing time:\t%v\tms\n", float64(b.data.MaxTime)/float64(time.Millisecond))
	log.Printf("Avg processing time:\t%v\tms\n", float64(b.calculateAvg())/float64(time.Millisecond))
	log.Printf("Med processing time:\t%v\tms\n", float64(b.calculateMedian())/float64(time.Millisecond))
	log.Printf("Total processing time:\t%v\tms\n", float64(b.data.TotalTime)/float64(time.Millisecond))
}
