package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Bill holds the schema definition for the Bill entity.
type Bill struct {
	ent.Schema
}

// Fields of the Bill.
func (Bill) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").Positive(),
		field.String("annotation").NotEmpty(),
		field.String("payer").NotEmpty(),

	}
}

// Edges of the Bill.
func (Bill) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pharmacists", Pharmacist.Type).Ref("Bills").Unique(),
		edge.From("payments", Payment.Type).Ref("Bills").Unique(),
		edge.From("dispenseMedicines", DispenseMedicine.Type).Ref("Bills").Unique(),
	}
}
