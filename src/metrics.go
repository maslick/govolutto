package src

import (
	"expvar"
	_ "expvar"
	"runtime"
	"time"
)

func NewMetrics(duration int) {
	var interval = time.Duration(duration) * time.Second
	var goroutines = expvar.NewInt("num_goroutine")
	for {
		<-time.After(interval)
		goroutines.Set(int64(runtime.NumGoroutine()))
	}
}
