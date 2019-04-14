package main

import (
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	numOfInc := promauto.NewCounter(prometheus.CounterOpts{
		Name: "number_of_incrementals",
		Help: "The number of incrementals",
	})

	go func() {
		for {
			numOfInc.Inc()
			time.Sleep(1 * time.Second)
		}
	}()

	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":3000", nil)
}
