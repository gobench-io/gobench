package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			NotEmpty().
			MaxLen(255).
			StructTag(`json:"name"`),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("application", Application.Type).
			Ref("tags").
			Unique(),
	}
}

// Indexes tag name
func (Tag) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("name").
			Edges("application").
			Unique(),
	}
}
