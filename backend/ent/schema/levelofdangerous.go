package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)
// LevelOfDangerous holds the schema definition for the LevelOfDangerous entity.
type LevelOfDangerous struct {
	ent.Schema
}

// Fields of the LevelOfDangerous.
func (LevelOfDangerous) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
	}
}

// Edges of the LevelOfDangerous.
func (LevelOfDangerous) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("Medicine", Medicine.Type),
	}
}
