package main

import (
	"log"
	"sync"
	"time"
)

// BenchApp represents one instance of benchdb
type benchApp struct {
	db   *benchDB
	wp   *workerPool
	data *benchAppData
}

// benchAppData represent the data to report
type benchAppData struct {
	mu         sync.Mutex
	NumQueries int
	TotalTime  int
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
		mu:         sync.Mutex{},
		NumQueries: 0,
		TotalTime:  0,
	}
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

// report BenchApp Data
func (b *benchApp) reportData() {
	log.Printf("Total queries processed: %v\n", b.data.NumQueries)
	log.Printf("Total processing time: %v ms\n", float64(b.data.TotalTime)/float64(time.Millisecond))
}
