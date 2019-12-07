//+build !test

package src

import (
	"github.com/prometheus/client_golang/prometheus"
	"runtime"
	"time"
)

func init() {
	goroutinesMetric := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "goroutines_number",
		Help: "Number of goroutines",
	})
	go goroutineMetricHandler(goroutinesMetric)
	prometheus.MustRegister(goroutinesMetric)
}

func goroutineMetricHandler(gauge prometheus.Gauge) {
	var interval = time.Duration(5) * time.Second
	for {
		<-time.After(interval)
		gauge.Set(float64(runtime.NumGoroutine()))
	}
}
