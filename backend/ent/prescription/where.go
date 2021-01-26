// Code generated by entc, DO NOT EDIT.

package prescription

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/sut63/team01/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Value applies equality check predicate on the "Value" field. It's identical to ValueEQ.
func Value(v int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// Symptom applies equality check predicate on the "Symptom" field. It's identical to SymptomEQ.
func Symptom(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// Annotation applies equality check predicate on the "Annotation" field. It's identical to AnnotationEQ.
func Annotation(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnnotation), v))
	})
}

// ValueEQ applies the EQ predicate on the "Value" field.
func ValueEQ(v int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "Value" field.
func ValueNEQ(v int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "Value" field.
func ValueIn(vs ...int) predicate.Prescription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prescription(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "Value" field.
func ValueNotIn(vs ...int) predicate.Prescription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prescription(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "Value" field.
func ValueGT(v int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "Value" field.
func ValueGTE(v int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "Value" field.
func ValueLT(v int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "Value" field.
func ValueLTE(v int) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// SymptomEQ applies the EQ predicate on the "Symptom" field.
func SymptomEQ(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSymptom), v))
	})
}

// SymptomNEQ applies the NEQ predicate on the "Symptom" field.
func SymptomNEQ(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSymptom), v))
	})
}

// SymptomIn applies the In predicate on the "Symptom" field.
func SymptomIn(vs ...string) predicate.Prescription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prescription(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSymptom), v...))
	})
}

// SymptomNotIn applies the NotIn predicate on the "Symptom" field.
func SymptomNotIn(vs ...string) predicate.Prescription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prescription(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSymptom), v...))
	})
}

// SymptomGT applies the GT predicate on the "Symptom" field.
func SymptomGT(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSymptom), v))
	})
}

// SymptomGTE applies the GTE predicate on the "Symptom" field.
func SymptomGTE(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSymptom), v))
	})
}

// SymptomLT applies the LT predicate on the "Symptom" field.
func SymptomLT(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSymptom), v))
	})
}

// SymptomLTE applies the LTE predicate on the "Symptom" field.
func SymptomLTE(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSymptom), v))
	})
}

// SymptomContains applies the Contains predicate on the "Symptom" field.
func SymptomContains(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSymptom), v))
	})
}

// SymptomHasPrefix applies the HasPrefix predicate on the "Symptom" field.
func SymptomHasPrefix(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSymptom), v))
	})
}

// SymptomHasSuffix applies the HasSuffix predicate on the "Symptom" field.
func SymptomHasSuffix(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSymptom), v))
	})
}

// SymptomEqualFold applies the EqualFold predicate on the "Symptom" field.
func SymptomEqualFold(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSymptom), v))
	})
}

// SymptomContainsFold applies the ContainsFold predicate on the "Symptom" field.
func SymptomContainsFold(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSymptom), v))
	})
}

// AnnotationEQ applies the EQ predicate on the "Annotation" field.
func AnnotationEQ(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAnnotation), v))
	})
}

// AnnotationNEQ applies the NEQ predicate on the "Annotation" field.
func AnnotationNEQ(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAnnotation), v))
	})
}

// AnnotationIn applies the In predicate on the "Annotation" field.
func AnnotationIn(vs ...string) predicate.Prescription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prescription(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAnnotation), v...))
	})
}

// AnnotationNotIn applies the NotIn predicate on the "Annotation" field.
func AnnotationNotIn(vs ...string) predicate.Prescription {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Prescription(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAnnotation), v...))
	})
}

// AnnotationGT applies the GT predicate on the "Annotation" field.
func AnnotationGT(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAnnotation), v))
	})
}

// AnnotationGTE applies the GTE predicate on the "Annotation" field.
func AnnotationGTE(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAnnotation), v))
	})
}

// AnnotationLT applies the LT predicate on the "Annotation" field.
func AnnotationLT(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAnnotation), v))
	})
}

// AnnotationLTE applies the LTE predicate on the "Annotation" field.
func AnnotationLTE(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAnnotation), v))
	})
}

// AnnotationContains applies the Contains predicate on the "Annotation" field.
func AnnotationContains(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAnnotation), v))
	})
}

// AnnotationHasPrefix applies the HasPrefix predicate on the "Annotation" field.
func AnnotationHasPrefix(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAnnotation), v))
	})
}

// AnnotationHasSuffix applies the HasSuffix predicate on the "Annotation" field.
func AnnotationHasSuffix(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAnnotation), v))
	})
}

// AnnotationEqualFold applies the EqualFold predicate on the "Annotation" field.
func AnnotationEqualFold(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAnnotation), v))
	})
}

// AnnotationContainsFold applies the ContainsFold predicate on the "Annotation" field.
func AnnotationContainsFold(v string) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAnnotation), v))
	})
}

// HasPrescriptionpatient applies the HasEdge predicate on the "prescriptionpatient" edge.
func HasPrescriptionpatient() predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptionpatientTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrescriptionpatientTable, PrescriptionpatientColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrescriptionpatientWith applies the HasEdge predicate on the "prescriptionpatient" edge with a given conditions (other predicates).
func HasPrescriptionpatientWith(preds ...predicate.PatientInfo) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptionpatientInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrescriptionpatientTable, PrescriptionpatientColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrescriptiondoctor applies the HasEdge predicate on the "prescriptiondoctor" edge.
func HasPrescriptiondoctor() predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptiondoctorTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrescriptiondoctorTable, PrescriptiondoctorColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrescriptiondoctorWith applies the HasEdge predicate on the "prescriptiondoctor" edge with a given conditions (other predicates).
func HasPrescriptiondoctorWith(preds ...predicate.Doctor) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptiondoctorInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrescriptiondoctorTable, PrescriptiondoctorColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasPrescriptionmedicine applies the HasEdge predicate on the "prescriptionmedicine" edge.
func HasPrescriptionmedicine() predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptionmedicineTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrescriptionmedicineTable, PrescriptionmedicineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPrescriptionmedicineWith applies the HasEdge predicate on the "prescriptionmedicine" edge with a given conditions (other predicates).
func HasPrescriptionmedicineWith(preds ...predicate.Medicine) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PrescriptionmedicineInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, PrescriptionmedicineTable, PrescriptionmedicineColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDispensemedicine applies the HasEdge predicate on the "dispensemedicine" edge.
func HasDispensemedicine() predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DispensemedicineTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, DispensemedicineTable, DispensemedicineColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDispensemedicineWith applies the HasEdge predicate on the "dispensemedicine" edge with a given conditions (other predicates).
func HasDispensemedicineWith(preds ...predicate.DispenseMedicine) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(DispensemedicineInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, DispensemedicineTable, DispensemedicineColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Prescription) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Prescription) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
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
func Not(p predicate.Prescription) predicate.Prescription {
	return predicate.Prescription(func(s *sql.Selector) {
		p(s.Not())
	})
}
