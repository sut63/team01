package schema

import "github.com/facebookincubator/ent"

// Medicine holds the schema definition for the Medicine entity.
type Medicine struct {
	ent.Schema
}

// Fields of the Medicine.
func (Medicine) Fields() []ent.Field {
	return nil
}

// Edges of the Medicine.
func (Medicine) Edges() []ent.Edge {
	return nil
}
