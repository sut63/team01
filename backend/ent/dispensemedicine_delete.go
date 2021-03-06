// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/predicate"
)

// DispenseMedicineDelete is the builder for deleting a DispenseMedicine entity.
type DispenseMedicineDelete struct {
	config
	hooks      []Hook
	mutation   *DispenseMedicineMutation
	predicates []predicate.DispenseMedicine
}

// Where adds a new predicate to the delete builder.
func (dmd *DispenseMedicineDelete) Where(ps ...predicate.DispenseMedicine) *DispenseMedicineDelete {
	dmd.predicates = append(dmd.predicates, ps...)
	return dmd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (dmd *DispenseMedicineDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(dmd.hooks) == 0 {
		affected, err = dmd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DispenseMedicineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dmd.mutation = mutation
			affected, err = dmd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(dmd.hooks) - 1; i >= 0; i-- {
			mut = dmd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dmd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (dmd *DispenseMedicineDelete) ExecX(ctx context.Context) int {
	n, err := dmd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (dmd *DispenseMedicineDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: dispensemedicine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: dispensemedicine.FieldID,
			},
		},
	}
	if ps := dmd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, dmd.driver, _spec)
}

// DispenseMedicineDeleteOne is the builder for deleting a single DispenseMedicine entity.
type DispenseMedicineDeleteOne struct {
	dmd *DispenseMedicineDelete
}

// Exec executes the deletion query.
func (dmdo *DispenseMedicineDeleteOne) Exec(ctx context.Context) error {
	n, err := dmdo.dmd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{dispensemedicine.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (dmdo *DispenseMedicineDeleteOne) ExecX(ctx context.Context) {
	dmdo.dmd.ExecX(ctx)
}
