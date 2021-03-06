// Code generated by entc, DO NOT EDIT.

package dispensemedicine

import (
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team01/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
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
func IDGT(id int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Datetime applies equality check predicate on the "datetime" field. It's identical to DatetimeEQ.
func Datetime(v time.Time) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetime), v))
	})
}

// Note applies equality check predicate on the "note" field. It's identical to NoteEQ.
func Note(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// Amountchangemedicine applies equality check predicate on the "amountchangemedicine" field. It's identical to AmountchangemedicineEQ.
func Amountchangemedicine(v int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmountchangemedicine), v))
	})
}

// Detailchangemedicine applies equality check predicate on the "detailchangemedicine" field. It's identical to DetailchangemedicineEQ.
func Detailchangemedicine(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetailchangemedicine), v))
	})
}

// DatetimeEQ applies the EQ predicate on the "datetime" field.
func DatetimeEQ(v time.Time) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDatetime), v))
	})
}

// DatetimeNEQ applies the NEQ predicate on the "datetime" field.
func DatetimeNEQ(v time.Time) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDatetime), v))
	})
}

// DatetimeIn applies the In predicate on the "datetime" field.
func DatetimeIn(vs ...time.Time) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDatetime), v...))
	})
}

// DatetimeNotIn applies the NotIn predicate on the "datetime" field.
func DatetimeNotIn(vs ...time.Time) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDatetime), v...))
	})
}

// DatetimeGT applies the GT predicate on the "datetime" field.
func DatetimeGT(v time.Time) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDatetime), v))
	})
}

// DatetimeGTE applies the GTE predicate on the "datetime" field.
func DatetimeGTE(v time.Time) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDatetime), v))
	})
}

// DatetimeLT applies the LT predicate on the "datetime" field.
func DatetimeLT(v time.Time) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDatetime), v))
	})
}

// DatetimeLTE applies the LTE predicate on the "datetime" field.
func DatetimeLTE(v time.Time) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDatetime), v))
	})
}

// NoteEQ applies the EQ predicate on the "note" field.
func NoteEQ(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldNote), v))
	})
}

// NoteNEQ applies the NEQ predicate on the "note" field.
func NoteNEQ(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldNote), v))
	})
}

// NoteIn applies the In predicate on the "note" field.
func NoteIn(vs ...string) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldNote), v...))
	})
}

// NoteNotIn applies the NotIn predicate on the "note" field.
func NoteNotIn(vs ...string) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldNote), v...))
	})
}

// NoteGT applies the GT predicate on the "note" field.
func NoteGT(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldNote), v))
	})
}

// NoteGTE applies the GTE predicate on the "note" field.
func NoteGTE(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldNote), v))
	})
}

// NoteLT applies the LT predicate on the "note" field.
func NoteLT(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldNote), v))
	})
}

// NoteLTE applies the LTE predicate on the "note" field.
func NoteLTE(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldNote), v))
	})
}

// NoteContains applies the Contains predicate on the "note" field.
func NoteContains(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldNote), v))
	})
}

// NoteHasPrefix applies the HasPrefix predicate on the "note" field.
func NoteHasPrefix(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldNote), v))
	})
}

// NoteHasSuffix applies the HasSuffix predicate on the "note" field.
func NoteHasSuffix(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldNote), v))
	})
}

// NoteEqualFold applies the EqualFold predicate on the "note" field.
func NoteEqualFold(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldNote), v))
	})
}

// NoteContainsFold applies the ContainsFold predicate on the "note" field.
func NoteContainsFold(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldNote), v))
	})
}

// AmountchangemedicineEQ applies the EQ predicate on the "amountchangemedicine" field.
func AmountchangemedicineEQ(v int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAmountchangemedicine), v))
	})
}

// AmountchangemedicineNEQ applies the NEQ predicate on the "amountchangemedicine" field.
func AmountchangemedicineNEQ(v int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAmountchangemedicine), v))
	})
}

// AmountchangemedicineIn applies the In predicate on the "amountchangemedicine" field.
func AmountchangemedicineIn(vs ...int) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAmountchangemedicine), v...))
	})
}

// AmountchangemedicineNotIn applies the NotIn predicate on the "amountchangemedicine" field.
func AmountchangemedicineNotIn(vs ...int) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAmountchangemedicine), v...))
	})
}

// AmountchangemedicineGT applies the GT predicate on the "amountchangemedicine" field.
func AmountchangemedicineGT(v int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAmountchangemedicine), v))
	})
}

// AmountchangemedicineGTE applies the GTE predicate on the "amountchangemedicine" field.
func AmountchangemedicineGTE(v int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAmountchangemedicine), v))
	})
}

// AmountchangemedicineLT applies the LT predicate on the "amountchangemedicine" field.
func AmountchangemedicineLT(v int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAmountchangemedicine), v))
	})
}

// AmountchangemedicineLTE applies the LTE predicate on the "amountchangemedicine" field.
func AmountchangemedicineLTE(v int) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAmountchangemedicine), v))
	})
}

// DetailchangemedicineEQ applies the EQ predicate on the "detailchangemedicine" field.
func DetailchangemedicineEQ(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineNEQ applies the NEQ predicate on the "detailchangemedicine" field.
func DetailchangemedicineNEQ(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineIn applies the In predicate on the "detailchangemedicine" field.
func DetailchangemedicineIn(vs ...string) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDetailchangemedicine), v...))
	})
}

// DetailchangemedicineNotIn applies the NotIn predicate on the "detailchangemedicine" field.
func DetailchangemedicineNotIn(vs ...string) predicate.DispenseMedicine {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDetailchangemedicine), v...))
	})
}

// DetailchangemedicineGT applies the GT predicate on the "detailchangemedicine" field.
func DetailchangemedicineGT(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineGTE applies the GTE predicate on the "detailchangemedicine" field.
func DetailchangemedicineGTE(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineLT applies the LT predicate on the "detailchangemedicine" field.
func DetailchangemedicineLT(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineLTE applies the LTE predicate on the "detailchangemedicine" field.
func DetailchangemedicineLTE(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineContains applies the Contains predicate on the "detailchangemedicine" field.
func DetailchangemedicineContains(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineHasPrefix applies the HasPrefix predicate on the "detailchangemedicine" field.
func DetailchangemedicineHasPrefix(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineHasSuffix applies the HasSuffix predicate on the "detailchangemedicine" field.
func DetailchangemedicineHasSuffix(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineEqualFold applies the EqualFold predicate on the "detailchangemedicine" field.
func DetailchangemedicineEqualFold(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDetailchangemedicine), v))
	})
}

// DetailchangemedicineContainsFold applies the ContainsFold predicate on the "detailchangemedicine" field.
func DetailchangemedicineContainsFold(v string) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDetailchangemedicine), v))
	})
}

// HasPharmacist applies the HasEdge predicate on the "pharmacist" edge.
func HasPharmacist() predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PharmacistTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PharmacistTable, PharmacistColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPharmacistWith applies the HasEdge predicate on the "pharmacist" edge with a given conditions (other predicates).
func HasPharmacistWith(preds ...predicate.Pharmacist) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PharmacistInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PharmacistTable, PharmacistColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasAnnotation applies the HasEdge predicate on the "annotation" edge.
func HasAnnotation() predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AnnotationTable, AnnotationColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasAnnotationWith applies the HasEdge predicate on the "annotation" edge with a given conditions (other predicates).
func HasAnnotationWith(preds ...predicate.Annotation) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(AnnotationInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, AnnotationTable, AnnotationColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrescription applies the HasEdge predicate on the "prescription" edge.
func HasPrescription() predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptionTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrescriptionTable, PrescriptionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrescriptionWith applies the HasEdge predicate on the "prescription" edge with a given conditions (other predicates).
func HasPrescriptionWith(preds ...predicate.Prescription) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, true, PrescriptionTable, PrescriptionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasBills applies the HasEdge predicate on the "Bills" edge.
func HasBills() predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillsTable, BillsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasBillsWith applies the HasEdge predicate on the "Bills" edge with a given conditions (other predicates).
func HasBillsWith(preds ...predicate.Bill) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(BillsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, BillsTable, BillsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.DispenseMedicine) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.DispenseMedicine) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
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
func Not(p predicate.DispenseMedicine) predicate.DispenseMedicine {
	return predicate.DispenseMedicine(func(s *sql.Selector) {
		p(s.Not())
	})
}
