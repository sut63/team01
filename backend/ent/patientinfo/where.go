// Code generated by entc, DO NOT EDIT.

package patientinfo

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team01/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// CardNumber applies equality check predicate on the "cardNumber" field. It's identical to CardNumberEQ.
func CardNumber(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardNumber), v))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// Gender applies equality check predicate on the "gender" field. It's identical to GenderEQ.
func Gender(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// CardNumberEQ applies the EQ predicate on the "cardNumber" field.
func CardNumberEQ(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCardNumber), v))
	})
}

// CardNumberNEQ applies the NEQ predicate on the "cardNumber" field.
func CardNumberNEQ(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCardNumber), v))
	})
}

// CardNumberIn applies the In predicate on the "cardNumber" field.
func CardNumberIn(vs ...string) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCardNumber), v...))
	})
}

// CardNumberNotIn applies the NotIn predicate on the "cardNumber" field.
func CardNumberNotIn(vs ...string) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCardNumber), v...))
	})
}

// CardNumberGT applies the GT predicate on the "cardNumber" field.
func CardNumberGT(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCardNumber), v))
	})
}

// CardNumberGTE applies the GTE predicate on the "cardNumber" field.
func CardNumberGTE(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCardNumber), v))
	})
}

// CardNumberLT applies the LT predicate on the "cardNumber" field.
func CardNumberLT(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCardNumber), v))
	})
}

// CardNumberLTE applies the LTE predicate on the "cardNumber" field.
func CardNumberLTE(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCardNumber), v))
	})
}

// CardNumberContains applies the Contains predicate on the "cardNumber" field.
func CardNumberContains(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldCardNumber), v))
	})
}

// CardNumberHasPrefix applies the HasPrefix predicate on the "cardNumber" field.
func CardNumberHasPrefix(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldCardNumber), v))
	})
}

// CardNumberHasSuffix applies the HasSuffix predicate on the "cardNumber" field.
func CardNumberHasSuffix(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldCardNumber), v))
	})
}

// CardNumberEqualFold applies the EqualFold predicate on the "cardNumber" field.
func CardNumberEqualFold(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldCardNumber), v))
	})
}

// CardNumberContainsFold applies the ContainsFold predicate on the "cardNumber" field.
func CardNumberContainsFold(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldCardNumber), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	})
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGender), v))
	})
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...string) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGender), v...))
	})
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...string) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGender), v...))
	})
}

// GenderGT applies the GT predicate on the "gender" field.
func GenderGT(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldGender), v))
	})
}

// GenderGTE applies the GTE predicate on the "gender" field.
func GenderGTE(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldGender), v))
	})
}

// GenderLT applies the LT predicate on the "gender" field.
func GenderLT(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldGender), v))
	})
}

// GenderLTE applies the LTE predicate on the "gender" field.
func GenderLTE(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldGender), v))
	})
}

// GenderContains applies the Contains predicate on the "gender" field.
func GenderContains(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldGender), v))
	})
}

// GenderHasPrefix applies the HasPrefix predicate on the "gender" field.
func GenderHasPrefix(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldGender), v))
	})
}

// GenderHasSuffix applies the HasSuffix predicate on the "gender" field.
func GenderHasSuffix(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldGender), v))
	})
}

// GenderEqualFold applies the EqualFold predicate on the "gender" field.
func GenderEqualFold(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldGender), v))
	})
}

// GenderContainsFold applies the ContainsFold predicate on the "gender" field.
func GenderContainsFold(v string) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldGender), v))
	})
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	})
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	})
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	})
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.PatientInfo {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.PatientInfo(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	})
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	})
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	})
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	})
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	})
}

// HasDrugallergys applies the HasEdge predicate on the "drugallergys" edge.
func HasDrugallergys() predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DrugallergysTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DrugallergysTable, DrugallergysColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDrugallergysWith applies the HasEdge predicate on the "drugallergys" edge with a given conditions (other predicates).
func HasDrugallergysWith(preds ...predicate.DrugAllergy) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DrugallergysInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, DrugallergysTable, DrugallergysColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPatientprescription applies the HasEdge predicate on the "patientprescription" edge.
func HasPatientprescription() predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientprescriptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientprescriptionTable, PatientprescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPatientprescriptionWith applies the HasEdge predicate on the "patientprescription" edge with a given conditions (other predicates).
func HasPatientprescriptionWith(preds ...predicate.Prescription) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PatientprescriptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PatientprescriptionTable, PatientprescriptionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.PatientInfo) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.PatientInfo) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.PatientInfo) predicate.PatientInfo {
	return predicate.PatientInfo(func(s *sql.Selector) {
		p(s.Not())
	})
}