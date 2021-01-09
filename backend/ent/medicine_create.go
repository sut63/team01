// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/drugallergy"
	"github.com/sut63/team01/ent/levelofdangerous"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/medicinetype"
	"github.com/sut63/team01/ent/order"
	"github.com/sut63/team01/ent/prescription"
	"github.com/sut63/team01/ent/unitofmedicine"
)

// MedicineCreate is the builder for creating a Medicine entity.
type MedicineCreate struct {
	config
	mutation *MedicineMutation
	hooks    []Hook
}

// SetName sets the name field.
func (mc *MedicineCreate) SetName(s string) *MedicineCreate {
	mc.mutation.SetName(s)
	return mc
}

// SetSerial sets the serial field.
func (mc *MedicineCreate) SetSerial(s string) *MedicineCreate {
	mc.mutation.SetSerial(s)
	return mc
}

// SetBrand sets the brand field.
func (mc *MedicineCreate) SetBrand(s string) *MedicineCreate {
	mc.mutation.SetBrand(s)
	return mc
}

// SetAmount sets the amount field.
func (mc *MedicineCreate) SetAmount(i int) *MedicineCreate {
	mc.mutation.SetAmount(i)
	return mc
}

// SetPrice sets the price field.
func (mc *MedicineCreate) SetPrice(f float64) *MedicineCreate {
	mc.mutation.SetPrice(f)
	return mc
}

// SetHowtouse sets the howtouse field.
func (mc *MedicineCreate) SetHowtouse(s string) *MedicineCreate {
	mc.mutation.SetHowtouse(s)
	return mc
}

// SetLevelOfDangerousID sets the LevelOfDangerous edge to LevelOfDangerous by id.
func (mc *MedicineCreate) SetLevelOfDangerousID(id int) *MedicineCreate {
	mc.mutation.SetLevelOfDangerousID(id)
	return mc
}

// SetNillableLevelOfDangerousID sets the LevelOfDangerous edge to LevelOfDangerous by id if the given value is not nil.
func (mc *MedicineCreate) SetNillableLevelOfDangerousID(id *int) *MedicineCreate {
	if id != nil {
		mc = mc.SetLevelOfDangerousID(*id)
	}
	return mc
}

// SetLevelOfDangerous sets the LevelOfDangerous edge to LevelOfDangerous.
func (mc *MedicineCreate) SetLevelOfDangerous(l *LevelOfDangerous) *MedicineCreate {
	return mc.SetLevelOfDangerousID(l.ID)
}

// SetMedicineTypeID sets the MedicineType edge to MedicineType by id.
func (mc *MedicineCreate) SetMedicineTypeID(id int) *MedicineCreate {
	mc.mutation.SetMedicineTypeID(id)
	return mc
}

// SetNillableMedicineTypeID sets the MedicineType edge to MedicineType by id if the given value is not nil.
func (mc *MedicineCreate) SetNillableMedicineTypeID(id *int) *MedicineCreate {
	if id != nil {
		mc = mc.SetMedicineTypeID(*id)
	}
	return mc
}

// SetMedicineType sets the MedicineType edge to MedicineType.
func (mc *MedicineCreate) SetMedicineType(m *MedicineType) *MedicineCreate {
	return mc.SetMedicineTypeID(m.ID)
}

// SetUnitOfMedicineID sets the UnitOfMedicine edge to UnitOfMedicine by id.
func (mc *MedicineCreate) SetUnitOfMedicineID(id int) *MedicineCreate {
	mc.mutation.SetUnitOfMedicineID(id)
	return mc
}

// SetNillableUnitOfMedicineID sets the UnitOfMedicine edge to UnitOfMedicine by id if the given value is not nil.
func (mc *MedicineCreate) SetNillableUnitOfMedicineID(id *int) *MedicineCreate {
	if id != nil {
		mc = mc.SetUnitOfMedicineID(*id)
	}
	return mc
}

// SetUnitOfMedicine sets the UnitOfMedicine edge to UnitOfMedicine.
func (mc *MedicineCreate) SetUnitOfMedicine(u *UnitOfMedicine) *MedicineCreate {
	return mc.SetUnitOfMedicineID(u.ID)
}

// AddDrugallergyIDs adds the drugallergys edge to DrugAllergy by ids.
func (mc *MedicineCreate) AddDrugallergyIDs(ids ...int) *MedicineCreate {
	mc.mutation.AddDrugallergyIDs(ids...)
	return mc
}

// AddDrugallergys adds the drugallergys edges to DrugAllergy.
func (mc *MedicineCreate) AddDrugallergys(d ...*DrugAllergy) *MedicineCreate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mc.AddDrugallergyIDs(ids...)
}

// AddMedicinepresciptionIDs adds the medicinepresciption edge to Prescription by ids.
func (mc *MedicineCreate) AddMedicinepresciptionIDs(ids ...int) *MedicineCreate {
	mc.mutation.AddMedicinepresciptionIDs(ids...)
	return mc
}

// AddMedicinepresciption adds the medicinepresciption edges to Prescription.
func (mc *MedicineCreate) AddMedicinepresciption(p ...*Prescription) *MedicineCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mc.AddMedicinepresciptionIDs(ids...)
}

// AddOrderIDs adds the order edge to Order by ids.
func (mc *MedicineCreate) AddOrderIDs(ids ...int) *MedicineCreate {
	mc.mutation.AddOrderIDs(ids...)
	return mc
}

// AddOrder adds the order edges to Order.
func (mc *MedicineCreate) AddOrder(o ...*Order) *MedicineCreate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mc.AddOrderIDs(ids...)
}

// Mutation returns the MedicineMutation object of the builder.
func (mc *MedicineCreate) Mutation() *MedicineMutation {
	return mc.mutation
}

// Save creates the Medicine in the database.
func (mc *MedicineCreate) Save(ctx context.Context) (*Medicine, error) {
	if _, ok := mc.mutation.Name(); !ok {
		return nil, &ValidationError{Name: "name", err: errors.New("ent: missing required field \"name\"")}
	}
	if v, ok := mc.mutation.Name(); ok {
		if err := medicine.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Serial(); !ok {
		return nil, &ValidationError{Name: "serial", err: errors.New("ent: missing required field \"serial\"")}
	}
	if v, ok := mc.mutation.Serial(); ok {
		if err := medicine.SerialValidator(v); err != nil {
			return nil, &ValidationError{Name: "serial", err: fmt.Errorf("ent: validator failed for field \"serial\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Brand(); !ok {
		return nil, &ValidationError{Name: "brand", err: errors.New("ent: missing required field \"brand\"")}
	}
	if v, ok := mc.mutation.Brand(); ok {
		if err := medicine.BrandValidator(v); err != nil {
			return nil, &ValidationError{Name: "brand", err: fmt.Errorf("ent: validator failed for field \"brand\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Amount(); !ok {
		return nil, &ValidationError{Name: "amount", err: errors.New("ent: missing required field \"amount\"")}
	}
	if v, ok := mc.mutation.Amount(); ok {
		if err := medicine.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Price(); !ok {
		return nil, &ValidationError{Name: "price", err: errors.New("ent: missing required field \"price\"")}
	}
	if v, ok := mc.mutation.Price(); ok {
		if err := medicine.PriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "price", err: fmt.Errorf("ent: validator failed for field \"price\": %w", err)}
		}
	}
	if _, ok := mc.mutation.Howtouse(); !ok {
		return nil, &ValidationError{Name: "howtouse", err: errors.New("ent: missing required field \"howtouse\"")}
	}
	if v, ok := mc.mutation.Howtouse(); ok {
		if err := medicine.HowtouseValidator(v); err != nil {
			return nil, &ValidationError{Name: "howtouse", err: fmt.Errorf("ent: validator failed for field \"howtouse\": %w", err)}
		}
	}
	var (
		err  error
		node *Medicine
	)
	if len(mc.hooks) == 0 {
		node, err = mc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mc.mutation = mutation
			node, err = mc.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(mc.hooks) - 1; i >= 0; i-- {
			mut = mc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MedicineCreate) SaveX(ctx context.Context) *Medicine {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (mc *MedicineCreate) sqlSave(ctx context.Context) (*Medicine, error) {
	m, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	m.ID = int(id)
	return m, nil
}

func (mc *MedicineCreate) createSpec() (*Medicine, *sqlgraph.CreateSpec) {
	var (
		m     = &Medicine{config: mc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: medicine.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicine.FieldID,
			},
		}
	)
	if value, ok := mc.mutation.Name(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldName,
		})
		m.Name = value
	}
	if value, ok := mc.mutation.Serial(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldSerial,
		})
		m.Serial = value
	}
	if value, ok := mc.mutation.Brand(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldBrand,
		})
		m.Brand = value
	}
	if value, ok := mc.mutation.Amount(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldAmount,
		})
		m.Amount = value
	}
	if value, ok := mc.mutation.Price(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeFloat64,
			Value:  value,
			Column: medicine.FieldPrice,
		})
		m.Price = value
	}
	if value, ok := mc.mutation.Howtouse(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldHowtouse,
		})
		m.Howtouse = value
	}
	if nodes := mc.mutation.LevelOfDangerousIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicine.LevelOfDangerousTable,
			Columns: []string{medicine.LevelOfDangerousColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: levelofdangerous.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.MedicineTypeIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicine.MedicineTypeTable,
			Columns: []string{medicine.MedicineTypeColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: medicinetype.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.UnitOfMedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   medicine.UnitOfMedicineTable,
			Columns: []string{medicine.UnitOfMedicineColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: unitofmedicine.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := mc.mutation.DrugallergysIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.DrugallergysTable,
			Columns: []string{medicine.DrugallergysColumn},
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
	if nodes := mc.mutation.MedicinepresciptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.MedicinepresciptionTable,
			Columns: []string{medicine.MedicinepresciptionColumn},
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
	if nodes := mc.mutation.OrderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.OrderTable,
			Columns: []string{medicine.OrderColumn},
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
	return m, _spec
}
