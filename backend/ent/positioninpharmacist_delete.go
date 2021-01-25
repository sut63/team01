// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/positioninpharmacist"
	"github.com/sut63/team01/ent/predicate"
)

// PositionInPharmacistDelete is the builder for deleting a PositionInPharmacist entity.
type PositionInPharmacistDelete struct {
	config
	hooks      []Hook
	mutation   *PositionInPharmacistMutation
	predicates []predicate.PositionInPharmacist
}

// Where adds a new predicate to the delete builder.
func (pipd *PositionInPharmacistDelete) Where(ps ...predicate.PositionInPharmacist) *PositionInPharmacistDelete {
	pipd.predicates = append(pipd.predicates, ps...)
	return pipd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (pipd *PositionInPharmacistDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pipd.hooks) == 0 {
		affected, err = pipd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PositionInPharmacistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pipd.mutation = mutation
			affected, err = pipd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pipd.hooks) - 1; i >= 0; i-- {
			mut = pipd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pipd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (pipd *PositionInPharmacistDelete) ExecX(ctx context.Context) int {
	n, err := pipd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (pipd *PositionInPharmacistDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: positioninpharmacist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: positioninpharmacist.FieldID,
			},
		},
	}
	if ps := pipd.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return sqlgraph.DeleteNodes(ctx, pipd.driver, _spec)
}

// PositionInPharmacistDeleteOne is the builder for deleting a single PositionInPharmacist entity.
type PositionInPharmacistDeleteOne struct {
	pipd *PositionInPharmacistDelete
}

// Exec executes the deletion query.
func (pipdo *PositionInPharmacistDeleteOne) Exec(ctx context.Context) error {
	n, err := pipdo.pipd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{positioninpharmacist.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (pipdo *PositionInPharmacistDeleteOne) ExecX(ctx context.Context) {
	pipdo.pipd.ExecX(ctx)
}