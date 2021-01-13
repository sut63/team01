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
	"github.com/sut63/team01/ent/drugallergy"
	"github.com/sut63/team01/ent/order"
	"github.com/sut63/team01/ent/pharmacist"
)

// PharmacistCreate is the builder for creating a Pharmacist entity.
type PharmacistCreate struct {
	config
	mutation *PharmacistMutation
	hooks    []Hook
}

// SetEmail sets the email field.
func (pc *PharmacistCreate) SetEmail(s string) *PharmacistCreate {
	pc.mutation.SetEmail(s)
	return pc
}

// SetPassword sets the password field.
func (pc *PharmacistCreate) SetPassword(s string) *PharmacistCreate {
	pc.mutation.SetPassword(s)
	return pc
}

// SetName sets the name field.
func (pc *PharmacistCreate) SetName(s string) *PharmacistCreate {
	pc.mutation.SetName(s)
	return pc
}

// AddDispensemedicineIDs adds the dispensemedicine edge to DispenseMedicine by ids.
func (pc *PharmacistCreate) AddDispensemedicineIDs(ids ...int) *PharmacistCreate {
	pc.mutation.AddDispensemedicineIDs(ids...)
	return pc
}

// AddDispensemedicine adds the dispensemedicine edges to DispenseMedicine.
func (pc *PharmacistCreate) AddDispensemedicine(d ...*DispenseMedicine) *PharmacistCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDispensemedicineIDs(ids...)
}

// AddDrugallergyIDs adds the drugallergys edge to DrugAllergy by ids.
func (pc *PharmacistCreate) AddDrugallergyIDs(ids ...int) *PharmacistCreate {
	pc.mutation.AddDrugallergyIDs(ids...)
	return pc
}

// AddDrugallergys adds the drugallergys edges to DrugAllergy.
func (pc *PharmacistCreate) AddDrugallergys(d ...*DrugAllergy) *PharmacistCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return pc.AddDrugallergyIDs(ids...)
}

// AddOrderpharmacistIDs adds the orderpharmacist edge to Order by ids.
func (pc *PharmacistCreate) AddOrderpharmacistIDs(ids ...int) *PharmacistCreate {
	pc.mutation.AddOrderpharmacistIDs(ids...)
	return pc
}

// AddOrderpharmacist adds the orderpharmacist edges to Order.
func (pc *PharmacistCreate) AddOrderpharmacist(o ...*Order) *PharmacistCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return pc.AddOrderpharmacistIDs(ids...)
}

// AddBillIDs adds the Bills edge to Bill by ids.
func (pc *PharmacistCreate) AddBillIDs(ids ...int) *PharmacistCreate {
	pc.mutation.AddBillIDs(ids...)
	return pc
}

// AddBills adds the Bills edges to Bill.
func (pc *PharmacistCreate) AddBills(b ...*Bill) *PharmacistCreate {
	ids := make([]int, len(b))
	for i := range b {
		ids[i] = b[i].ID
	}
	return pc.AddBillIDs(ids...)
}

// Mutation returns the PharmacistMutation object of the builder.
func (pc *PharmacistCreate) Mutation() *PharmacistMutation {
	return pc.mutation
}

// Save creates the Pharmacist in the database.
func (pc *PharmacistCreate) Save(ctx context.Context) (*Pharmacist, error) {
	if _, ok := pc.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := pc.mutation.Email(); ok {
		if err := pharmacist.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if v, ok := pc.mutation.Password(); ok {
		if err := pharmacist.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if _, ok := pc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := pharmacist.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *Pharmacist
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PharmacistMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			pc.mutation = mutation
			node, err = pc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pc.hooks) - 1; i >= 0; i-- {
			mut = pc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pc *PharmacistCreate) SaveX(ctx context.Context) *Pharmacist {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PharmacistCreate) sqlSave(ctx context.Context) (*Pharmacist, error) {
	ph, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	ph.ID = int(id)
	return ph, nil
}

func (pc *PharmacistCreate) createSpec() (*Pharmacist, *sqlgraph.CreateSpec) {
	var (
		ph    = &Pharmacist{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: pharmacist.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: pharmacist.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pharmacist.FieldEmail,
		})
		ph.Email = value
	}
	if value, ok := pc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pharmacist.FieldPassword,
		})
		ph.Password = value
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: pharmacist.FieldName,
		})
		ph.Name = value
	}
	if nodes := pc.mutation.DispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pharmacist.DispensemedicineTable,
			Columns: []string{pharmacist.DispensemedicineColumn},
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
	if nodes := pc.mutation.DrugallergysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pharmacist.DrugallergysTable,
			Columns: []string{pharmacist.DrugallergysColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: drugallergy.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.OrderpharmacistIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pharmacist.OrderpharmacistTable,
			Columns: []string{pharmacist.OrderpharmacistColumn},
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.BillsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   pharmacist.BillsTable,
			Columns: []string{pharmacist.BillsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: bill.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return ph, _spec
}
