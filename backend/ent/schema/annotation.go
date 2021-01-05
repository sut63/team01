package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Annotation holds the schema definition for the Annotation entity.
type Annotation struct {
	ent.Schema
}

// Fields of the Annotation.
func (Annotation) Fields() []ent.Field {
	return []ent.Field{
		field.String("messages").NotEmpty(),
	}
}

// Edges of the Annotation.
func (Annotation) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("dispensemedicine", DispenseMedicine.Type).StorageKey(edge.Column("annotation_id")),
	}
}
