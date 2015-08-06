package gq

import (
	"testing"
	"time"
)

var count int

type WorkerTest struct {
	Delay time.Duration
}

func (w WorkerTest) Work() {
	count++
}

func (w WorkerTest) Data() string {
	return ""
}

func (w WorkerTest) DelayTime() time.Duration {
	return w.Delay
}

func (w WorkerTest) Preprocess() string {
	return ""
}

func (w WorkerTest) Postprocess() string {
	return ""
}

func nullLogger(...interface{}) {}

func init() {
	Logger(nullLogger)
	StartDispatcher(10)
	for i := 0; i < 10; i++ {
		work := WorkerTest{Delay: 0 * time.Second}
		WorkQueue <- work
	}
	time.Sleep(1 * time.Second)
}

func TestIncrement(t *testing.T) {
	if count != 10 {
		t.Errorf("expected count to be 10, got %d\n", count)
	}
}
