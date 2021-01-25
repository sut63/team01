package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// PositionInPharmacist holds the schema definition for the PositionInPharmacist entity.
type PositionInPharmacist struct {
	ent.Schema
}

// Fields of the PositionInPharmacist.
func (PositionInPharmacist) Fields() []ent.Field {
	return []ent.Field{
		field.String("position").NotEmpty().Unique(),
	}
}

// Edges of the PositionInPharmacist.
func (PositionInPharmacist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("pharmacist", Pharmacist.Type).StorageKey(edge.Column("positioninpharmacist_id")),
	}
}
