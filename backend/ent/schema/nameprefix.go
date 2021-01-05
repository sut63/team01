package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// NamePrefix holds the schema definition for the NamePrefix entity.
type NamePrefix struct {
	ent.Schema
}

// Fields of the NamePrefix.
func (NamePrefix) Fields() []ent.Field {
	return []ent.Field{
		field.String("prefix").NotEmpty(),
	}
}

// Edges of the NamePrefix.
func (NamePrefix) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pharmacist", Pharmacist.Type).StorageKey(edge.Column("nameprefix_id")),
	}
}
