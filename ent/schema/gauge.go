package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Gauge holds the schema definition for the Gauge entity.
type Gauge struct {
	ent.Schema
}

// Fields of the Gauge.
func (Gauge) Fields() []ent.Field {
	// counter: id, time, count
	// historgram: id, time, count, min, max, mean, stddev, median, 75, 95, 99, 99.9
	// gauge: id, time, value
	return []ent.Field{
		field.Int64("time").StructTag(`json:"time"`),
		field.Int64("value").StructTag(`json:"value"`),

		field.String("wID").StructTag(`json:"wId"`),
	}
}

// Edges of the Gauge.
func (Gauge) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metric", Metric.Type).Ref("gauges").Unique(),
	}
}
