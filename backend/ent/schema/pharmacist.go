package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Pharmacist holds the schema definition for the Pharmacist entity.
type Pharmacist struct {
	ent.Schema
}

// Fields of the Pharmacist.
func (Pharmacist) Fields() []ent.Field {
	return []ent.Field{
		field.String("email").NotEmpty().Unique(),
		field.String("password").NotEmpty(),
		field.String("name").NotEmpty(),
	}
}

// Edges of the Pharmacist.
func (Pharmacist) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dispensemedicine", DispenseMedicine.Type).StorageKey(edge.Column("pharmacist_id")),
		edge.To("drugallergys", DrugAllergy.Type).StorageKey(edge.Column("pharmacist_id")),
		edge.To("orderpharmacist", Order.Type),
		edge.To("Bills", Bill.Type).StorageKey(edge.Column("pharmacist_id")),
	}
}
