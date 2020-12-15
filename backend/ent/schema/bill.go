package schema

import "github.com/facebookincubator/ent"

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return nil
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return nil
}
