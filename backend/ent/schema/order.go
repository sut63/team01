package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/facebookincubator/ent/schema/edge"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Time("addedtime"),
		field.Int("price"),
		field.Int("amount"),

	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("medicine", Medicine.Type).
		Ref("ordermedicine").
		Unique(),
		edge.From("company", Company.Type).
		Ref("ordercompany").
		Unique(),
		edge.From("pharmacist",Pharmacist.Type).
		Ref("orderpharmacist").
		Unique(),
	}
}

