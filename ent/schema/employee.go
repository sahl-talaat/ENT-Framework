package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Emplyee holds the schema definition for the Emplyee entity.
type Employee struct {
	ent.Schema
}

// Fields of the Emplyee.
func (Employee) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Int("age"),
		field.Float("salary"),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Emplyee.
func (Employee) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("department", Department.Type).Ref("employees").Unique().Required(),
		edge.From("branch", Branch.Type).Ref("employees").Unique().Required(),
	}
}
