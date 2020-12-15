package schema

import "github.com/facebookincubator/ent"

// DispenseMedicine holds the schema definition for the DispenseMedicine entity.
type DispenseMedicine struct {
	ent.Schema
}

// Fields of the DispenseMedicine.
func (DispenseMedicine) Fields() []ent.Field {
	return nil
}

// Edges of the DispenseMedicine.
func (DispenseMedicine) Edges() []ent.Edge {
	return nil
}
