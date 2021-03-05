package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	klog "k8s.io/klog/v2"
)

// LVM collector, listen to port 9080 path /metrics
func main() {
	lvmVgCollector := newLvmVgCollector()
	prometheus.MustRegister(lvmVgCollector)

	lvmLvCollector := newLvmLvCollector()
	prometheus.MustRegister(lvmLvCollector)

	http.Handle("/metrics", promhttp.Handler())
	klog.Info("Beginning to serve on port :9080")
	klog.Fatal(http.ListenAndServe(":9080", nil))
}
