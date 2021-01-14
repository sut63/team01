// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/doctor"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/prescription"
)

// PrescriptionCreate is the builder for creating a Prescription entity.
type PrescriptionCreate struct {
	config
	mutation *PrescriptionMutation
	hooks    []Hook
}

// SetValue sets the Value field.
func (pc *PrescriptionCreate) SetValue(i int) *PrescriptionCreate {
	pc.mutation.SetValue(i)
	return pc
}

// SetPrescriptionpatientID sets the prescriptionpatient edge to PatientInfo by id.
func (pc *PrescriptionCreate) SetPrescriptionpatientID(id int) *PrescriptionCreate {
	pc.mutation.SetPrescriptionpatientID(id)
	return pc
}

// SetNillablePrescriptionpatientID sets the prescriptionpatient edge to PatientInfo by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillablePrescriptionpatientID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetPrescriptionpatientID(*id)
	}
	return pc
}

// SetPrescriptionpatient sets the prescriptionpatient edge to PatientInfo.
func (pc *PrescriptionCreate) SetPrescriptionpatient(p *PatientInfo) *PrescriptionCreate {
	return pc.SetPrescriptionpatientID(p.ID)
}

// SetPrescriptiondoctorID sets the prescriptiondoctor edge to Doctor by id.
func (pc *PrescriptionCreate) SetPrescriptiondoctorID(id int) *PrescriptionCreate {
	pc.mutation.SetPrescriptiondoctorID(id)
	return pc
}

// SetNillablePrescriptiondoctorID sets the prescriptiondoctor edge to Doctor by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillablePrescriptiondoctorID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetPrescriptiondoctorID(*id)
	}
	return pc
}

// SetPrescriptiondoctor sets the prescriptiondoctor edge to Doctor.
func (pc *PrescriptionCreate) SetPrescriptiondoctor(d *Doctor) *PrescriptionCreate {
	return pc.SetPrescriptiondoctorID(d.ID)
}

// SetPrescriptionmedicineID sets the prescriptionmedicine edge to Medicine by id.
func (pc *PrescriptionCreate) SetPrescriptionmedicineID(id int) *PrescriptionCreate {
	pc.mutation.SetPrescriptionmedicineID(id)
	return pc
}

// SetNillablePrescriptionmedicineID sets the prescriptionmedicine edge to Medicine by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillablePrescriptionmedicineID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetPrescriptionmedicineID(*id)
	}
	return pc
}

// SetPrescriptionmedicine sets the prescriptionmedicine edge to Medicine.
func (pc *PrescriptionCreate) SetPrescriptionmedicine(m *Medicine) *PrescriptionCreate {
	return pc.SetPrescriptionmedicineID(m.ID)
}

// SetDispensemedicineID sets the dispensemedicine edge to DispenseMedicine by id.
func (pc *PrescriptionCreate) SetDispensemedicineID(id int) *PrescriptionCreate {
	pc.mutation.SetDispensemedicineID(id)
	return pc
}

// SetNillableDispensemedicineID sets the dispensemedicine edge to DispenseMedicine by id if the given value is not nil.
func (pc *PrescriptionCreate) SetNillableDispensemedicineID(id *int) *PrescriptionCreate {
	if id != nil {
		pc = pc.SetDispensemedicineID(*id)
	}
	return pc
}

// SetDispensemedicine sets the dispensemedicine edge to DispenseMedicine.
func (pc *PrescriptionCreate) SetDispensemedicine(d *DispenseMedicine) *PrescriptionCreate {
	return pc.SetDispensemedicineID(d.ID)
}

// Mutation returns the PrescriptionMutation object of the builder.
func (pc *PrescriptionCreate) Mutation() *PrescriptionMutation {
	return pc.mutation
}

// Save creates the Prescription in the database.
func (pc *PrescriptionCreate) Save(ctx context.Context) (*Prescription, error) {
	if _, ok := pc.mutation.Value(); !ok {
		return nil, &ValidationError{Name: "Value", err: errors.New("ent: missing required field \"Value\"")}
	}
	if v, ok := pc.mutation.Value(); ok {
		if err := prescription.ValueValidator(v); err != nil {
			return nil, &ValidationError{Name: "Value", err: fmt.Errorf("ent: validator failed for field \"Value\": %w", err)}
		}
	}
	var (
		err  error
		node *Prescription
	)
	if len(pc.hooks) == 0 {
		node, err = pc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrescriptionMutation)
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
func (pc *PrescriptionCreate) SaveX(ctx context.Context) *Prescription {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (pc *PrescriptionCreate) sqlSave(ctx context.Context) (*Prescription, error) {
	pr, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	pr.ID = int(id)
	return pr, nil
}

func (pc *PrescriptionCreate) createSpec() (*Prescription, *sqlgraph.CreateSpec) {
	var (
		pr    = &Prescription{config: pc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: prescription.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: prescription.FieldID,
			},
		}
	)
	if value, ok := pc.mutation.Value(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: prescription.FieldValue,
		})
		pr.Value = value
	}
	if nodes := pc.mutation.PrescriptionpatientIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionpatientTable,
			Columns: []string{prescription.PrescriptionpatientColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: patientinfo.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PrescriptiondoctorIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptiondoctorTable,
			Columns: []string{prescription.PrescriptiondoctorColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: doctor.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.PrescriptionmedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   prescription.PrescriptionmedicineTable,
			Columns: []string{prescription.PrescriptionmedicineColumn},
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
	if nodes := pc.mutation.DispensemedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   prescription.DispensemedicineTable,
			Columns: []string{prescription.DispensemedicineColumn},
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
	return pr, _spec
}
