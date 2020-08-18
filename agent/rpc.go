package agent

import "github.com/gobench-io/gobench/metrics"

// RPC is the struct that expose RPC interface
// avoid the rpc.Register: method "xxx" has 1 input parameters; needs exactly three
type RPC struct {
	ml metricLoggerRPC
}

func newRPC(ml metricLoggerRPC) (*RPC, error) {
	return &RPC{ml: ml}, nil
}

// FindCreateGroup wrapper func
func (ar *RPC) FindCreateGroup(req *metrics.FCGroupReq, res *metrics.FCGroupRes) error {
	return ar.ml.FindCreateGroup(req, res)
}

// FindCreateGraph wrapper func
func (ar *RPC) FindCreateGraph(req *metrics.FCGraphReq, res *metrics.FCGraphRes) error {
	return ar.ml.FindCreateGraph(req, res)
}

// FindCreateMetric wrapper func
func (ar *RPC) FindCreateMetric(req *metrics.FCMetricReq, res *metrics.FCMetricRes) error {
	return ar.ml.FindCreateMetric(req, res)
}

// Counter wrapper func
func (ar *RPC) Counter(req *metrics.CounterReq, res *metrics.CounterRes) error {
	return ar.ml.Counter(req, res)
}

// Histogram wrapper func
func (ar *RPC) Histogram(req *metrics.HistogramReq, res *metrics.HistogramRes) error {
	return ar.ml.Histogram(req, res)
}

// Gauge wrapper func
func (ar *RPC) Gauge(req *metrics.GaugeReq, res *metrics.GaugeRes) error {
	return ar.ml.Gauge(req, res)
}
