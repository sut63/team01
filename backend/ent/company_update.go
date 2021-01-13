// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/company"
	"github.com/sut63/team01/ent/order"
	"github.com/sut63/team01/ent/predicate"
)

// CompanyUpdate is the builder for updating Company entities.
type CompanyUpdate struct {
	config
	hooks      []Hook
	mutation   *CompanyMutation
	predicates []predicate.Company
}

// Where adds a new predicate for the builder.
func (cu *CompanyUpdate) Where(ps ...predicate.Company) *CompanyUpdate {
	cu.predicates = append(cu.predicates, ps...)
	return cu
}

// SetName sets the Name field.
func (cu *CompanyUpdate) SetName(s string) *CompanyUpdate {
	cu.mutation.SetName(s)
	return cu
}

// AddOrdercompanyIDs adds the ordercompany edge to Order by ids.
func (cu *CompanyUpdate) AddOrdercompanyIDs(ids ...int) *CompanyUpdate {
	cu.mutation.AddOrdercompanyIDs(ids...)
	return cu
}

// AddOrdercompany adds the ordercompany edges to Order.
func (cu *CompanyUpdate) AddOrdercompany(o ...*Order) *CompanyUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.AddOrdercompanyIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cu *CompanyUpdate) Mutation() *CompanyMutation {
	return cu.mutation
}

// RemoveOrdercompanyIDs removes the ordercompany edge to Order by ids.
func (cu *CompanyUpdate) RemoveOrdercompanyIDs(ids ...int) *CompanyUpdate {
	cu.mutation.RemoveOrdercompanyIDs(ids...)
	return cu
}

// RemoveOrdercompany removes ordercompany edges to Order.
func (cu *CompanyUpdate) RemoveOrdercompany(o ...*Order) *CompanyUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cu.RemoveOrdercompanyIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (cu *CompanyUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := cu.mutation.Name(); ok {
		if err := company.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *CompanyUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *CompanyUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *CompanyUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cu *CompanyUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   company.Table,
			Columns: company.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		},
	}
	if ps := cu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldName,
		})
	}
	if nodes := cu.mutation.RemovedOrdercompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.OrdercompanyTable,
			Columns: []string{company.OrdercompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.OrdercompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.OrdercompanyTable,
			Columns: []string{company.OrdercompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// CompanyUpdateOne is the builder for updating a single Company entity.
type CompanyUpdateOne struct {
	config
	hooks    []Hook
	mutation *CompanyMutation
}

// SetName sets the Name field.
func (cuo *CompanyUpdateOne) SetName(s string) *CompanyUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// AddOrdercompanyIDs adds the ordercompany edge to Order by ids.
func (cuo *CompanyUpdateOne) AddOrdercompanyIDs(ids ...int) *CompanyUpdateOne {
	cuo.mutation.AddOrdercompanyIDs(ids...)
	return cuo
}

// AddOrdercompany adds the ordercompany edges to Order.
func (cuo *CompanyUpdateOne) AddOrdercompany(o ...*Order) *CompanyUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.AddOrdercompanyIDs(ids...)
}

// Mutation returns the CompanyMutation object of the builder.
func (cuo *CompanyUpdateOne) Mutation() *CompanyMutation {
	return cuo.mutation
}

// RemoveOrdercompanyIDs removes the ordercompany edge to Order by ids.
func (cuo *CompanyUpdateOne) RemoveOrdercompanyIDs(ids ...int) *CompanyUpdateOne {
	cuo.mutation.RemoveOrdercompanyIDs(ids...)
	return cuo
}

// RemoveOrdercompany removes ordercompany edges to Order.
func (cuo *CompanyUpdateOne) RemoveOrdercompany(o ...*Order) *CompanyUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return cuo.RemoveOrdercompanyIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (cuo *CompanyUpdateOne) Save(ctx context.Context) (*Company, error) {
	if v, ok := cuo.mutation.Name(); ok {
		if err := company.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "Name", err: fmt.Errorf("ent: validator failed for field \"Name\": %w", err)}
		}
	}

	var (
		err  error
		node *Company
	)
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*CompanyMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *CompanyUpdateOne) SaveX(ctx context.Context) *Company {
	c, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return c
}

// Exec executes the query on the entity.
func (cuo *CompanyUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *CompanyUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (cuo *CompanyUpdateOne) sqlSave(ctx context.Context) (c *Company, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   company.Table,
			Columns: company.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: company.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Company.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: company.FieldName,
		})
	}
	if nodes := cuo.mutation.RemovedOrdercompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.OrdercompanyTable,
			Columns: []string{company.OrdercompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.OrdercompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   company.OrdercompanyTable,
			Columns: []string{company.OrdercompanyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: order.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	c = &Company{config: cuo.config}
	_spec.Assign = c.assignValues
	_spec.ScanValues = c.scanValues()
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{company.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return c, nil
}
