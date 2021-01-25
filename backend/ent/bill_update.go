// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/bill"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/payment"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/predicate"
)

// BillUpdate is the builder for updating Bill entities.
type BillUpdate struct {
	config
	hooks      []Hook
	mutation   *BillMutation
	predicates []predicate.Bill
}

// Where adds a new predicate for the builder.
func (bu *BillUpdate) Where(ps ...predicate.Bill) *BillUpdate {
	bu.predicates = append(bu.predicates, ps...)
	return bu
}

// SetAmount sets the amount field.
func (bu *BillUpdate) SetAmount(i int) *BillUpdate {
	bu.mutation.ResetAmount()
	bu.mutation.SetAmount(i)
	return bu
}

// AddAmount adds i to amount.
func (bu *BillUpdate) AddAmount(i int) *BillUpdate {
	bu.mutation.AddAmount(i)
	return bu
}

// SetAnnotation sets the annotation field.
func (bu *BillUpdate) SetAnnotation(s string) *BillUpdate {
	bu.mutation.SetAnnotation(s)
	return bu
}

// SetPayer sets the payer field.
func (bu *BillUpdate) SetPayer(s string) *BillUpdate {
	bu.mutation.SetPayer(s)
	return bu
}

// SetPharmacistsID sets the pharmacists edge to Pharmacist by id.
func (bu *BillUpdate) SetPharmacistsID(id int) *BillUpdate {
	bu.mutation.SetPharmacistsID(id)
	return bu
}

// SetNillablePharmacistsID sets the pharmacists edge to Pharmacist by id if the given value is not nil.
func (bu *BillUpdate) SetNillablePharmacistsID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetPharmacistsID(*id)
	}
	return bu
}

// SetPharmacists sets the pharmacists edge to Pharmacist.
func (bu *BillUpdate) SetPharmacists(p *Pharmacist) *BillUpdate {
	return bu.SetPharmacistsID(p.ID)
}

// SetPaymentsID sets the payments edge to Payment by id.
func (bu *BillUpdate) SetPaymentsID(id int) *BillUpdate {
	bu.mutation.SetPaymentsID(id)
	return bu
}

// SetNillablePaymentsID sets the payments edge to Payment by id if the given value is not nil.
func (bu *BillUpdate) SetNillablePaymentsID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetPaymentsID(*id)
	}
	return bu
}

// SetPayments sets the payments edge to Payment.
func (bu *BillUpdate) SetPayments(p *Payment) *BillUpdate {
	return bu.SetPaymentsID(p.ID)
}

// SetDispenseMedicinesID sets the dispenseMedicines edge to DispenseMedicine by id.
func (bu *BillUpdate) SetDispenseMedicinesID(id int) *BillUpdate {
	bu.mutation.SetDispenseMedicinesID(id)
	return bu
}

// SetNillableDispenseMedicinesID sets the dispenseMedicines edge to DispenseMedicine by id if the given value is not nil.
func (bu *BillUpdate) SetNillableDispenseMedicinesID(id *int) *BillUpdate {
	if id != nil {
		bu = bu.SetDispenseMedicinesID(*id)
	}
	return bu
}

// SetDispenseMedicines sets the dispenseMedicines edge to DispenseMedicine.
func (bu *BillUpdate) SetDispenseMedicines(d *DispenseMedicine) *BillUpdate {
	return bu.SetDispenseMedicinesID(d.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bu *BillUpdate) Mutation() *BillMutation {
	return bu.mutation
}

// ClearPharmacists clears the pharmacists edge to Pharmacist.
func (bu *BillUpdate) ClearPharmacists() *BillUpdate {
	bu.mutation.ClearPharmacists()
	return bu
}

// ClearPayments clears the payments edge to Payment.
func (bu *BillUpdate) ClearPayments() *BillUpdate {
	bu.mutation.ClearPayments()
	return bu
}

// ClearDispenseMedicines clears the dispenseMedicines edge to DispenseMedicine.
func (bu *BillUpdate) ClearDispenseMedicines() *BillUpdate {
	bu.mutation.ClearDispenseMedicines()
	return bu
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (bu *BillUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := bu.mutation.Amount(); ok {
		if err := bill.AmountValidator(v); err != nil {
			return 0, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if v, ok := bu.mutation.Annotation(); ok {
		if err := bill.AnnotationValidator(v); err != nil {
			return 0, &ValidationError{Name: "annotation", err: fmt.Errorf("ent: validator failed for field \"annotation\": %w", err)}
		}
	}
	if v, ok := bu.mutation.Payer(); ok {
		if err := bill.PayerValidator(v); err != nil {
			return 0, &ValidationError{Name: "payer", err: fmt.Errorf("ent: validator failed for field \"payer\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(bu.hooks) == 0 {
		affected, err = bu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bu.mutation = mutation
			affected, err = bu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(bu.hooks) - 1; i >= 0; i-- {
			mut = bu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (bu *BillUpdate) SaveX(ctx context.Context) int {
	affected, err := bu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (bu *BillUpdate) Exec(ctx context.Context) error {
	_, err := bu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (bu *BillUpdate) ExecX(ctx context.Context) {
	if err := bu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (bu *BillUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bill.Table,
			Columns: bill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bill.FieldID,
			},
		},
	}
	if ps := bu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := bu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldAmount,
		})
	}
	if value, ok := bu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldAmount,
		})
	}
	if value, ok := bu.mutation.Annotation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldAnnotation,
		})
	}
	if value, ok := bu.mutation.Payer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayer,
		})
	}
	if bu.mutation.PharmacistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PharmacistsTable,
			Columns: []string{bill.PharmacistsColumn},
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
	if nodes := bu.mutation.PharmacistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PharmacistsTable,
			Columns: []string{bill.PharmacistsColumn},
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
	if bu.mutation.PaymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PaymentsTable,
			Columns: []string{bill.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := bu.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PaymentsTable,
			Columns: []string{bill.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if bu.mutation.DispenseMedicinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DispenseMedicinesTable,
			Columns: []string{bill.DispenseMedicinesColumn},
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
	if nodes := bu.mutation.DispenseMedicinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DispenseMedicinesTable,
			Columns: []string{bill.DispenseMedicinesColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, bu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// BillUpdateOne is the builder for updating a single Bill entity.
type BillUpdateOne struct {
	config
	hooks    []Hook
	mutation *BillMutation
}

// SetAmount sets the amount field.
func (buo *BillUpdateOne) SetAmount(i int) *BillUpdateOne {
	buo.mutation.ResetAmount()
	buo.mutation.SetAmount(i)
	return buo
}

// AddAmount adds i to amount.
func (buo *BillUpdateOne) AddAmount(i int) *BillUpdateOne {
	buo.mutation.AddAmount(i)
	return buo
}

// SetAnnotation sets the annotation field.
func (buo *BillUpdateOne) SetAnnotation(s string) *BillUpdateOne {
	buo.mutation.SetAnnotation(s)
	return buo
}

// SetPayer sets the payer field.
func (buo *BillUpdateOne) SetPayer(s string) *BillUpdateOne {
	buo.mutation.SetPayer(s)
	return buo
}

// SetPharmacistsID sets the pharmacists edge to Pharmacist by id.
func (buo *BillUpdateOne) SetPharmacistsID(id int) *BillUpdateOne {
	buo.mutation.SetPharmacistsID(id)
	return buo
}

// SetNillablePharmacistsID sets the pharmacists edge to Pharmacist by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillablePharmacistsID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetPharmacistsID(*id)
	}
	return buo
}

// SetPharmacists sets the pharmacists edge to Pharmacist.
func (buo *BillUpdateOne) SetPharmacists(p *Pharmacist) *BillUpdateOne {
	return buo.SetPharmacistsID(p.ID)
}

// SetPaymentsID sets the payments edge to Payment by id.
func (buo *BillUpdateOne) SetPaymentsID(id int) *BillUpdateOne {
	buo.mutation.SetPaymentsID(id)
	return buo
}

// SetNillablePaymentsID sets the payments edge to Payment by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillablePaymentsID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetPaymentsID(*id)
	}
	return buo
}

// SetPayments sets the payments edge to Payment.
func (buo *BillUpdateOne) SetPayments(p *Payment) *BillUpdateOne {
	return buo.SetPaymentsID(p.ID)
}

// SetDispenseMedicinesID sets the dispenseMedicines edge to DispenseMedicine by id.
func (buo *BillUpdateOne) SetDispenseMedicinesID(id int) *BillUpdateOne {
	buo.mutation.SetDispenseMedicinesID(id)
	return buo
}

// SetNillableDispenseMedicinesID sets the dispenseMedicines edge to DispenseMedicine by id if the given value is not nil.
func (buo *BillUpdateOne) SetNillableDispenseMedicinesID(id *int) *BillUpdateOne {
	if id != nil {
		buo = buo.SetDispenseMedicinesID(*id)
	}
	return buo
}

// SetDispenseMedicines sets the dispenseMedicines edge to DispenseMedicine.
func (buo *BillUpdateOne) SetDispenseMedicines(d *DispenseMedicine) *BillUpdateOne {
	return buo.SetDispenseMedicinesID(d.ID)
}

// Mutation returns the BillMutation object of the builder.
func (buo *BillUpdateOne) Mutation() *BillMutation {
	return buo.mutation
}

// ClearPharmacists clears the pharmacists edge to Pharmacist.
func (buo *BillUpdateOne) ClearPharmacists() *BillUpdateOne {
	buo.mutation.ClearPharmacists()
	return buo
}

// ClearPayments clears the payments edge to Payment.
func (buo *BillUpdateOne) ClearPayments() *BillUpdateOne {
	buo.mutation.ClearPayments()
	return buo
}

// ClearDispenseMedicines clears the dispenseMedicines edge to DispenseMedicine.
func (buo *BillUpdateOne) ClearDispenseMedicines() *BillUpdateOne {
	buo.mutation.ClearDispenseMedicines()
	return buo
}

// Save executes the query and returns the updated entity.
func (buo *BillUpdateOne) Save(ctx context.Context) (*Bill, error) {
	if v, ok := buo.mutation.Amount(); ok {
		if err := bill.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if v, ok := buo.mutation.Annotation(); ok {
		if err := bill.AnnotationValidator(v); err != nil {
			return nil, &ValidationError{Name: "annotation", err: fmt.Errorf("ent: validator failed for field \"annotation\": %w", err)}
		}
	}
	if v, ok := buo.mutation.Payer(); ok {
		if err := bill.PayerValidator(v); err != nil {
			return nil, &ValidationError{Name: "payer", err: fmt.Errorf("ent: validator failed for field \"payer\": %w", err)}
		}
	}

	var (
		err  error
		node *Bill
	)
	if len(buo.hooks) == 0 {
		node, err = buo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			buo.mutation = mutation
			node, err = buo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(buo.hooks) - 1; i >= 0; i-- {
			mut = buo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, buo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (buo *BillUpdateOne) SaveX(ctx context.Context) *Bill {
	b, err := buo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return b
}

// Exec executes the query on the entity.
func (buo *BillUpdateOne) Exec(ctx context.Context) error {
	_, err := buo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (buo *BillUpdateOne) ExecX(ctx context.Context) {
	if err := buo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (buo *BillUpdateOne) sqlSave(ctx context.Context) (b *Bill, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   bill.Table,
			Columns: bill.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bill.FieldID,
			},
		},
	}
	id, ok := buo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Bill.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := buo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldAmount,
		})
	}
	if value, ok := buo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldAmount,
		})
	}
	if value, ok := buo.mutation.Annotation(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldAnnotation,
		})
	}
	if value, ok := buo.mutation.Payer(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayer,
		})
	}
	if buo.mutation.PharmacistsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PharmacistsTable,
			Columns: []string{bill.PharmacistsColumn},
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
	if nodes := buo.mutation.PharmacistsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PharmacistsTable,
			Columns: []string{bill.PharmacistsColumn},
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
	if buo.mutation.PaymentsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PaymentsTable,
			Columns: []string{bill.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := buo.mutation.PaymentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   bill.PaymentsTable,
			Columns: []string{bill.PaymentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: payment.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if buo.mutation.DispenseMedicinesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DispenseMedicinesTable,
			Columns: []string{bill.DispenseMedicinesColumn},
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
	if nodes := buo.mutation.DispenseMedicinesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: true,
			Table:   bill.DispenseMedicinesTable,
			Columns: []string{bill.DispenseMedicinesColumn},
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
	b = &Bill{config: buo.config}
	_spec.Assign = b.assignValues
	_spec.ScanValues = b.scanValues()
	if err = sqlgraph.UpdateNode(ctx, buo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{bill.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return b, nil
}
