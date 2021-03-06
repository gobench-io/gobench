// Code generated by entc, DO NOT EDIT.

package counter

const (
	// Label holds the string label denoting the counter type in the database.
	Label = "counter"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTime holds the string denoting the time field in the database.
	FieldTime = "time"
	// FieldCount holds the string denoting the count field in the database.
	FieldCount = "count"
	// FieldWID holds the string denoting the wid field in the database.
	FieldWID = "w_id"

	// EdgeMetric holds the string denoting the metric edge name in mutations.
	EdgeMetric = "metric"

	// Table holds the table name of the counter in the database.
	Table = "counters"
	// MetricTable is the table the holds the metric relation/edge.
	MetricTable = "counters"
	// MetricInverseTable is the table name for the Metric entity.
	// It exists in this package in order to avoid circular dependency with the "metric" package.
	MetricInverseTable = "metrics"
	// MetricColumn is the table column denoting the metric relation/edge.
	MetricColumn = "metric_counters"
)

// Columns holds all SQL columns for counter fields.
var Columns = []string{
	FieldID,
	FieldTime,
	FieldCount,
	FieldWID,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Counter type.
var ForeignKeys = []string{
	"metric_counters",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}
