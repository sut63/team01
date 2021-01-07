// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/levelofdangerous"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/predicate"
)

// LevelOfDangerousUpdate is the builder for updating LevelOfDangerous entities.
type LevelOfDangerousUpdate struct {
	config
	hooks      []Hook
	mutation   *LevelOfDangerousMutation
	predicates []predicate.LevelOfDangerous
}

// Where adds a new predicate for the builder.
func (lodu *LevelOfDangerousUpdate) Where(ps ...predicate.LevelOfDangerous) *LevelOfDangerousUpdate {
	lodu.predicates = append(lodu.predicates, ps...)
	return lodu
}

// SetName sets the name field.
func (lodu *LevelOfDangerousUpdate) SetName(s string) *LevelOfDangerousUpdate {
	lodu.mutation.SetName(s)
	return lodu
}

// AddMedicineIDs adds the Medicine edge to Medicine by ids.
func (lodu *LevelOfDangerousUpdate) AddMedicineIDs(ids ...int) *LevelOfDangerousUpdate {
	lodu.mutation.AddMedicineIDs(ids...)
	return lodu
}

// AddMedicine adds the Medicine edges to Medicine.
func (lodu *LevelOfDangerousUpdate) AddMedicine(m ...*Medicine) *LevelOfDangerousUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return lodu.AddMedicineIDs(ids...)
}

// Mutation returns the LevelOfDangerousMutation object of the builder.
func (lodu *LevelOfDangerousUpdate) Mutation() *LevelOfDangerousMutation {
	return lodu.mutation
}

// RemoveMedicineIDs removes the Medicine edge to Medicine by ids.
func (lodu *LevelOfDangerousUpdate) RemoveMedicineIDs(ids ...int) *LevelOfDangerousUpdate {
	lodu.mutation.RemoveMedicineIDs(ids...)
	return lodu
}

// RemoveMedicine removes Medicine edges to Medicine.
func (lodu *LevelOfDangerousUpdate) RemoveMedicine(m ...*Medicine) *LevelOfDangerousUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return lodu.RemoveMedicineIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (lodu *LevelOfDangerousUpdate) Save(ctx context.Context) (int, error) {

	var (
		err      error
		affected int
	)
	if len(lodu.hooks) == 0 {
		affected, err = lodu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LevelOfDangerousMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lodu.mutation = mutation
			affected, err = lodu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lodu.hooks) - 1; i >= 0; i-- {
			mut = lodu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lodu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lodu *LevelOfDangerousUpdate) SaveX(ctx context.Context) int {
	affected, err := lodu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lodu *LevelOfDangerousUpdate) Exec(ctx context.Context) error {
	_, err := lodu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lodu *LevelOfDangerousUpdate) ExecX(ctx context.Context) {
	if err := lodu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (lodu *LevelOfDangerousUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   levelofdangerous.Table,
			Columns: levelofdangerous.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: levelofdangerous.FieldID,
			},
		},
	}
	if ps := lodu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lodu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: levelofdangerous.FieldName,
		})
	}
	if nodes := lodu.mutation.RemovedMedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   levelofdangerous.MedicineTable,
			Columns: []string{levelofdangerous.MedicineColumn},
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
	if nodes := lodu.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   levelofdangerous.MedicineTable,
			Columns: []string{levelofdangerous.MedicineColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, lodu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{levelofdangerous.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// LevelOfDangerousUpdateOne is the builder for updating a single LevelOfDangerous entity.
type LevelOfDangerousUpdateOne struct {
	config
	hooks    []Hook
	mutation *LevelOfDangerousMutation
}

// SetName sets the name field.
func (loduo *LevelOfDangerousUpdateOne) SetName(s string) *LevelOfDangerousUpdateOne {
	loduo.mutation.SetName(s)
	return loduo
}

// AddMedicineIDs adds the Medicine edge to Medicine by ids.
func (loduo *LevelOfDangerousUpdateOne) AddMedicineIDs(ids ...int) *LevelOfDangerousUpdateOne {
	loduo.mutation.AddMedicineIDs(ids...)
	return loduo
}

// AddMedicine adds the Medicine edges to Medicine.
func (loduo *LevelOfDangerousUpdateOne) AddMedicine(m ...*Medicine) *LevelOfDangerousUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return loduo.AddMedicineIDs(ids...)
}

// Mutation returns the LevelOfDangerousMutation object of the builder.
func (loduo *LevelOfDangerousUpdateOne) Mutation() *LevelOfDangerousMutation {
	return loduo.mutation
}

// RemoveMedicineIDs removes the Medicine edge to Medicine by ids.
func (loduo *LevelOfDangerousUpdateOne) RemoveMedicineIDs(ids ...int) *LevelOfDangerousUpdateOne {
	loduo.mutation.RemoveMedicineIDs(ids...)
	return loduo
}

// RemoveMedicine removes Medicine edges to Medicine.
func (loduo *LevelOfDangerousUpdateOne) RemoveMedicine(m ...*Medicine) *LevelOfDangerousUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return loduo.RemoveMedicineIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (loduo *LevelOfDangerousUpdateOne) Save(ctx context.Context) (*LevelOfDangerous, error) {

	var (
		err  error
		node *LevelOfDangerous
	)
	if len(loduo.hooks) == 0 {
		node, err = loduo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*LevelOfDangerousMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			loduo.mutation = mutation
			node, err = loduo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(loduo.hooks) - 1; i >= 0; i-- {
			mut = loduo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, loduo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (loduo *LevelOfDangerousUpdateOne) SaveX(ctx context.Context) *LevelOfDangerous {
	lod, err := loduo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return lod
}

// Exec executes the query on the entity.
func (loduo *LevelOfDangerousUpdateOne) Exec(ctx context.Context) error {
	_, err := loduo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (loduo *LevelOfDangerousUpdateOne) ExecX(ctx context.Context) {
	if err := loduo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (loduo *LevelOfDangerousUpdateOne) sqlSave(ctx context.Context) (lod *LevelOfDangerous, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   levelofdangerous.Table,
			Columns: levelofdangerous.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: levelofdangerous.FieldID,
			},
		},
	}
	id, ok := loduo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing LevelOfDangerous.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := loduo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: levelofdangerous.FieldName,
		})
	}
	if nodes := loduo.mutation.RemovedMedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   levelofdangerous.MedicineTable,
			Columns: []string{levelofdangerous.MedicineColumn},
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
	if nodes := loduo.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   levelofdangerous.MedicineTable,
			Columns: []string{levelofdangerous.MedicineColumn},
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
	lod = &LevelOfDangerous{config: loduo.config}
	_spec.Assign = lod.assignValues
	_spec.ScanValues = lod.scanValues()
	if err = sqlgraph.UpdateNode(ctx, loduo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{levelofdangerous.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return lod, nil
}