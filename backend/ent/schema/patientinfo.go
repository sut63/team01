package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// PatientInfo holds the schema definition for the PatientInfo entity.
type PatientInfo struct {
	ent.Schema
}

// Fields of the PatientInfo.
func (PatientInfo) Fields() []ent.Field {
	return []ent.Field{
		field.String("cardNumber").NotEmpty().Unique(),
		field.String("name").NotEmpty(),
		field.String("surname").NotEmpty(),
		field.Int("age").Positive(),
		field.String("address"),
		field.String("chronicDisease"),
		field.String("congenitalDisease"),
		field.String("drugAllergy"),
	}
}

// Edges of the PatientInfo.
func (PatientInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("nametitle", NameTitle.Type).Ref("patientinfo").Unique().Required(),
		edge.From("gender", Gender.Type).Ref("patientinfo").Unique().Required(),
		edge.From("bloodtype", BloodType.Type).Ref("patientinfo").Unique().Required(),
		edge.To("medicalhistorys", MedicalHistory.Type).StorageKey(edge.Column("owner_id")),
	}
}
