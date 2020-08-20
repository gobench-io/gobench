package schema

import (
	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Counter holds the schema definition for the Counter entity.
type Counter struct {
	ent.Schema
}

// Fields of the Counter.
func (Counter) Fields() []ent.Field {
	// counter: id, time, count
	// historgram: id, time, count, min, max, mean, stddev, median, 75, 95, 99, 99.9
	// gauge: id, time, value
	return []ent.Field{
		field.Int64("time").StructTag(`json:"time"`),
		field.Int64("count").StructTag(`json:"count"`),

		field.String("wID").StructTag(`json:"wId"`),
	}
}

// Edges of the Counter.
func (Counter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("metric", Metric.Type).Ref("counters").Unique(),
	}
}
