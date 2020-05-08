package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Metric holds the schema definition for the Metric entity.
type Metric struct {
	ent.Schema
}

// Fields of the Metric.
func (Metric) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Immutable().Unique().StructTag(`json:"title"`),
		field.String("type").StructTag(`json:"type"`),
	}
}

// Edges of the Metric.
func (Metric) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("graph", Graph.Type).Ref("metrics").Unique().StructTag("json:\"-\""),

		edge.To("histograms", Histogram.Type),
		edge.To("counters", Counter.Type),
		edge.To("gauges", Gauge.Type),
	}
}
