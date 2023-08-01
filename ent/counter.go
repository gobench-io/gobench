// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebook/ent/dialect/sql"
	"github.com/gobench-io/gobench/v2/ent/counter"
	"github.com/gobench-io/gobench/v2/ent/metric"
)

// Counter is the model entity for the Counter schema.
type Counter struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Time holds the value of the "time" field.
	Time int64 `json:"time"`
	// Count holds the value of the "count" field.
	Count int64 `json:"count"`
	// WID holds the value of the "wID" field.
	WID string `json:"wId"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CounterQuery when eager-loading is set.
	Edges           CounterEdges `json:"edges"`
	metric_counters *int
}

// CounterEdges holds the relations/edges for other nodes in the graph.
type CounterEdges struct {
	// Metric holds the value of the metric edge.
	Metric *Metric
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MetricOrErr returns the Metric value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e CounterEdges) MetricOrErr() (*Metric, error) {
	if e.loadedTypes[0] {
		if e.Metric == nil {
			// The edge metric was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: metric.Label}
		}
		return e.Metric, nil
	}
	return nil, &NotLoadedError{edge: "metric"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Counter) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // time
		&sql.NullInt64{},  // count
		&sql.NullString{}, // wID
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Counter) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // metric_counters
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Counter fields.
func (c *Counter) assignValues(values ...interface{}) error {
	if m, n := len(values), len(counter.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	c.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field time", values[0])
	} else if value.Valid {
		c.Time = value.Int64
	}
	if value, ok := values[1].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field count", values[1])
	} else if value.Valid {
		c.Count = value.Int64
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field wID", values[2])
	} else if value.Valid {
		c.WID = value.String
	}
	values = values[3:]
	if len(values) == len(counter.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field metric_counters", value)
		} else if value.Valid {
			c.metric_counters = new(int)
			*c.metric_counters = int(value.Int64)
		}
	}
	return nil
}

// QueryMetric queries the metric edge of the Counter.
func (c *Counter) QueryMetric() *MetricQuery {
	return (&CounterClient{config: c.config}).QueryMetric(c)
}

// Update returns a builder for updating this Counter.
// Note that, you need to call Counter.Unwrap() before calling this method, if this Counter
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Counter) Update() *CounterUpdateOne {
	return (&CounterClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (c *Counter) Unwrap() *Counter {
	tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Counter is not a transactional entity")
	}
	c.config.driver = tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Counter) String() string {
	var builder strings.Builder
	builder.WriteString("Counter(")
	builder.WriteString(fmt.Sprintf("id=%v", c.ID))
	builder.WriteString(", time=")
	builder.WriteString(fmt.Sprintf("%v", c.Time))
	builder.WriteString(", count=")
	builder.WriteString(fmt.Sprintf("%v", c.Count))
	builder.WriteString(", wID=")
	builder.WriteString(c.WID)
	builder.WriteByte(')')
	return builder.String()
}

// Counters is a parsable slice of Counter.
type Counters []*Counter

func (c Counters) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
