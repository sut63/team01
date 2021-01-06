package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Prescription holds the schema definition for the Prescription entity.
type Prescription struct {
	ent.Schema
}

// Fields of the Prescription.
func (Prescription) Fields() []ent.Field {
	return []ent.Field{
		field.Int("value"),
	}
}

// Edges of the Prescription.
func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prescriptionpatient", PatientInfo.Type).Ref("patientprescription").Unique(),
		edge.From("prescriptiondoctor", Doctor.Type).Ref("doctorprescription").Unique(),
		edge.From("prescriptionmedicine", Medicine.Type).Ref("medicinepresciption").Unique(),
	}
}