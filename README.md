gq is a small package used for setting up a worker queue system in Go.

[![GoDoc](https://godoc.org/github.com/buddhamagnet/gq?status.svg)](https://godoc.org/github.com/buddhamagnet/gq)

###EXAMPLE

```go
package main

import (
	"flag"
	"fmt"
	"log"
	"time"

	"github.com/buddhamagnet/gq"
)

type WorkRequest struct {
	Delay time.Duration
}

func (w WorkRequest) Work() {
	fmt.Println("testing testing")
}

func (w WorkRequest) Data() string {
	return ""
}

func (w WorkRequest) DelayTime() time.Duration {
	return w.Delay
}

func (w WorkRequest) Preprocess() string {
	return fmt.Sprintf("Received work request, delaying for %f seconds", w.DelayTime().Seconds())
}

func (w WorkRequest) Postprocess() string {
	return fmt.Sprintf("Finished processing job")
}

var nworkers int

func init() {
	flag.IntVar(&nworkers, "n", 10, "Number of workers")
}

func main() {
	flag.Parse()
	gq.Logger(log.Println)
	gq.StartDispatcher(nworkers)
	work := WorkRequest{Delay: 5 * time.Second}
	gq.WorkQueue <- work
	time.Sleep(15 * time.Second)
}
```
