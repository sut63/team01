// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/annotation"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/predicate"
)

// AnnotationUpdate is the builder for updating Annotation entities.
type AnnotationUpdate struct {
	config
	hooks      []Hook
	mutation   *AnnotationMutation
	predicates []predicate.Annotation
}

// Where adds a new predicate for the builder.
func (au *AnnotationUpdate) Where(ps ...predicate.Annotation) *AnnotationUpdate {
	au.predicates = append(au.predicates, ps...)
	return au
}

// SetMessages sets the messages field.
func (au *AnnotationUpdate) SetMessages(s string) *AnnotationUpdate {
	au.mutation.SetMessages(s)
	return au
}

// AddDispensemedicineIDs adds the dispensemedicine edge to DispenseMedicine by ids.
func (au *AnnotationUpdate) AddDispensemedicineIDs(ids ...int) *AnnotationUpdate {
	au.mutation.AddDispensemedicineIDs(ids...)
	return au
}

// AddDispensemedicine adds the dispensemedicine edges to DispenseMedicine.
func (au *AnnotationUpdate) AddDispensemedicine(d ...*DispenseMedicine) *AnnotationUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return au.AddDispensemedicineIDs(ids...)
}

// Mutation returns the AnnotationMutation object of the builder.
func (au *AnnotationUpdate) Mutation() *AnnotationMutation {
	return au.mutation
}

// RemoveDispensemedicineIDs removes the dispensemedicine edge to DispenseMedicine by ids.
func (au *AnnotationUpdate) RemoveDispensemedicineIDs(ids ...int) *AnnotationUpdate {
	au.mutation.RemoveDispensemedicineIDs(ids...)
	return au
}

// RemoveDispensemedicine removes dispensemedicine edges to DispenseMedicine.
func (au *AnnotationUpdate) RemoveDispensemedicine(d ...*DispenseMedicine) *AnnotationUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return au.RemoveDispensemedicineIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (au *AnnotationUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := au.mutation.Messages(); ok {
		if err := annotation.MessagesValidator(v); err != nil {
			return 0, &ValidationError{Name: "messages", err: fmt.Errorf("ent: validator failed for field \"messages\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(au.hooks) == 0 {
		affected, err = au.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnotationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			au.mutation = mutation
			affected, err = au.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(au.hooks) - 1; i >= 0; i-- {
			mut = au.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, au.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (au *AnnotationUpdate) SaveX(ctx context.Context) int {
	affected, err := au.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (au *AnnotationUpdate) Exec(ctx context.Context) error {
	_, err := au.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (au *AnnotationUpdate) ExecX(ctx context.Context) {
	if err := au.Exec(ctx); err != nil {
		panic(err)
	}
}

func (au *AnnotationUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   annotation.Table,
			Columns: annotation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: annotation.FieldID,
			},
		},
	}
	if ps := au.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := au.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: annotation.FieldMessages,
		})
	}
	if nodes := au.mutation.RemovedDispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   annotation.DispensemedicineTable,
			Columns: []string{annotation.DispensemedicineColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := au.mutation.DispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   annotation.DispensemedicineTable,
			Columns: []string{annotation.DispensemedicineColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, au.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{annotation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// AnnotationUpdateOne is the builder for updating a single Annotation entity.
type AnnotationUpdateOne struct {
	config
	hooks    []Hook
	mutation *AnnotationMutation
}

// SetMessages sets the messages field.
func (auo *AnnotationUpdateOne) SetMessages(s string) *AnnotationUpdateOne {
	auo.mutation.SetMessages(s)
	return auo
}

// AddDispensemedicineIDs adds the dispensemedicine edge to DispenseMedicine by ids.
func (auo *AnnotationUpdateOne) AddDispensemedicineIDs(ids ...int) *AnnotationUpdateOne {
	auo.mutation.AddDispensemedicineIDs(ids...)
	return auo
}

// AddDispensemedicine adds the dispensemedicine edges to DispenseMedicine.
func (auo *AnnotationUpdateOne) AddDispensemedicine(d ...*DispenseMedicine) *AnnotationUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return auo.AddDispensemedicineIDs(ids...)
}

// Mutation returns the AnnotationMutation object of the builder.
func (auo *AnnotationUpdateOne) Mutation() *AnnotationMutation {
	return auo.mutation
}

// RemoveDispensemedicineIDs removes the dispensemedicine edge to DispenseMedicine by ids.
func (auo *AnnotationUpdateOne) RemoveDispensemedicineIDs(ids ...int) *AnnotationUpdateOne {
	auo.mutation.RemoveDispensemedicineIDs(ids...)
	return auo
}

// RemoveDispensemedicine removes dispensemedicine edges to DispenseMedicine.
func (auo *AnnotationUpdateOne) RemoveDispensemedicine(d ...*DispenseMedicine) *AnnotationUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return auo.RemoveDispensemedicineIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (auo *AnnotationUpdateOne) Save(ctx context.Context) (*Annotation, error) {
	if v, ok := auo.mutation.Messages(); ok {
		if err := annotation.MessagesValidator(v); err != nil {
			return nil, &ValidationError{Name: "messages", err: fmt.Errorf("ent: validator failed for field \"messages\": %w", err)}
		}
	}

	var (
		err  error
		node *Annotation
	)
	if len(auo.hooks) == 0 {
		node, err = auo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AnnotationMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			auo.mutation = mutation
			node, err = auo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(auo.hooks) - 1; i >= 0; i-- {
			mut = auo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (auo *AnnotationUpdateOne) SaveX(ctx context.Context) *Annotation {
	a, err := auo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return a
}

// Exec executes the query on the entity.
func (auo *AnnotationUpdateOne) Exec(ctx context.Context) error {
	_, err := auo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auo *AnnotationUpdateOne) ExecX(ctx context.Context) {
	if err := auo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (auo *AnnotationUpdateOne) sqlSave(ctx context.Context) (a *Annotation, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   annotation.Table,
			Columns: annotation.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: annotation.FieldID,
			},
		},
	}
	id, ok := auo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Annotation.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := auo.mutation.Messages(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: annotation.FieldMessages,
		})
	}
	if nodes := auo.mutation.RemovedDispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   annotation.DispensemedicineTable,
			Columns: []string{annotation.DispensemedicineColumn},
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := auo.mutation.DispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   annotation.DispensemedicineTable,
			Columns: []string{annotation.DispensemedicineColumn},
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
	a = &Annotation{config: auo.config}
	_spec.Assign = a.assignValues
	_spec.ScanValues = a.scanValues()
	if err = sqlgraph.UpdateNode(ctx, auo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{annotation.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return a, nil
}