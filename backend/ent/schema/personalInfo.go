package schema

import (
	"github.com/facebook/ent"
	//"github.com/facebook/ent/schema/field"
	//"github.com/facebook/ent/schema/edge"

)
// PersonalInfo holds the schema definition for the PersonalInfo entity.
type PersonalInfo struct {
	ent.Schema
}

// Fields of the PersonalInfo.
func (PersonalInfo) Fields() []ent.Field {
	return []ent.Field{
		//field.String("name").NotEmpty(),

	}
}

