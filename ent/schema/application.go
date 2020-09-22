package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
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
		field.String("tags").
			Default(""),
	}
}

// Edges of the Application.
func (Application) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("groups", Group.Type),
	}
}
