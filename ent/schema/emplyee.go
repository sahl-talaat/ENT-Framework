package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Emplyee holds the schema definition for the Emplyee entity.
type Emplyee struct {
	ent.Schema
}

/*
package ent

// User entity.
type User struct {
    RequiredName string `json:"required_name,omitempty"`
    OptionalName string `json:"optional_name,omitempty"`
    NillableName *string `json:"nillable_name,omitempty"`
}
*/

// Fields of the Emplyee.
func (Emplyee) Fields() []ent.Field {
	return []ent.Field{
		field.Int16("age"),
		field.String("name"),
		field.Float32("salary").Optional(),
		field.String("username").Unique(),
		field.Time("created_at").Default(time.Now).Immutable(), // Immutable cant update
	}
}

// Edges of the Emplyee.
func (Emplyee) Edges() []ent.Edge {
	return nil
}
