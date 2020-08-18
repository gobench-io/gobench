package agent

import "github.com/gobench-io/gobench/metrics"

type metricLoggerRPC interface {
	FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) error
	FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) error
	FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) error
	Counter(req *metrics.CounterReq, res *metrics.CounterRes) error
	Histogram(req *metrics.HistogramReq, res *metrics.HistogramRes) error
	Gauge(req *metrics.GaugeReq, res *metrics.GaugeRes) error
}

// nop metric logger
type nopLog struct{}

func (n *nopLog) FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) error {
	return nil
}
func (n *nopLog) FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) error {
	return nil
}
func (n *nopLog) FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) error {
	return nil
}
func (n *nopLog) Counter(req *metrics.CounterReq, res *metrics.CounterRes) error {
	return nil
}
func (n *nopLog) Histogram(req *metrics.HistogramReq, res *metrics.HistogramRes) error {
	return nil
}
func (n *nopLog) Gauge(req *metrics.GaugeReq, res *metrics.GaugeRes) error {
	return nil
}

func newNopMetricLog() metricLoggerRPC {
	return &nopLog{}
}
