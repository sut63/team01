// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/company"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/order"
	"github.com/sut63/team01/ent/pharmacist"
)

// OrderCreate is the builder for creating a Order entity.
type OrderCreate struct {
	config
	mutation *OrderMutation
	hooks    []Hook
}

// SetAmount sets the amount field.
func (oc *OrderCreate) SetAmount(i int) *OrderCreate {
	oc.mutation.SetAmount(i)
	return oc
}

// SetPrice sets the price field.
func (oc *OrderCreate) SetPrice(f float64) *OrderCreate {
	oc.mutation.SetPrice(f)
	return oc
}

// SetTotal sets the total field.
func (oc *OrderCreate) SetTotal(f float64) *OrderCreate {
	oc.mutation.SetTotal(f)
	return oc
}

// SetDatetime sets the datetime field.
func (oc *OrderCreate) SetDatetime(t time.Time) *OrderCreate {
	oc.mutation.SetDatetime(t)
	return oc
}

// SetPharmacistID sets the pharmacist edge to Pharmacist by id.
func (oc *OrderCreate) SetPharmacistID(id int) *OrderCreate {
	oc.mutation.SetPharmacistID(id)
	return oc
}

// SetNillablePharmacistID sets the pharmacist edge to Pharmacist by id if the given value is not nil.
func (oc *OrderCreate) SetNillablePharmacistID(id *int) *OrderCreate {
	if id != nil {
		oc = oc.SetPharmacistID(*id)
	}
	return oc
}

// SetPharmacist sets the pharmacist edge to Pharmacist.
func (oc *OrderCreate) SetPharmacist(p *Pharmacist) *OrderCreate {
	return oc.SetPharmacistID(p.ID)
}

// AddCompanyIDs adds the company edge to Company by ids.
func (oc *OrderCreate) AddCompanyIDs(ids ...int) *OrderCreate {
	oc.mutation.AddCompanyIDs(ids...)
	return oc
}

// AddCompany adds the company edges to Company.
func (oc *OrderCreate) AddCompany(c ...*Company) *OrderCreate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return oc.AddCompanyIDs(ids...)
}

// AddMedicineIDs adds the medicine edge to Medicine by ids.
func (oc *OrderCreate) AddMedicineIDs(ids ...int) *OrderCreate {
	oc.mutation.AddMedicineIDs(ids...)
	return oc
}

// AddMedicine adds the medicine edges to Medicine.
func (oc *OrderCreate) AddMedicine(m ...*Medicine) *OrderCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return oc.AddMedicineIDs(ids...)
}

// Mutation returns the OrderMutation object of the builder.
func (oc *OrderCreate) Mutation() *OrderMutation {
	return oc.mutation
}

// Save creates the Order in the database.
func (oc *OrderCreate) Save(ctx context.Context) (*Order, error) {
	if _, ok := oc.mutation.Amount(); !ok {
		return nil, &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if v, ok := oc.mutation.Amount(); ok {
		if err := order.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if _, ok := oc.mutation.Price(); !ok {
		return nil, &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if _, ok := oc.mutation.Total(); !ok {
		return nil, &ValidationError{Name: "total", err: errors.New("ent: missing required field \"total\"")}
	}
	if _, ok := oc.mutation.Datetime(); !ok {
		return nil, &ValidationError{Name: "datetime", err: errors.New("ent: missing required field \"datetime\"")}
	}
	var (
		err  error
		node *Order
	)
	if len(oc.hooks) == 0 {
		node, err = oc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*OrderMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			oc.mutation = mutation
			node, err = oc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(oc.hooks) - 1; i >= 0; i-- {
			mut = oc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, oc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (oc *OrderCreate) SaveX(ctx context.Context) *Order {
	v, err := oc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (oc *OrderCreate) sqlSave(ctx context.Context) (*Order, error) {
	o, _spec := oc.createSpec()
	if err := sqlgraph.CreateNode(ctx, oc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	o.ID = int(id)
	return o, nil
}

func (oc *OrderCreate) createSpec() (*Order, *sqlgraph.CreateSpec) {
	var (
		o     = &Order{config: oc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: order.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: order.FieldID,
			},
		}
	)
	if value, ok := oc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: order.FieldAmount,
		})
		o.Amount = value
	}
	if value, ok := oc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldPrice,
		})
		o.Price = value
	}
	if value, ok := oc.mutation.Total(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: order.FieldTotal,
		})
		o.Total = value
	}
	if value, ok := oc.mutation.Datetime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: order.FieldDatetime,
		})
		o.Datetime = value
	}
	if nodes := oc.mutation.PharmacistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   order.PharmacistTable,
			Columns: []string{order.PharmacistColumn},
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
	if nodes := oc.mutation.CompanyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.CompanyTable,
			Columns: order.CompanyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: company.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := oc.mutation.MedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   order.MedicineTable,
			Columns: order.MedicinePrimaryKey,
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return o, _spec
}
