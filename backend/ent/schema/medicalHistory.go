package schema

import (
	"github.com/facebook/ent"
	//"github.com/facebook/ent/schema/field"
	//"github.com/facebook/ent/schema/edge"

)
// MedicalHistory holds the schema definition for the MedicalHistory entity.
type MedicalHistory struct {
	ent.Schema
}

// Fields of the MedicalHistory.
func (MedicalHistory) Fields() []ent.Field {
	return []ent.Field{
		//field.String("").NotEmpty(),
	}
}