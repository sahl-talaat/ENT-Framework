package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Department holds the schema definition for the Department entity.
type Department struct {
	ent.Schema
}

// Fields of the Department.
func (Department) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
		field.Time("created_at").Default(time.Now),
	}
}

// Edges of the Department.
func (Department) Edges() []ent.Edge {
	return nil
}
