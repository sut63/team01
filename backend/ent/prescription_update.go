// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/doctor"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/predicate"
	"github.com/sut63/team01/ent/prescription"
)

// PrescriptionUpdate is the builder for updating Prescription entities.
type PrescriptionUpdate struct {
	config
	hooks      []Hook
	mutation   *PrescriptionMutation
	predicates []predicate.Prescription
}

// Where adds a new predicate for the builder.
func (pu *PrescriptionUpdate) Where(ps ...predicate.Prescription) *PrescriptionUpdate {
	pu.predicates = append(pu.predicates, ps...)
	return pu
}

// SetValue sets the value field.
func (pu *PrescriptionUpdate) SetValue(i int) *PrescriptionUpdate {
	pu.mutation.ResetValue()
	pu.mutation.SetValue(i)
	return pu
}

// AddValue adds i to value.
func (pu *PrescriptionUpdate) AddValue(i int) *PrescriptionUpdate {
	pu.mutation.AddValue(i)
	return pu
}

// SetPrescriptionpatientID sets the prescriptionpatient edge to PatientInfo by id.
func (pu *PrescriptionUpdate) SetPrescriptionpatientID(id int) *PrescriptionUpdate {
	pu.mutation.SetPrescriptionpatientID(id)
	return pu
}

// SetNillablePrescriptionpatientID sets the prescriptionpatient edge to PatientInfo by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillablePrescriptionpatientID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetPrescriptionpatientID(*id)
	}
	return pu
}

// SetPrescriptionpatient sets the prescriptionpatient edge to PatientInfo.
func (pu *PrescriptionUpdate) SetPrescriptionpatient(p *PatientInfo) *PrescriptionUpdate {
	return pu.SetPrescriptionpatientID(p.ID)
}

// SetPrescriptiondoctorID sets the prescriptiondoctor edge to Doctor by id.
func (pu *PrescriptionUpdate) SetPrescriptiondoctorID(id int) *PrescriptionUpdate {
	pu.mutation.SetPrescriptiondoctorID(id)
	return pu
}

// SetNillablePrescriptiondoctorID sets the prescriptiondoctor edge to Doctor by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillablePrescriptiondoctorID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetPrescriptiondoctorID(*id)
	}
	return pu
}

// SetPrescriptiondoctor sets the prescriptiondoctor edge to Doctor.
func (pu *PrescriptionUpdate) SetPrescriptiondoctor(d *Doctor) *PrescriptionUpdate {
	return pu.SetPrescriptiondoctorID(d.ID)
}

// SetPrescriptionmedicineID sets the prescriptionmedicine edge to Medicine by id.
func (pu *PrescriptionUpdate) SetPrescriptionmedicineID(id int) *PrescriptionUpdate {
	pu.mutation.SetPrescriptionmedicineID(id)
	return pu
}

// SetNillablePrescriptionmedicineID sets the prescriptionmedicine edge to Medicine by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillablePrescriptionmedicineID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetPrescriptionmedicineID(*id)
	}
	return pu
}

// SetPrescriptionmedicine sets the prescriptionmedicine edge to Medicine.
func (pu *PrescriptionUpdate) SetPrescriptionmedicine(m *Medicine) *PrescriptionUpdate {
	return pu.SetPrescriptionmedicineID(m.ID)
}

// SetDispensemedicineID sets the dispensemedicine edge to DispenseMedicine by id.
func (pu *PrescriptionUpdate) SetDispensemedicineID(id int) *PrescriptionUpdate {
	pu.mutation.SetDispensemedicineID(id)
	return pu
}

// SetNillableDispensemedicineID sets the dispensemedicine edge to DispenseMedicine by id if the given value is not nil.
func (pu *PrescriptionUpdate) SetNillableDispensemedicineID(id *int) *PrescriptionUpdate {
	if id != nil {
		pu = pu.SetDispensemedicineID(*id)
	}
	return pu
}

// SetDispensemedicine sets the dispensemedicine edge to DispenseMedicine.
func (pu *PrescriptionUpdate) SetDispensemedicine(d *DispenseMedicine) *PrescriptionUpdate {
	return pu.SetDispensemedicineID(d.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (pu *PrescriptionUpdate) Mutation() *PrescriptionMutation {
	return pu.mutation
}

// ClearPrescriptionpatient clears the prescriptionpatient edge to PatientInfo.
func (pu *PrescriptionUpdate) ClearPrescriptionpatient() *PrescriptionUpdate {
	pu.mutation.ClearPrescriptionpatient()
	return pu
}

// ClearPrescriptiondoctor clears the prescriptiondoctor edge to Doctor.
func (pu *PrescriptionUpdate) ClearPrescriptiondoctor() *PrescriptionUpdate {
	pu.mutation.ClearPrescriptiondoctor()
	return pu
}

// ClearPrescriptionmedicine clears the prescriptionmedicine edge to Medicine.
func (pu *PrescriptionUpdate) ClearPrescriptionmedicine() *PrescriptionUpdate {
	pu.mutation.ClearPrescriptionmedicine()
	return pu
}

// ClearDispensemedicine clears the dispensemedicine edge to DispenseMedicine.
func (pu *PrescriptionUpdate) ClearDispensemedicine() *PrescriptionUpdate {
	pu.mutation.ClearDispensemedicine()
	return pu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (pu *PrescriptionUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(pu.hooks) == 0 {
		affected, err = pu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pu.mutation = mutation
			affected, err = pu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pu.hooks) - 1; i >= 0; i-- {
			mut = pu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pu *PrescriptionUpdate) SaveX(ctx context.Context) int {
	affected, err := pu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pu *PrescriptionUpdate) Exec(ctx context.Context) error {
	_, err := pu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pu *PrescriptionUpdate) ExecX(ctx context.Context) {
	if err := pu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pu *PrescriptionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prescription.Table,
			Columns: prescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		},
	}
	if ps := pu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pu.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: prescription.FieldValue,
		})
	}
	if value, ok := pu.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: prescription.FieldValue,
		})
	}
	if pu.mutation.PrescriptionpatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionpatientTable,
			Columns: []string{prescription.PrescriptionpatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrescriptionpatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionpatientTable,
			Columns: []string{prescription.PrescriptionpatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PrescriptiondoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptiondoctorTable,
			Columns: []string{prescription.PrescriptiondoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrescriptiondoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptiondoctorTable,
			Columns: []string{prescription.PrescriptiondoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.PrescriptionmedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionmedicineTable,
			Columns: []string{prescription.PrescriptionmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.PrescriptionmedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionmedicineTable,
			Columns: []string{prescription.PrescriptionmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pu.mutation.DispensemedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescription.DispensemedicineTable,
			Columns: []string{prescription.DispensemedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispensemedicine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pu.mutation.DispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescription.DispensemedicineTable,
			Columns: []string{prescription.DispensemedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispensemedicine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescription.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// PrescriptionUpdateOne is the builder for updating a single Prescription entity.
type PrescriptionUpdateOne struct {
	config
	hooks    []Hook
	mutation *PrescriptionMutation
}

// SetValue sets the value field.
func (puo *PrescriptionUpdateOne) SetValue(i int) *PrescriptionUpdateOne {
	puo.mutation.ResetValue()
	puo.mutation.SetValue(i)
	return puo
}

// AddValue adds i to value.
func (puo *PrescriptionUpdateOne) AddValue(i int) *PrescriptionUpdateOne {
	puo.mutation.AddValue(i)
	return puo
}

// SetPrescriptionpatientID sets the prescriptionpatient edge to PatientInfo by id.
func (puo *PrescriptionUpdateOne) SetPrescriptionpatientID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetPrescriptionpatientID(id)
	return puo
}

// SetNillablePrescriptionpatientID sets the prescriptionpatient edge to PatientInfo by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillablePrescriptionpatientID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetPrescriptionpatientID(*id)
	}
	return puo
}

// SetPrescriptionpatient sets the prescriptionpatient edge to PatientInfo.
func (puo *PrescriptionUpdateOne) SetPrescriptionpatient(p *PatientInfo) *PrescriptionUpdateOne {
	return puo.SetPrescriptionpatientID(p.ID)
}

// SetPrescriptiondoctorID sets the prescriptiondoctor edge to Doctor by id.
func (puo *PrescriptionUpdateOne) SetPrescriptiondoctorID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetPrescriptiondoctorID(id)
	return puo
}

// SetNillablePrescriptiondoctorID sets the prescriptiondoctor edge to Doctor by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillablePrescriptiondoctorID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetPrescriptiondoctorID(*id)
	}
	return puo
}

// SetPrescriptiondoctor sets the prescriptiondoctor edge to Doctor.
func (puo *PrescriptionUpdateOne) SetPrescriptiondoctor(d *Doctor) *PrescriptionUpdateOne {
	return puo.SetPrescriptiondoctorID(d.ID)
}

// SetPrescriptionmedicineID sets the prescriptionmedicine edge to Medicine by id.
func (puo *PrescriptionUpdateOne) SetPrescriptionmedicineID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetPrescriptionmedicineID(id)
	return puo
}

// SetNillablePrescriptionmedicineID sets the prescriptionmedicine edge to Medicine by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillablePrescriptionmedicineID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetPrescriptionmedicineID(*id)
	}
	return puo
}

// SetPrescriptionmedicine sets the prescriptionmedicine edge to Medicine.
func (puo *PrescriptionUpdateOne) SetPrescriptionmedicine(m *Medicine) *PrescriptionUpdateOne {
	return puo.SetPrescriptionmedicineID(m.ID)
}

// SetDispensemedicineID sets the dispensemedicine edge to DispenseMedicine by id.
func (puo *PrescriptionUpdateOne) SetDispensemedicineID(id int) *PrescriptionUpdateOne {
	puo.mutation.SetDispensemedicineID(id)
	return puo
}

// SetNillableDispensemedicineID sets the dispensemedicine edge to DispenseMedicine by id if the given value is not nil.
func (puo *PrescriptionUpdateOne) SetNillableDispensemedicineID(id *int) *PrescriptionUpdateOne {
	if id != nil {
		puo = puo.SetDispensemedicineID(*id)
	}
	return puo
}

// SetDispensemedicine sets the dispensemedicine edge to DispenseMedicine.
func (puo *PrescriptionUpdateOne) SetDispensemedicine(d *DispenseMedicine) *PrescriptionUpdateOne {
	return puo.SetDispensemedicineID(d.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (puo *PrescriptionUpdateOne) Mutation() *PrescriptionMutation {
	return puo.mutation
}

// ClearPrescriptionpatient clears the prescriptionpatient edge to PatientInfo.
func (puo *PrescriptionUpdateOne) ClearPrescriptionpatient() *PrescriptionUpdateOne {
	puo.mutation.ClearPrescriptionpatient()
	return puo
}

// ClearPrescriptiondoctor clears the prescriptiondoctor edge to Doctor.
func (puo *PrescriptionUpdateOne) ClearPrescriptiondoctor() *PrescriptionUpdateOne {
	puo.mutation.ClearPrescriptiondoctor()
	return puo
}

// ClearPrescriptionmedicine clears the prescriptionmedicine edge to Medicine.
func (puo *PrescriptionUpdateOne) ClearPrescriptionmedicine() *PrescriptionUpdateOne {
	puo.mutation.ClearPrescriptionmedicine()
	return puo
}

// ClearDispensemedicine clears the dispensemedicine edge to DispenseMedicine.
func (puo *PrescriptionUpdateOne) ClearDispensemedicine() *PrescriptionUpdateOne {
	puo.mutation.ClearDispensemedicine()
	return puo
}

// Save executes the query and returns the updated entity.
func (puo *PrescriptionUpdateOne) Save(ctx context.Context) (*Prescription, error) {

	var (
		err  error
		node *Prescription
	)
	if len(puo.hooks) == 0 {
		node, err = puo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			puo.mutation = mutation
			node, err = puo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(puo.hooks) - 1; i >= 0; i-- {
			mut = puo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, puo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (puo *PrescriptionUpdateOne) SaveX(ctx context.Context) *Prescription {
	pr, err := puo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return pr
}

// Exec executes the query on the entity.
func (puo *PrescriptionUpdateOne) Exec(ctx context.Context) error {
	_, err := puo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (puo *PrescriptionUpdateOne) ExecX(ctx context.Context) {
	if err := puo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (puo *PrescriptionUpdateOne) sqlSave(ctx context.Context) (pr *Prescription, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   prescription.Table,
			Columns: prescription.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		},
	}
	id, ok := puo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Prescription.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := puo.mutation.Value(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: prescription.FieldValue,
		})
	}
	if value, ok := puo.mutation.AddedValue(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: prescription.FieldValue,
		})
	}
	if puo.mutation.PrescriptionpatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionpatientTable,
			Columns: []string{prescription.PrescriptionpatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientinfo.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrescriptionpatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionpatientTable,
			Columns: []string{prescription.PrescriptionpatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PrescriptiondoctorCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptiondoctorTable,
			Columns: []string{prescription.PrescriptiondoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrescriptiondoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptiondoctorTable,
			Columns: []string{prescription.PrescriptiondoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.PrescriptionmedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionmedicineTable,
			Columns: []string{prescription.PrescriptionmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.PrescriptionmedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionmedicineTable,
			Columns: []string{prescription.PrescriptionmedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if puo.mutation.DispensemedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescription.DispensemedicineTable,
			Columns: []string{prescription.DispensemedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispensemedicine.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := puo.mutation.DispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescription.DispensemedicineTable,
			Columns: []string{prescription.DispensemedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: dispensemedicine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	pr = &Prescription{config: puo.config}
	_spec.Assign = pr.assignValues
	_spec.ScanValues = pr.scanValues()
	if err = sqlgraph.UpdateNode(ctx, puo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{prescription.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return pr, nil
}
