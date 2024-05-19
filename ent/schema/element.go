package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Element holds the schema definition for the Element entity.
type Element struct {
	ent.Schema
}

// Fields of the Element.
func (Element) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Time("nillable_created_at").Default(time.Now).Nillable(),
	}
}

// Edges of the Element.
func (Element) Edges() []ent.Edge {
	return nil
}
