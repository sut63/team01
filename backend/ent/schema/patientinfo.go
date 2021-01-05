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
		field.String("gender").NotEmpty(),
		field.Int("age").Positive(),
	}
}

// Edges of the PatientInfo.
func (PatientInfo) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("drugallergys", DrugAllergy.Type).StorageKey(edge.Column("patient_id")),
	}
}
