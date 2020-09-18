package schema

import (
	"time"

	"github.com/facebook/ent"
	"github.com/facebook/ent/schema/edge"
	"github.com/facebook/ent/schema/field"
)

// Accesstoken holds the schema definition for the Accesstoken entity.
type Accesstoken struct {
	ent.Schema
}

// Fields of the Accesstoken.
func (Accesstoken) Fields() []ent.Field {
	return []ent.Field{
		field.String("accesstoken"),
		field.Uint64("ttl"),
		field.Time("created_at").
			Default(time.Now),
	}
}

// Edges of the Accesstoken.
func (Accesstoken) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).Ref("accesstokens"),
	}
}
