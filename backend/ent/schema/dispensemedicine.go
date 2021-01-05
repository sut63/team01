package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// DispenseMedicine holds the schema definition for the DispenseMedicine entity.
type DispenseMedicine struct {
	ent.Schema
}

// Fields of the DispenseMedicine.
func (DispenseMedicine) Fields() []ent.Field {
	return []ent.Field{
		field.Time("datetime"),
	}
}

// Edges of the DispenseMedicine.
func (DispenseMedicine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pharmacist", Pharmacist.Type).Ref("dispensemedicine").Unique(),
		edge.From("annotation", Annotation.Type).Ref("dispensemedicine").Unique(),
	}
}
