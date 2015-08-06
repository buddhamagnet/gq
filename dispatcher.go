package gq

import "fmt"

// WorkerQueue is a buffered channel that holds the work channels
var (
	WorkQueue   chan WorkRequestInterface
	WorkerQueue chan chan WorkRequestInterface
	logger      func(...interface{})
)

func Logger(logFunc func(...interface{})) {
	logger = logFunc
}

// StartDispatcher starts the dispatcher
func StartDispatcher(nworkers int) {
	WorkQueue = make(chan WorkRequestInterface, nworkers)
	// Initialize channel where worker's work channel gets sent to.
	WorkerQueue = make(chan chan WorkRequestInterface, nworkers)

	// Create our workers
	for i := 0; i < nworkers; i++ {
		logger(fmt.Sprintf("Starting worker %d", i+1))
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				logger("Received work request")
				go func() {
					worker := <-WorkerQueue

					logger("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
