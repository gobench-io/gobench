package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Log holds the schema definition for the Log entity.
type Log struct {
	ent.Schema
}

// Fields of the Log.
func (Log) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable().StructTag(`json:"name"`),
		field.String("message").Immutable().StructTag(`json:"message"`),
		field.String("level").Immutable().StructTag(`json:"level"`),
		field.String("source").Immutable().StructTag(`json:"source"`),
		field.String("created_at").Immutable().StructTag(`json:"created_at"`),
	}
}

// Edges of the Log.
func (Log) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("logs").
			Unique(),
	}
}
