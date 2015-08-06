package gq

import (
	"log"
	"time"
)

// WorkRequest interface
type WorkRequestInterface interface {
	Work()
	DelayTime() time.Duration
	Data() string
	Preprocess() string
	Postprocess() string
}

// Worker object
type Worker struct {
	ID          int
	Work        chan WorkRequestInterface
	WorkerQueue chan chan WorkRequestInterface
	QuitChan    chan bool
}

// NewWorker creates and returns a new Worker object
func NewWorker(id int, workerQueue chan chan WorkRequestInterface) Worker {
	worker := Worker{
		ID:          id,
		Work:        make(chan WorkRequestInterface),
		WorkerQueue: workerQueue,
		QuitChan:    make(chan bool),
	}

	return worker
}

// Start starts the worker
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Work

			select {
			case work := <-w.Work:
				// Receive a work request.
				log.Printf("worker%d: %s", w.ID, work.Preprocess())
				time.Sleep(work.DelayTime())

				log.Printf("worker%d: %s", w.ID, work.Postprocess())

				work.Work()

			case <-w.QuitChan:
				log.Printf("worker%d stopping\n", w.ID)
				return
			}
		}
	}()
}

// Stop stops the worker
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}
