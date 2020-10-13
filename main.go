package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	log "github.com/sirupsen/logrus"
)

// LVM collector, listen to port 9080 path /metrics
func main() {
	lvmVgCollector := newLvmVgCollector()
	prometheus.MustRegister(lvmVgCollector)

	lvmLvCollector := newLvmLvCollector()
	prometheus.MustRegister(lvmLvCollector)

	http.Handle("/metrics", promhttp.Handler())
	log.Info("Beginning to serve on port :9080")
	log.Fatal(http.ListenAndServe(":9080", nil))
}
