package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type lvmLvCollector struct {
	lvTotalSizeMetric *prometheus.Desc
	node              string
}

func newLvmLvCollector(node string) *lvmLvCollector {
	return &lvmLvCollector{
		lvTotalSizeMetric: prometheus.NewDesc("lvm_lv_total_size_bytes",
			"Shows LVM LV total size in Bytes",
			[]string{"lv_name", "vg_name", "node"}, nil,
		),
		node: node,
	}
}

func (collector *lvmLvCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.lvTotalSizeMetric
}

func (collector *lvmLvCollector) Collect(ch chan<- prometheus.Metric) {
	out, err := exec.Command("/sbin/lvs", "--units", "B", "--separator", ",", "-o", "lv_size,lv_name,vg_name", "--noheadings").Output()
	if err != nil {
		log.Print(err)
	}
	lines := strings.SplitSeq(string(out), "\n")
	for line := range lines {
		values := strings.Split(strings.TrimSpace(line), ",")
		if len(values) < 3 {
			continue
		}

		logicalVolumeName := values[1]
		volumeGroupName := values[2]
		size, err := strconv.ParseFloat(strings.Trim(values[0], "B"), 64)
		if err != nil {
			continue
		}

		ch <- prometheus.MustNewConstMetric(collector.lvTotalSizeMetric, prometheus.GaugeValue, size, logicalVolumeName, volumeGroupName, collector.node)
	}
}
