package main

import (
	"log"
	"os/exec"
	"strconv"
	"strings"

	"github.com/prometheus/client_golang/prometheus"
)

type lvmLvCollector struct {
	lvDataSizeMetric       *prometheus.Desc
	lvMetadataSizeMetric   *prometheus.Desc
	lvDataFilledMetric     *prometheus.Desc
	lvMetadataFilledMetric *prometheus.Desc
}

func newLvmLvCollector() *lvmLvCollector {
	return &lvmLvCollector{
		lvDataSizeMetric: prometheus.NewDesc("lvm_lv_size_bytes",
			"Shows LVM LV data size in Bytes",
			[]string{"lv_pool_name", "lv_name", "vg_name"}, nil,
		),
		lvMetadataSizeMetric: prometheus.NewDesc("lvm_lv_metadata_size_bytes",
			"Shows LVM LV metadata size in Bytes",
			[]string{"lv_pool_name", "lv_name", "vg_name"}, nil,
		),
		lvDataFilledMetric: prometheus.NewDesc("lvm_lv_filled_percentage",
			"Shows LVM LV data filled percentage",
			[]string{"lv_pool_name", "lv_name", "vg_name"}, nil,
		),
		lvMetadataFilledMetric: prometheus.NewDesc("lvm_lv_metadata_filled_percentage",
			"Shows LVM LV metadata filled percentage",
			[]string{"lv_pool_name", "lv_name", "vg_name"}, nil,
		),
	}
}

func (collector *lvmLvCollector) Describe(ch chan<- *prometheus.Desc) {
	ch <- collector.lvDataSizeMetric
	ch <- collector.lvMetadataSizeMetric
	ch <- collector.lvDataFilledMetric
	ch <- collector.lvMetadataFilledMetric
}

func (collector *lvmLvCollector) Collect(ch chan<- prometheus.Metric) {
	// /sbin/lvs --units B --separator , -o lv_size,lv_metadata_size,data_percent,metadata_percent,pool_lv,lv_name,vg_name --noheadings
	out, err := exec.Command("/sbin/lvs", "--units", "B", "--separator", ",", "-o", "lv_size,lv_metadata_size,data_percent,metadata_percent,pool_lv,lv_name,vg_name", "--noheadings").Output()
	if err != nil {
		log.Print(err)
	}

	lines := strings.Split(string(out), "\n")
	for _, line := range lines {
		values := strings.Split(line, ",")
		if len(values) < 6 {
			continue
		}

		descriptors := []*prometheus.Desc{collector.lvDataSizeMetric, collector.lvMetadataSizeMetric, collector.lvDataFilledMetric, collector.lvMetadataFilledMetric}
		poolName := values[4]
		logicalVolumeName := values[5]
		volumeGroupName := values[6]

		for index, descriptor := range descriptors {
			value, err := strconv.ParseFloat(strings.Trim(values[index], "B"), 64)
			if err != nil {
				continue
			}

			ch <- prometheus.MustNewConstMetric(descriptor, prometheus.GaugeValue, value, poolName, logicalVolumeName, volumeGroupName)
		}
	}
}
