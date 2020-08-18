// Code generated by entc, DO NOT EDIT.

package histogram

const (
	// Label holds the string label denoting the histogram type in the database.
	Label = "histogram"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldMin holds the string denoting the min field in the database.
	FieldMin = "min"
	// FieldMax holds the string denoting the max field in the database.
	FieldMax = "max"
	// FieldMean holds the string denoting the mean field in the database.
	FieldMean = "mean"
	// FieldStddev holds the string denoting the stddev field in the database.
	FieldStddev = "stddev"
	// FieldMedian holds the string denoting the median field in the database.
	FieldMedian = "median"
	// FieldP75 holds the string denoting the p75 field in the database.
	FieldP75 = "p75"
	// FieldP95 holds the string denoting the p95 field in the database.
	FieldP95 = "p95"
	// FieldP99 holds the string denoting the p99 field in the database.
	FieldP99 = "p99"
	// FieldP999 holds the string denoting the p999 field in the database.
	FieldP999 = "p999"
	// FieldWID holds the string denoting the wid field in the database.
	FieldWID = "w_id"

	// EdgeMetric holds the string denoting the metric edge name in mutations.
	EdgeMetric = "metric"

	// Table holds the table name of the histogram in the database.
	Table = "histograms"
	// MetricTable is the table the holds the metric relation/edge.
	MetricTable = "histograms"
	// MetricInverseTable is the table name for the Metric entity.
	// It exists in this package in order to avoid circular dependency with the "metric" package.
	MetricInverseTable = "metrics"
	// MetricColumn is the table column denoting the metric relation/edge.
	MetricColumn = "metric_histograms"
)

// Columns holds all SQL columns for histogram fields.
var Columns = []string{
	FieldID,
	FieldTime,
	FieldCount,
	FieldMin,
	FieldMax,
	FieldMean,
	FieldStddev,
	FieldMedian,
	FieldP75,
	FieldP95,
	FieldP99,
	FieldP999,
	FieldWID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Histogram type.
var ForeignKeys = []string{
	"metric_histograms",
}
