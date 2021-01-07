// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/drugallergy"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/predicate"
)

// DrugAllergyUpdate is the builder for updating DrugAllergy entities.
type DrugAllergyUpdate struct {
	config
	hooks      []Hook
	mutation   *DrugAllergyMutation
	predicates []predicate.DrugAllergy
}

// Where adds a new predicate for the builder.
func (dau *DrugAllergyUpdate) Where(ps ...predicate.DrugAllergy) *DrugAllergyUpdate {
	dau.predicates = append(dau.predicates, ps...)
	return dau
}

// SetDateTime sets the dateTime field.
func (dau *DrugAllergyUpdate) SetDateTime(t time.Time) *DrugAllergyUpdate {
	dau.mutation.SetDateTime(t)
	return dau
}

// SetPatientID sets the patient edge to PatientInfo by id.
func (dau *DrugAllergyUpdate) SetPatientID(id int) *DrugAllergyUpdate {
	dau.mutation.SetPatientID(id)
	return dau
}

// SetNillablePatientID sets the patient edge to PatientInfo by id if the given value is not nil.
func (dau *DrugAllergyUpdate) SetNillablePatientID(id *int) *DrugAllergyUpdate {
	if id != nil {
		dau = dau.SetPatientID(*id)
	}
	return dau
}

// SetPatient sets the patient edge to PatientInfo.
func (dau *DrugAllergyUpdate) SetPatient(p *PatientInfo) *DrugAllergyUpdate {
	return dau.SetPatientID(p.ID)
}

// SetMedicineID sets the medicine edge to Medicine by id.
func (dau *DrugAllergyUpdate) SetMedicineID(id int) *DrugAllergyUpdate {
	dau.mutation.SetMedicineID(id)
	return dau
}

// SetNillableMedicineID sets the medicine edge to Medicine by id if the given value is not nil.
func (dau *DrugAllergyUpdate) SetNillableMedicineID(id *int) *DrugAllergyUpdate {
	if id != nil {
		dau = dau.SetMedicineID(*id)
	}
	return dau
}

// SetMedicine sets the medicine edge to Medicine.
func (dau *DrugAllergyUpdate) SetMedicine(m *Medicine) *DrugAllergyUpdate {
	return dau.SetMedicineID(m.ID)
}

// SetPharmacistID sets the pharmacist edge to Pharmacist by id.
func (dau *DrugAllergyUpdate) SetPharmacistID(id int) *DrugAllergyUpdate {
	dau.mutation.SetPharmacistID(id)
	return dau
}

// SetNillablePharmacistID sets the pharmacist edge to Pharmacist by id if the given value is not nil.
func (dau *DrugAllergyUpdate) SetNillablePharmacistID(id *int) *DrugAllergyUpdate {
	if id != nil {
		dau = dau.SetPharmacistID(*id)
	}
	return dau
}

// SetPharmacist sets the pharmacist edge to Pharmacist.
func (dau *DrugAllergyUpdate) SetPharmacist(p *Pharmacist) *DrugAllergyUpdate {
	return dau.SetPharmacistID(p.ID)
}

// Mutation returns the DrugAllergyMutation object of the builder.
func (dau *DrugAllergyUpdate) Mutation() *DrugAllergyMutation {
	return dau.mutation
}

// ClearPatient clears the patient edge to PatientInfo.
func (dau *DrugAllergyUpdate) ClearPatient() *DrugAllergyUpdate {
	dau.mutation.ClearPatient()
	return dau
}

// ClearMedicine clears the medicine edge to Medicine.
func (dau *DrugAllergyUpdate) ClearMedicine() *DrugAllergyUpdate {
	dau.mutation.ClearMedicine()
	return dau
}

// ClearPharmacist clears the pharmacist edge to Pharmacist.
func (dau *DrugAllergyUpdate) ClearPharmacist() *DrugAllergyUpdate {
	dau.mutation.ClearPharmacist()
	return dau
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (dau *DrugAllergyUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(dau.hooks) == 0 {
		affected, err = dau.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DrugAllergyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dau.mutation = mutation
			affected, err = dau.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dau.hooks) - 1; i >= 0; i-- {
			mut = dau.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dau.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (dau *DrugAllergyUpdate) SaveX(ctx context.Context) int {
	affected, err := dau.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dau *DrugAllergyUpdate) Exec(ctx context.Context) error {
	_, err := dau.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dau *DrugAllergyUpdate) ExecX(ctx context.Context) {
	if err := dau.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dau *DrugAllergyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drugallergy.Table,
			Columns: drugallergy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drugallergy.FieldID,
			},
		},
	}
	if ps := dau.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dau.mutation.DateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: drugallergy.FieldDateTime,
		})
	}
	if dau.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PatientTable,
			Columns: []string{drugallergy.PatientColumn},
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
	if nodes := dau.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PatientTable,
			Columns: []string{drugallergy.PatientColumn},
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
	if dau.mutation.MedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.MedicineTable,
			Columns: []string{drugallergy.MedicineColumn},
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
	if nodes := dau.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.MedicineTable,
			Columns: []string{drugallergy.MedicineColumn},
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
	if dau.mutation.PharmacistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PharmacistTable,
			Columns: []string{drugallergy.PharmacistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pharmacist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dau.mutation.PharmacistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PharmacistTable,
			Columns: []string{drugallergy.PharmacistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pharmacist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, dau.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{drugallergy.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// DrugAllergyUpdateOne is the builder for updating a single DrugAllergy entity.
type DrugAllergyUpdateOne struct {
	config
	hooks    []Hook
	mutation *DrugAllergyMutation
}

// SetDateTime sets the dateTime field.
func (dauo *DrugAllergyUpdateOne) SetDateTime(t time.Time) *DrugAllergyUpdateOne {
	dauo.mutation.SetDateTime(t)
	return dauo
}

// SetPatientID sets the patient edge to PatientInfo by id.
func (dauo *DrugAllergyUpdateOne) SetPatientID(id int) *DrugAllergyUpdateOne {
	dauo.mutation.SetPatientID(id)
	return dauo
}

// SetNillablePatientID sets the patient edge to PatientInfo by id if the given value is not nil.
func (dauo *DrugAllergyUpdateOne) SetNillablePatientID(id *int) *DrugAllergyUpdateOne {
	if id != nil {
		dauo = dauo.SetPatientID(*id)
	}
	return dauo
}

// SetPatient sets the patient edge to PatientInfo.
func (dauo *DrugAllergyUpdateOne) SetPatient(p *PatientInfo) *DrugAllergyUpdateOne {
	return dauo.SetPatientID(p.ID)
}

// SetMedicineID sets the medicine edge to Medicine by id.
func (dauo *DrugAllergyUpdateOne) SetMedicineID(id int) *DrugAllergyUpdateOne {
	dauo.mutation.SetMedicineID(id)
	return dauo
}

// SetNillableMedicineID sets the medicine edge to Medicine by id if the given value is not nil.
func (dauo *DrugAllergyUpdateOne) SetNillableMedicineID(id *int) *DrugAllergyUpdateOne {
	if id != nil {
		dauo = dauo.SetMedicineID(*id)
	}
	return dauo
}

// SetMedicine sets the medicine edge to Medicine.
func (dauo *DrugAllergyUpdateOne) SetMedicine(m *Medicine) *DrugAllergyUpdateOne {
	return dauo.SetMedicineID(m.ID)
}

// SetPharmacistID sets the pharmacist edge to Pharmacist by id.
func (dauo *DrugAllergyUpdateOne) SetPharmacistID(id int) *DrugAllergyUpdateOne {
	dauo.mutation.SetPharmacistID(id)
	return dauo
}

// SetNillablePharmacistID sets the pharmacist edge to Pharmacist by id if the given value is not nil.
func (dauo *DrugAllergyUpdateOne) SetNillablePharmacistID(id *int) *DrugAllergyUpdateOne {
	if id != nil {
		dauo = dauo.SetPharmacistID(*id)
	}
	return dauo
}

// SetPharmacist sets the pharmacist edge to Pharmacist.
func (dauo *DrugAllergyUpdateOne) SetPharmacist(p *Pharmacist) *DrugAllergyUpdateOne {
	return dauo.SetPharmacistID(p.ID)
}

// Mutation returns the DrugAllergyMutation object of the builder.
func (dauo *DrugAllergyUpdateOne) Mutation() *DrugAllergyMutation {
	return dauo.mutation
}

// ClearPatient clears the patient edge to PatientInfo.
func (dauo *DrugAllergyUpdateOne) ClearPatient() *DrugAllergyUpdateOne {
	dauo.mutation.ClearPatient()
	return dauo
}

// ClearMedicine clears the medicine edge to Medicine.
func (dauo *DrugAllergyUpdateOne) ClearMedicine() *DrugAllergyUpdateOne {
	dauo.mutation.ClearMedicine()
	return dauo
}

// ClearPharmacist clears the pharmacist edge to Pharmacist.
func (dauo *DrugAllergyUpdateOne) ClearPharmacist() *DrugAllergyUpdateOne {
	dauo.mutation.ClearPharmacist()
	return dauo
}

// Save executes the query and returns the updated entity.
func (dauo *DrugAllergyUpdateOne) Save(ctx context.Context) (*DrugAllergy, error) {

	var (
		err  error
		node *DrugAllergy
	)
	if len(dauo.hooks) == 0 {
		node, err = dauo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DrugAllergyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dauo.mutation = mutation
			node, err = dauo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dauo.hooks) - 1; i >= 0; i-- {
			mut = dauo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dauo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (dauo *DrugAllergyUpdateOne) SaveX(ctx context.Context) *DrugAllergy {
	da, err := dauo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return da
}

// Exec executes the query on the entity.
func (dauo *DrugAllergyUpdateOne) Exec(ctx context.Context) error {
	_, err := dauo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dauo *DrugAllergyUpdateOne) ExecX(ctx context.Context) {
	if err := dauo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (dauo *DrugAllergyUpdateOne) sqlSave(ctx context.Context) (da *DrugAllergy, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   drugallergy.Table,
			Columns: drugallergy.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: drugallergy.FieldID,
			},
		},
	}
	id, ok := dauo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing DrugAllergy.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := dauo.mutation.DateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: drugallergy.FieldDateTime,
		})
	}
	if dauo.mutation.PatientCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PatientTable,
			Columns: []string{drugallergy.PatientColumn},
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
	if nodes := dauo.mutation.PatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PatientTable,
			Columns: []string{drugallergy.PatientColumn},
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
	if dauo.mutation.MedicineCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.MedicineTable,
			Columns: []string{drugallergy.MedicineColumn},
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
	if nodes := dauo.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.MedicineTable,
			Columns: []string{drugallergy.MedicineColumn},
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
	if dauo.mutation.PharmacistCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PharmacistTable,
			Columns: []string{drugallergy.PharmacistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pharmacist.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := dauo.mutation.PharmacistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   drugallergy.PharmacistTable,
			Columns: []string{drugallergy.PharmacistColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: pharmacist.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	da = &DrugAllergy{config: dauo.config}
	_spec.Assign = da.assignValues
	_spec.ScanValues = da.scanValues()
	if err = sqlgraph.UpdateNode(ctx, dauo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{drugallergy.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return da, nil
}