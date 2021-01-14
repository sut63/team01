package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// UnitOfMedicine holds the schema definition for the UnitOfMedicine entity.
type UnitOfMedicine struct {
	ent.Schema
}

// Fields of the UnitOfMedicine.
func (UnitOfMedicine) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").Unique(),
	}
}

// Edges of the UnitOfMedicine.
func (UnitOfMedicine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Medicine", Medicine.Type),
	}
}
