// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/predicate"
	"github.com/sut63/team01/ent/unitofmedicine"
)

// UnitOfMedicineUpdate is the builder for updating UnitOfMedicine entities.
type UnitOfMedicineUpdate struct {
	config
	hooks      []Hook
	mutation   *UnitOfMedicineMutation
	predicates []predicate.UnitOfMedicine
}

// Where adds a new predicate for the builder.
func (uomu *UnitOfMedicineUpdate) Where(ps ...predicate.UnitOfMedicine) *UnitOfMedicineUpdate {
	uomu.predicates = append(uomu.predicates, ps...)
	return uomu
}

// SetName sets the name field.
func (uomu *UnitOfMedicineUpdate) SetName(s string) *UnitOfMedicineUpdate {
	uomu.mutation.SetName(s)
	return uomu
}

// AddMedicineIDs adds the Medicine edge to Medicine by ids.
func (uomu *UnitOfMedicineUpdate) AddMedicineIDs(ids ...int) *UnitOfMedicineUpdate {
	uomu.mutation.AddMedicineIDs(ids...)
	return uomu
}

// AddMedicine adds the Medicine edges to Medicine.
func (uomu *UnitOfMedicineUpdate) AddMedicine(m ...*Medicine) *UnitOfMedicineUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uomu.AddMedicineIDs(ids...)
}

// Mutation returns the UnitOfMedicineMutation object of the builder.
func (uomu *UnitOfMedicineUpdate) Mutation() *UnitOfMedicineMutation {
	return uomu.mutation
}

// RemoveMedicineIDs removes the Medicine edge to Medicine by ids.
func (uomu *UnitOfMedicineUpdate) RemoveMedicineIDs(ids ...int) *UnitOfMedicineUpdate {
	uomu.mutation.RemoveMedicineIDs(ids...)
	return uomu
}

// RemoveMedicine removes Medicine edges to Medicine.
func (uomu *UnitOfMedicineUpdate) RemoveMedicine(m ...*Medicine) *UnitOfMedicineUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uomu.RemoveMedicineIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (uomu *UnitOfMedicineUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(uomu.hooks) == 0 {
		affected, err = uomu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnitOfMedicineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uomu.mutation = mutation
			affected, err = uomu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(uomu.hooks) - 1; i >= 0; i-- {
			mut = uomu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uomu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (uomu *UnitOfMedicineUpdate) SaveX(ctx context.Context) int {
	affected, err := uomu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (uomu *UnitOfMedicineUpdate) Exec(ctx context.Context) error {
	_, err := uomu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uomu *UnitOfMedicineUpdate) ExecX(ctx context.Context) {
	if err := uomu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uomu *UnitOfMedicineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unitofmedicine.Table,
			Columns: unitofmedicine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unitofmedicine.FieldID,
			},
		},
	}
	if ps := uomu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := uomu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unitofmedicine.FieldName,
		})
	}
	if nodes := uomu.mutation.RemovedMedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   unitofmedicine.MedicineTable,
			Columns: []string{unitofmedicine.MedicineColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uomu.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   unitofmedicine.MedicineTable,
			Columns: []string{unitofmedicine.MedicineColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, uomu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unitofmedicine.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// UnitOfMedicineUpdateOne is the builder for updating a single UnitOfMedicine entity.
type UnitOfMedicineUpdateOne struct {
	config
	hooks    []Hook
	mutation *UnitOfMedicineMutation
}

// SetName sets the name field.
func (uomuo *UnitOfMedicineUpdateOne) SetName(s string) *UnitOfMedicineUpdateOne {
	uomuo.mutation.SetName(s)
	return uomuo
}

// AddMedicineIDs adds the Medicine edge to Medicine by ids.
func (uomuo *UnitOfMedicineUpdateOne) AddMedicineIDs(ids ...int) *UnitOfMedicineUpdateOne {
	uomuo.mutation.AddMedicineIDs(ids...)
	return uomuo
}

// AddMedicine adds the Medicine edges to Medicine.
func (uomuo *UnitOfMedicineUpdateOne) AddMedicine(m ...*Medicine) *UnitOfMedicineUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uomuo.AddMedicineIDs(ids...)
}

// Mutation returns the UnitOfMedicineMutation object of the builder.
func (uomuo *UnitOfMedicineUpdateOne) Mutation() *UnitOfMedicineMutation {
	return uomuo.mutation
}

// RemoveMedicineIDs removes the Medicine edge to Medicine by ids.
func (uomuo *UnitOfMedicineUpdateOne) RemoveMedicineIDs(ids ...int) *UnitOfMedicineUpdateOne {
	uomuo.mutation.RemoveMedicineIDs(ids...)
	return uomuo
}

// RemoveMedicine removes Medicine edges to Medicine.
func (uomuo *UnitOfMedicineUpdateOne) RemoveMedicine(m ...*Medicine) *UnitOfMedicineUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return uomuo.RemoveMedicineIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (uomuo *UnitOfMedicineUpdateOne) Save(ctx context.Context) (*UnitOfMedicine, error) {

	var (
		err  error
		node *UnitOfMedicine
	)
	if len(uomuo.hooks) == 0 {
		node, err = uomuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UnitOfMedicineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			uomuo.mutation = mutation
			node, err = uomuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(uomuo.hooks) - 1; i >= 0; i-- {
			mut = uomuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, uomuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (uomuo *UnitOfMedicineUpdateOne) SaveX(ctx context.Context) *UnitOfMedicine {
	uom, err := uomuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return uom
}

// Exec executes the query on the entity.
func (uomuo *UnitOfMedicineUpdateOne) Exec(ctx context.Context) error {
	_, err := uomuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uomuo *UnitOfMedicineUpdateOne) ExecX(ctx context.Context) {
	if err := uomuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (uomuo *UnitOfMedicineUpdateOne) sqlSave(ctx context.Context) (uom *UnitOfMedicine, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   unitofmedicine.Table,
			Columns: unitofmedicine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: unitofmedicine.FieldID,
			},
		},
	}
	id, ok := uomuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UnitOfMedicine.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := uomuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: unitofmedicine.FieldName,
		})
	}
	if nodes := uomuo.mutation.RemovedMedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   unitofmedicine.MedicineTable,
			Columns: []string{unitofmedicine.MedicineColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := uomuo.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   unitofmedicine.MedicineTable,
			Columns: []string{unitofmedicine.MedicineColumn},
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
	uom = &UnitOfMedicine{config: uomuo.config}
	_spec.Assign = uom.assignValues
	_spec.ScanValues = uom.scanValues()
	if err = sqlgraph.UpdateNode(ctx, uomuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{unitofmedicine.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return uom, nil
}
