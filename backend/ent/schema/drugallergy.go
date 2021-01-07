package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// DrugAllergy holds the schema definition for the DrugAllergy entity.
type DrugAllergy struct {
	ent.Schema
}

// Fields of the DrugAllergy.
func (DrugAllergy) Fields() []ent.Field {
	return []ent.Field{
		field.Time("dateTime"),
	}
}

// Edges of the DrugAllergy.
func (DrugAllergy) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("patient", PatientInfo.Type).Ref("drugallergys").Unique(),
		edge.From("medicine", Medicine.Type).Ref("drugallergys").Unique(),
		edge.From("pharmacist", Pharmacist.Type).Ref("drugallergys").Unique(),
	}
}
