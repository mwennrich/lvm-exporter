package main

import (
	"net/http"
	"os"
	"time"

	"log"
	"log/slog"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// LVM collector, listen to port 9080 path /metrics
func main() {
	node := os.Getenv("KUBE_NODE_NAME")
	if len(node) == 0 {
		var err error
		node, err = os.Hostname()
		if err != nil {
			node = "Unkown"
		}
	}
	lvmVgCollector := newLvmVgCollector(node)
	prometheus.MustRegister(lvmVgCollector)

	lvmLvCollector := newLvmLvCollector(node)
	prometheus.MustRegister(lvmLvCollector)

	http.Handle("/metrics", promhttp.Handler())
	slog.Info("Beginning to serve on port :9080")
	server := &http.Server{
		Addr:              ":9080",
		ReadHeaderTimeout: 5 * time.Second,
	}
	log.Fatal(server.ListenAndServe())
}
