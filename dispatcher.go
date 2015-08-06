package gq

import "log"

// WorkerQueue is a buffered channel that holds the work channels
var (
	WorkQueue   chan WorkRequestInterface
	WorkerQueue chan chan WorkRequestInterface
)

// StartDispatcher starts the dispatcher
func StartDispatcher(nworkers int) {
	WorkQueue = make(chan WorkRequestInterface, nworkers)
	// Initialize channel where worker's work channel gets sent to.
	WorkerQueue = make(chan chan WorkRequestInterface, nworkers)

	// Create our workers
	for i := 0; i < nworkers; i++ {
		log.Printf("Starting worker %d", i+1)
		worker := NewWorker(i+1, WorkerQueue)
		worker.Start()
	}

	go func() {
		for {
			select {
			case work := <-WorkQueue:
				log.Println("Received work request")
				go func() {
					worker := <-WorkerQueue

					log.Println("Dispatching work request")
					worker <- work
				}()
			}
		}
	}()
}
