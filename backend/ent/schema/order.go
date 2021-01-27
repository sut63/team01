package schema
import "regexp"
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
		field.String("hospitalid").Match(regexp.MustCompile("[0]\\d{8}")),
		field.Time("addedtime"),
		field.Int("price").Min(0).Positive(),
		field.Int("amount").Min(0).Positive(),
		

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