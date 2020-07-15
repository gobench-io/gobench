package server

import (
	gometrics "github.com/rcrowley/go-metrics"
)

type metriclog struct{}

func (m *metriclog) Counter(id, title string, time, c int64) error {
	return nil
}

func (m *metriclog) Histogram(id, title string, time int64, h gometrics.Histogram) error {
	return nil
}

func (m *metriclog) Gauge(id, title string, time int64, g int64) error {
	return nil
}
