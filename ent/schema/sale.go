package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Sale holds the schema definition for the Sale entity.
type Sale struct {
	ent.Schema
}

// Fields of the Sale.
func (Sale) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Sale.
func (Sale) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("branch", Branch.Type).Ref("sales").Unique().Required(),
	}
}
