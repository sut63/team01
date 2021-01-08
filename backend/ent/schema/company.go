package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Company holds the schema definition for the Company entity.
type Company struct {
	ent.Schema
}

// Fields of the Company.
func (Company) Fields() []ent.Field {
	return []ent.Field{
		field.String("companyname").NotEmpty(),
	}
}

// Edges of the Company.
func (Company) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("order", Order.Type),
		
	}
}