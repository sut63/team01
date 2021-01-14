package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// MedicineType holds the schema definition for the MedicineType entity.
type MedicineType struct {
	ent.Schema
}

// Fields of the MedicineType.
func (MedicineType) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the MedicineType.
func (MedicineType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Medicine", Medicine.Type),
	}
}
