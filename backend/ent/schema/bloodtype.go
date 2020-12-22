package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// BloodType holds the schema definition for the BloodType entity.
type BloodType struct {
	ent.Schema
}

// Fields of the BloodType.
func (BloodType) Fields() []ent.Field {
	return []ent.Field{
		field.String("typeName").NotEmpty().Unique(),
	}
}

// Edges of the BloodType.
func (BloodType) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("patientinfo", PatientInfo.Type).StorageKey(edge.Column("bloodtype_id")).Unique(),
	}
}
