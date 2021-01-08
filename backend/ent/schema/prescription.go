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
<<<<<<< HEAD
		field.Int("Value"),
=======
		field.Int("value"),
>>>>>>> 68357c40a42f59ad7daf6e1c4e68a6b0e7ac2913
	}
}

// Edges of the Prescription.
func (Prescription) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("prescriptionpatient", PatientInfo.Type).Ref("patientprescription").Unique(),
		edge.From("prescriptiondoctor", Doctor.Type).Ref("doctorprescription").Unique(),
		edge.From("prescriptionmedicine", Medicine.Type).Ref("medicinepresciption").Unique(),
		edge.To("dispensemedicine", DispenseMedicine.Type).Unique().StorageKey(edge.Column("prescription_id")),
	}
}
