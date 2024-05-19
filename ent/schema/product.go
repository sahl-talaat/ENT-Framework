package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Product holds the schema definition for the Product entity.
type Product struct {
	ent.Schema
}

// Fields of the Product.
func (Product) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("description"),
		field.Time("nillable_created_at").Default(time.Now).Nillable(),
	}
}

// Edges of the Product.
func (Product) Edges() []ent.Edge {
	return nil
}
