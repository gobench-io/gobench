// Code generated by ent, DO NOT EDIT.

package metric

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the metric type in the database.
	Label = "metric"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeGraph holds the string denoting the graph edge name in mutations.
	EdgeGraph = "graph"
	// EdgeHistograms holds the string denoting the histograms edge name in mutations.
	EdgeHistograms = "histograms"
	// EdgeCounters holds the string denoting the counters edge name in mutations.
	EdgeCounters = "counters"
	// EdgeGauges holds the string denoting the gauges edge name in mutations.
	EdgeGauges = "gauges"
	// Table holds the table name of the metric in the database.
	Table = "metrics"
	// GraphTable is the table that holds the graph relation/edge.
	GraphTable = "metrics"
	// GraphInverseTable is the table name for the Graph entity.
	// It exists in this package in order to avoid circular dependency with the "graph" package.
	GraphInverseTable = "graphs"
	// GraphColumn is the table column denoting the graph relation/edge.
	GraphColumn = "graph_metrics"
	// HistogramsTable is the table that holds the histograms relation/edge.
	HistogramsTable = "histograms"
	// HistogramsInverseTable is the table name for the Histogram entity.
	// It exists in this package in order to avoid circular dependency with the "histogram" package.
	HistogramsInverseTable = "histograms"
	// HistogramsColumn is the table column denoting the histograms relation/edge.
	HistogramsColumn = "metric_histograms"
	// CountersTable is the table that holds the counters relation/edge.
	CountersTable = "counters"
	// CountersInverseTable is the table name for the Counter entity.
	// It exists in this package in order to avoid circular dependency with the "counter" package.
	CountersInverseTable = "counters"
	// CountersColumn is the table column denoting the counters relation/edge.
	CountersColumn = "metric_counters"
	// GaugesTable is the table that holds the gauges relation/edge.
	GaugesTable = "gauges"
	// GaugesInverseTable is the table name for the Gauge entity.
	// It exists in this package in order to avoid circular dependency with the "gauge" package.
	GaugesInverseTable = "gauges"
	// GaugesColumn is the table column denoting the gauges relation/edge.
	GaugesColumn = "metric_gauges"
)

// Columns holds all SQL columns for metric fields.
var Columns = []string{
	FieldID,
	FieldTitle,
	FieldType,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "metrics"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"graph_metrics",
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

// OrderOption defines the ordering options for the Metric queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByTitle orders the results by the title field.
func ByTitle(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTitle, opts...).ToFunc()
}

// ByType orders the results by the type field.
func ByType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldType, opts...).ToFunc()
}

// ByGraphField orders the results by graph field.
func ByGraphField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGraphStep(), sql.OrderByField(field, opts...))
	}
}

// ByHistogramsCount orders the results by histograms count.
func ByHistogramsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newHistogramsStep(), opts...)
	}
}

// ByHistograms orders the results by histograms terms.
func ByHistograms(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newHistogramsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCountersCount orders the results by counters count.
func ByCountersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCountersStep(), opts...)
	}
}

// ByCounters orders the results by counters terms.
func ByCounters(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCountersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByGaugesCount orders the results by gauges count.
func ByGaugesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newGaugesStep(), opts...)
	}
}

// ByGauges orders the results by gauges terms.
func ByGauges(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newGaugesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newGraphStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GraphInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, GraphTable, GraphColumn),
	)
}
func newHistogramsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(HistogramsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, HistogramsTable, HistogramsColumn),
	)
}
func newCountersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CountersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, CountersTable, CountersColumn),
	)
}
func newGaugesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(GaugesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, GaugesTable, GaugesColumn),
	)
}
