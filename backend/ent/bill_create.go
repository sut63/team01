// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/bill"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/payment"
	"github.com/sut63/team01/ent/pharmacist"
)

// BillCreate is the builder for creating a Bill entity.
type BillCreate struct {
	config
	mutation *BillMutation
	hooks    []Hook
}

// SetAmount sets the amount field.
func (bc *BillCreate) SetAmount(i int) *BillCreate {
	bc.mutation.SetAmount(i)
	return bc
}

// SetAnnotation sets the annotation field.
func (bc *BillCreate) SetAnnotation(s string) *BillCreate {
	bc.mutation.SetAnnotation(s)
	return bc
}

// SetPayer sets the payer field.
func (bc *BillCreate) SetPayer(s string) *BillCreate {
	bc.mutation.SetPayer(s)
	return bc
}

// SetPharmacistsID sets the pharmacists edge to Pharmacist by id.
func (bc *BillCreate) SetPharmacistsID(id int) *BillCreate {
	bc.mutation.SetPharmacistsID(id)
	return bc
}

// SetNillablePharmacistsID sets the pharmacists edge to Pharmacist by id if the given value is not nil.
func (bc *BillCreate) SetNillablePharmacistsID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetPharmacistsID(*id)
	}
	return bc
}

// SetPharmacists sets the pharmacists edge to Pharmacist.
func (bc *BillCreate) SetPharmacists(p *Pharmacist) *BillCreate {
	return bc.SetPharmacistsID(p.ID)
}

// SetPaymentsID sets the payments edge to Payment by id.
func (bc *BillCreate) SetPaymentsID(id int) *BillCreate {
	bc.mutation.SetPaymentsID(id)
	return bc
}

// SetNillablePaymentsID sets the payments edge to Payment by id if the given value is not nil.
func (bc *BillCreate) SetNillablePaymentsID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetPaymentsID(*id)
	}
	return bc
}

// SetPayments sets the payments edge to Payment.
func (bc *BillCreate) SetPayments(p *Payment) *BillCreate {
	return bc.SetPaymentsID(p.ID)
}

// SetDispenseMedicinesID sets the dispenseMedicines edge to DispenseMedicine by id.
func (bc *BillCreate) SetDispenseMedicinesID(id int) *BillCreate {
	bc.mutation.SetDispenseMedicinesID(id)
	return bc
}

// SetNillableDispenseMedicinesID sets the dispenseMedicines edge to DispenseMedicine by id if the given value is not nil.
func (bc *BillCreate) SetNillableDispenseMedicinesID(id *int) *BillCreate {
	if id != nil {
		bc = bc.SetDispenseMedicinesID(*id)
	}
	return bc
}

// SetDispenseMedicines sets the dispenseMedicines edge to DispenseMedicine.
func (bc *BillCreate) SetDispenseMedicines(d *DispenseMedicine) *BillCreate {
	return bc.SetDispenseMedicinesID(d.ID)
}

// Mutation returns the BillMutation object of the builder.
func (bc *BillCreate) Mutation() *BillMutation {
	return bc.mutation
}

// Save creates the Bill in the database.
func (bc *BillCreate) Save(ctx context.Context) (*Bill, error) {
	if _, ok := bc.mutation.Amount(); !ok {
		return nil, &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if v, ok := bc.mutation.Amount(); ok {
		if err := bill.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if _, ok := bc.mutation.Annotation(); !ok {
		return nil, &ValidationError{Name: "annotation", err: errors.New("ent: missing required field \"annotation\"")}
	}
	if v, ok := bc.mutation.Annotation(); ok {
		if err := bill.AnnotationValidator(v); err != nil {
			return nil, &ValidationError{Name: "annotation", err: fmt.Errorf("ent: validator failed for field \"annotation\": %w", err)}
		}
	}
	if _, ok := bc.mutation.Payer(); !ok {
		return nil, &ValidationError{Name: "payer", err: errors.New("ent: missing required field \"payer\"")}
	}
	if v, ok := bc.mutation.Payer(); ok {
		if err := bill.PayerValidator(v); err != nil {
			return nil, &ValidationError{Name: "payer", err: fmt.Errorf("ent: validator failed for field \"payer\": %w", err)}
		}
	}
	var (
		err  error
		node *Bill
	)
	if len(bc.hooks) == 0 {
		node, err = bc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*BillMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			bc.mutation = mutation
			node, err = bc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(bc.hooks) - 1; i >= 0; i-- {
			mut = bc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, bc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (bc *BillCreate) SaveX(ctx context.Context) *Bill {
	v, err := bc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (bc *BillCreate) sqlSave(ctx context.Context) (*Bill, error) {
	b, _spec := bc.createSpec()
	if err := sqlgraph.CreateNode(ctx, bc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	b.ID = int(id)
	return b, nil
}

func (bc *BillCreate) createSpec() (*Bill, *sqlgraph.CreateSpec) {
	var (
		b     = &Bill{config: bc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: bill.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: bill.FieldID,
			},
		}
	)
	if value, ok := bc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: bill.FieldAmount,
		})
		b.Amount = value
	}
	if value, ok := bc.mutation.Annotation(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldAnnotation,
		})
		b.Annotation = value
	}
	if value, ok := bc.mutation.Payer(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: bill.FieldPayer,
		})
		b.Payer = value
	}
	if nodes := bc.mutation.PharmacistsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.PaymentsIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := bc.mutation.DispenseMedicinesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return b, _spec
}
