package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Application holds the schema definition for the Application entity.
type Application struct {
	ent.Schema
}

// Fields of the Application.
func (Application) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("status"),
		field.Time("created_at").
			Default(time.Now),
		field.Time("started_at").
			Optional(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Text("scenario"),
		field.Text("gomod").
			Default(""),
		field.Text("gosum").
			Default(""),
		field.Int("vu").
			Positive().
			Default(0),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type),
		edge.To("tags", Tag.Type),
	}
}
