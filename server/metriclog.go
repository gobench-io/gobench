package server

import (
	gometrics "github.com/rcrowley/go-metrics"
)

func (m *master) Counter(id, title string, time, c int64) error {
	return nil
}

func (m *master) Histogram(id, title string, time int64, h gometrics.Histogram) error {
	return nil
}

func (m *master) Gauge(id, title string, time int64, g int64) error {
	return nil
}
