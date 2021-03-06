package schema

import (
	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Medicine holds the schema definition for the Medicine entity.
type Medicine struct {
	ent.Schema
}

// Fields of the Medicine.
func (Medicine) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").NotEmpty(),
		field.String("serial").NotEmpty(),
		field.String("brand").NotEmpty(),
		field.Int("amount").Positive(),
		field.Int("price").Positive(),
		field.String("howtouse").NotEmpty(),
	}
}

// Edges of the Medicine.
func (Medicine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("LevelOfDangerous", LevelOfDangerous.Type).Ref("Medicine").Unique(),
		edge.From("MedicineType", MedicineType.Type).Ref("Medicine").Unique(),
		edge.From("UnitOfMedicine", UnitOfMedicine.Type).Ref("Medicine").Unique(),
		edge.To("drugallergys", DrugAllergy.Type).StorageKey(edge.Column("medicine_id")),
		edge.To("medicinepresciption", Prescription.Type).StorageKey(edge.Column("medicine_id")),
		edge.To("ordermedicine", Order.Type),
	}
}
