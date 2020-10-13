package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type lvmVgCollector struct {
	vgFreeMetric *prometheus.Desc
	vgSizeMetric *prometheus.Desc
}

// LVM Collector contains VG size and VG free in MB
func newLvmVgCollector() *lvmVgCollector {
	return &lvmVgCollector{
		vgFreeMetric: prometheus.NewDesc("lvm_vg_free_bytes",
			"Shows LVM VG free size in Bytes",
			[]string{"vg_name"}, nil,
		),
		vgSizeMetric: prometheus.NewDesc("lvm_vg_total_size_bytes",
			"Shows LVM VG total size in Bytes",
			[]string{"vg_name"}, nil,
		),
	}
}

func (collector *lvmVgCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.vgFreeMetric
	ch <- collector.vgSizeMetric
}

// LVM Collect, call OS command and set values
func (collector *lvmVgCollector) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("/sbin/vgs", "--units", "B", "--separator", ",", "-o", "vg_name,vg_free,vg_size", "--noheadings").Output()
	if err != nil {
		log.Print(err)
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		values := strings.Split(line, ",")
		if len(values) == 3 {
			freeSize, err := strconv.ParseFloat(strings.Trim(values[1], "B"), 64)
			if err != nil {
				log.Print(err)
			} else {
				totalSize, err := strconv.ParseFloat(strings.Trim(values[2], "B"), 64)
				if err != nil {
					log.Print(err)
				} else {
					vgName := strings.Trim(values[0], " ")
					ch <- prometheus.MustNewConstMetric(collector.vgFreeMetric, prometheus.GaugeValue, freeSize, vgName)
					ch <- prometheus.MustNewConstMetric(collector.vgSizeMetric, prometheus.GaugeValue, totalSize, vgName)
				}
			}
		}
	}

}
