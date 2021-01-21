package schema

import (
	"errors"
	"regexp"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// DispenseMedicine holds the schema definition for the DispenseMedicine entity.
type DispenseMedicine struct {
	ent.Schema
}

// Fields of the DispenseMedicine.
func (DispenseMedicine) Fields() []ent.Field {
	return []ent.Field{
		field.Time("datetime"),
		field.String("note").Validate(func(s string) error {
			match, _ := regexp.Match("[`~!@#$%^&*_;?<>]", []byte(s))
			if match {
				return errors.New("รูปแบบรายละอียดหมายเหตุไม่ถูกต้องห้ามมีอักษรพิเศษ")
			}
			return nil
		}).NotEmpty(),
		field.Int("amountchangemedicine").Range(0, 20),
		field.String("detailchangemedicine").Validate(func(s string) error {
			match, _ := regexp.Match("[`~!@#$%^&*_;?<>]", []byte(s))
			if match {
				return errors.New("รูปแบบรายละอียดการเปลี่ยนยาไม่ถูกต้องห้ามมีอักษรพิเศษ")
			}
			return nil
		}).NotEmpty(),
	}
}

// Edges of the DispenseMedicine.
func (DispenseMedicine) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("pharmacist", Pharmacist.Type).Ref("dispensemedicine").Unique(),
		edge.From("annotation", Annotation.Type).Ref("dispensemedicine").Unique(),
		edge.From("prescription", Prescription.Type).Ref("dispensemedicine").Unique(),
		edge.To("Bills", Bill.Type).StorageKey(edge.Column("dispensemedicine_id")).Unique(),
	}
}
