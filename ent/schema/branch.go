package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Branch holds the schema definition for the Branch entity.
type Branch struct {
	ent.Schema
}

// Fields of the Branch.
func (Branch) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Default("ShopName").Comment("nickname"),
		field.String("address").Unique(),
		field.Time("opening_date"),
	}
}

// Edges of the Branch.
func (Branch) Edges() []ent.Edge {
	return nil
}
