package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// EventLog holds the schema definition for the EventLog entity.
type EventLog struct {
	ent.Schema
}

// Fields of the EventLog.
func (EventLog) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Immutable().Default("application").StructTag(`json:"name"`),
		field.String("message").Immutable().StructTag(`json:"message"`),
		field.String("level").Immutable().Default("info").StructTag(`json:"level"`),
		field.String("source").Immutable().StructTag(`json:"source"`),
		field.Time("created_at").Immutable().Default(time.Now).StructTag(`json:"created_at"`),
	}
}

// Edges of the EventLog.
func (EventLog) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("applications", Application.Type).
			Ref("eventLogs").
			Unique(),
	}
}
