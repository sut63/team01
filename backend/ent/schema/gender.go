package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// Gender holds the schema definition for the Gender entity.
type Gender struct {
	ent.Schema
}

// Fields of the Gender.
func (Gender) Fields() []ent.Field {
	return []ent.Field{
		field.String("genderName").NotEmpty().Unique(),
	}
}

// Edges of the Gender.
func (Gender) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("patientinfo", PatientInfo.Type).StorageKey(edge.Column("gender_id")).Unique(),
	}
}
