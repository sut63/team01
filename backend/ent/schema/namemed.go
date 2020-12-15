package schema

import "github.com/facebookincubator/ent"

// NameMed holds the schema definition for the NameMed entity.
type NameMed struct {
	ent.Schema
}

// Fields of the NameMed.
func (NameMed) Fields() []ent.Field {
	return nil
}

// Edges of the NameMed.
func (NameMed) Edges() []ent.Edge {
	return nil
}
