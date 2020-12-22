package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
 )

// NameTitle holds the schema definition for the NameTitle entity.
type NameTitle struct {
	ent.Schema
}

// Fields of the NameTitle.
func (NameTitle) Fields() []ent.Field {
	return []ent.Field{
		field.String("nameTitle").NotEmpty().Unique(),
	}
}

// Edges of the NameTitle.
func (NameTitle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("patientinfo", PatientInfo.Type).StorageKey(edge.Column("nametitle_id")).Unique(),
	}
}
