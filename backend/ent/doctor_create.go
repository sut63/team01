// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/doctor"
	"github.com/sut63/team01/ent/prescription"
)

// DoctorCreate is the builder for creating a Doctor entity.
type DoctorCreate struct {
	config
	mutation *DoctorMutation
	hooks    []Hook
}

// SetEmail sets the email field.
func (dc *DoctorCreate) SetEmail(s string) *DoctorCreate {
	dc.mutation.SetEmail(s)
	return dc
}

// SetPassword sets the password field.
func (dc *DoctorCreate) SetPassword(s string) *DoctorCreate {
	dc.mutation.SetPassword(s)
	return dc
}

// SetName sets the name field.
func (dc *DoctorCreate) SetName(s string) *DoctorCreate {
	dc.mutation.SetName(s)
	return dc
}

// AddDoctorprescriptionIDs adds the doctorprescription edge to Prescription by ids.
func (dc *DoctorCreate) AddDoctorprescriptionIDs(ids ...int) *DoctorCreate {
	dc.mutation.AddDoctorprescriptionIDs(ids...)
	return dc
}

// AddDoctorprescription adds the doctorprescription edges to Prescription.
func (dc *DoctorCreate) AddDoctorprescription(p ...*Prescription) *DoctorCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return dc.AddDoctorprescriptionIDs(ids...)
}

// Mutation returns the DoctorMutation object of the builder.
func (dc *DoctorCreate) Mutation() *DoctorMutation {
	return dc.mutation
}

// Save creates the Doctor in the database.
func (dc *DoctorCreate) Save(ctx context.Context) (*Doctor, error) {
	if _, ok := dc.mutation.Email(); !ok {
		return nil, &ValidationError{Name: "email", err: errors.New("ent: missing required field \"email\"")}
	}
	if v, ok := dc.mutation.Email(); ok {
		if err := doctor.EmailValidator(v); err != nil {
			return nil, &ValidationError{Name: "email", err: fmt.Errorf("ent: validator failed for field \"email\": %w", err)}
		}
	}
	if _, ok := dc.mutation.Password(); !ok {
		return nil, &ValidationError{Name: "password", err: errors.New("ent: missing required field \"password\"")}
	}
	if v, ok := dc.mutation.Password(); ok {
		if err := doctor.PasswordValidator(v); err != nil {
			return nil, &ValidationError{Name: "password", err: fmt.Errorf("ent: validator failed for field \"password\": %w", err)}
		}
	}
	if _, ok := dc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := dc.mutation.Name(); ok {
		if err := doctor.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	var (
		err  error
		node *Doctor
	)
	if len(dc.hooks) == 0 {
		node, err = dc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*DoctorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			dc.mutation = mutation
			node, err = dc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(dc.hooks) - 1; i >= 0; i-- {
			mut = dc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, dc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (dc *DoctorCreate) SaveX(ctx context.Context) *Doctor {
	v, err := dc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (dc *DoctorCreate) sqlSave(ctx context.Context) (*Doctor, error) {
	d, _spec := dc.createSpec()
	if err := sqlgraph.CreateNode(ctx, dc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	d.ID = int(id)
	return d, nil
}

func (dc *DoctorCreate) createSpec() (*Doctor, *sqlgraph.CreateSpec) {
	var (
		d     = &Doctor{config: dc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: doctor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: doctor.FieldID,
			},
		}
	)
	if value, ok := dc.mutation.Email(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldEmail,
		})
		d.Email = value
	}
	if value, ok := dc.mutation.Password(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldPassword,
		})
		d.Password = value
	}
	if value, ok := dc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: doctor.FieldName,
		})
		d.Name = value
	}
	if nodes := dc.mutation.DoctorprescriptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   doctor.DoctorprescriptionTable,
			Columns: []string{doctor.DoctorprescriptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: prescription.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return d, _spec
}