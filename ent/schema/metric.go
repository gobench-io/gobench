package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Metric holds the schema definition for the Metric entity.
type Metric struct {
	ent.Schema
}

// Fields of the Metric.
func (Metric) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Immutable().StructTag(`json:"title"`),
		field.String("type").StructTag(`json:"type"`),
	}
}

// Edges of the Metric.
func (Metric) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("graph", Graph.Type).Ref("metrics").Unique(),
		edge.To("histograms", Histogram.Type),
		edge.To("counters", Counter.Type),
		edge.To("gauges", Gauge.Type),
	}
}
