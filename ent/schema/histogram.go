package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Histogram holds the schema definition for the Histogram entity.
type Histogram struct {
	ent.Schema
}

// Fields of the Histogram.
func (Histogram) Fields() []ent.Field {
	// counter: id, time, count
	// historgram: id, time, count, min, max, mean, stddev, median, 75, 95, 99, 99.9
	// gauge: id, time, value
	return []ent.Field{
		field.Int64("time").StructTag(`json:"time"`),
		field.Int64("count").StructTag(`json:"count"`),
		field.Int64("min").StructTag(`json:"min"`),
		field.Int64("max").StructTag(`json:"max"`),
		field.Float("mean").StructTag(`json:"mean"`),
		field.Float("stddev").StructTag(`json:"stddev"`),
		field.Float("median").StructTag(`json:"median"`),
		field.Float("p75").StructTag(`json:"p75"`),
		field.Float("p95").StructTag(`json:"p95"`),
		field.Float("p99").StructTag(`json:"p99"`),
		field.Float("p999").StructTag(`json:"p999"`),
	}
}

// Edges of the Histogram.
func (Histogram) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metric", Metric.Type).Ref("histograms").Unique(),
	}
}
