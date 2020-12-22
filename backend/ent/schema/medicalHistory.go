package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )
 

// MedicalHistory holds the schema definition for the MedicalHistory entity.
type MedicalHistory struct {
	ent.Schema
}

// Fields of the MedicalHistory.
func (MedicalHistory) Fields() []ent.Field {
	return []ent.Field{
		field.Time("date"),
		field.Time("time"),
		field.Float("weight").Positive(),
		field.Float("hight").Positive(),
		field.String("bloodPressure"),
		field.String("symptom"),
	}
}

// Edges of the MedicalHistory.
func (MedicalHistory) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", PatientInfo.Type).Ref("medicalhistorys").Unique(),
	}
}
