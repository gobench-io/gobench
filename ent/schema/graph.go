package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Graph holds the schema definition for the Graph entity.
type Graph struct {
	ent.Schema
}

// Fields of the Graph.
func (Graph) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").Immutable().StructTag(`json:"title"`),
		field.String("unit").StructTag(`json:"unit"`),
	}
}

// Edges of the Graph.
func (Graph) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("group", Group.Type).Ref("graphs").Unique(),
		edge.To("metrics", Metric.Type),
	}
}
