package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Branch holds the schema definition for the Branch entity.
type Branch struct {
	ent.Schema
}

// Fields of the Branch.
func (Branch) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("address").Unique(),
		field.Time("opening_date").Default(time.Now),
	}
}

// Edges of the Branch.
func (Branch) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("employees", Employee.Type),
	}
}
