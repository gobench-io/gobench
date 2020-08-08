package metrics

type MetricType string

const (
	// Counter is single additive value. New values are simply added to the current one.
	Counter MetricType = "counter"
	// Histogram is a set of numerical values that quantify a distribution of values. New values are added to the distribution.
	Histogram = "histogram"
	// Gauge is a single non-additive value. New value replaces the previous one.
	Gauge = "gauge"
)

type Metric struct {
	Title string
	Type  MetricType
}

type Graph struct {
	Title   string
	Unit    string
	Metrics []Metric
}

type Group struct {
	Name   string
	Graphs []Graph
}

// rpc

// FCGroupReq for find or create group rpc request
type FCGroupReq struct {
	Name  string
	AppID int
}

type FCGroupRes struct {
	ID int
}

// FCGraphReq for find or create graph rpc request
type FCGraphReq struct {
	Title   string
	Unit    string
	GroupID int
}
type FCGraphRes struct {
	ID int
}

// FCMetricReq for find or create metric rpc request
type FCMetricReq struct {
	Title   string
	Type    MetricType
	GraphID int
}
type FCMetricRes struct {
	ID int
}

type basedReqMetric struct {
	EID   string // executor ID
	AppID int    // app ID
	MID   int    // metric ID
	Time  int64
}
type basedResMetric struct {
	Success bool
	AppID   int
}

type HistogramValues struct {
	Count  int64
	Min    int64
	Max    int64
	Mean   float64
	Stddev float64
	Median float64
	P75    float64
	P95    float64
	P99    float64
	P999   float64
}

type CounterReq struct {
	basedReqMetric
	Count int64
}
type CounterRes struct {
	basedResMetric
}

type HistogramReq struct {
	basedReqMetric
	HistogramValues
}
type HistogramRes struct {
	basedResMetric
}

type GaugeReq struct {
	basedReqMetric
	Gauge int64
}
type GaugeRes struct {
	basedResMetric
}
