package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Order holds the schema definition for the Order entity.
type Order struct {
	ent.Schema
}

// Fields of the Order.
func (Order) Fields() []ent.Field {
	return []ent.Field{
		field.Int("amount").Positive(),
		field.Float("price"),
        field.Float("total"),
		field.Time("datetime"),
	}
}

// Edges of the Order.
func (Order) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pharmacist", Pharmacist.Type).Ref("order").Unique(),
		edge.From("medicine", Medicine.Type).Ref("order").Unique(),
		edge.From("company", Company.Type).Ref("order").Unique(),
	}
}

