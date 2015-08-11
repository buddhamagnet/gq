package gq

import (
	"fmt"
	"time"
)

// WorkRequestInterface represents a work
// request, the main function being Work, which
// does the actual work required.
type WorkRequestInterface interface {
	Work()
	DelayTime() time.Duration
	Preprocess() string
	Postprocess() string
}

// Worker represents a worker, which contains
// a reference to the three channels used by
// gq to manage the queue.
type Worker struct {
	ID          int
	Work        chan WorkRequestInterface
	WorkerQueue chan chan WorkRequestInterface
	QuitChan    chan bool
}

// NewWorker creates and returns a new Worker object.
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
				logger(fmt.Sprintf("worker%d: %s", w.ID, work.Preprocess()))
				time.Sleep(work.DelayTime())
				work.Work()
				logger(fmt.Sprintf("worker%d: %s", w.ID, work.Postprocess()))
			case <-w.QuitChan:
				logger(fmt.Sprintf("worker%d stopping\n", w.ID))
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
