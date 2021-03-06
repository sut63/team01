// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/sut63/team01/ent/drugallergy"
	"github.com/sut63/team01/ent/levelofdangerous"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/medicinetype"
	"github.com/sut63/team01/ent/order"
	"github.com/sut63/team01/ent/predicate"
	"github.com/sut63/team01/ent/prescription"
	"github.com/sut63/team01/ent/unitofmedicine"
)

// MedicineUpdate is the builder for updating Medicine entities.
type MedicineUpdate struct {
	config
	hooks      []Hook
	mutation   *MedicineMutation
	predicates []predicate.Medicine
}

// Where adds a new predicate for the builder.
func (mu *MedicineUpdate) Where(ps ...predicate.Medicine) *MedicineUpdate {
	mu.predicates = append(mu.predicates, ps...)
	return mu
}

// SetName sets the name field.
func (mu *MedicineUpdate) SetName(s string) *MedicineUpdate {
	mu.mutation.SetName(s)
	return mu
}

// SetSerial sets the serial field.
func (mu *MedicineUpdate) SetSerial(s string) *MedicineUpdate {
	mu.mutation.SetSerial(s)
	return mu
}

// SetBrand sets the brand field.
func (mu *MedicineUpdate) SetBrand(s string) *MedicineUpdate {
	mu.mutation.SetBrand(s)
	return mu
}

// SetAmount sets the amount field.
func (mu *MedicineUpdate) SetAmount(i int) *MedicineUpdate {
	mu.mutation.ResetAmount()
	mu.mutation.SetAmount(i)
	return mu
}

// AddAmount adds i to amount.
func (mu *MedicineUpdate) AddAmount(i int) *MedicineUpdate {
	mu.mutation.AddAmount(i)
	return mu
}

// SetPrice sets the price field.
func (mu *MedicineUpdate) SetPrice(i int) *MedicineUpdate {
	mu.mutation.ResetPrice()
	mu.mutation.SetPrice(i)
	return mu
}

// AddPrice adds i to price.
func (mu *MedicineUpdate) AddPrice(i int) *MedicineUpdate {
	mu.mutation.AddPrice(i)
	return mu
}

// SetHowtouse sets the howtouse field.
func (mu *MedicineUpdate) SetHowtouse(s string) *MedicineUpdate {
	mu.mutation.SetHowtouse(s)
	return mu
}

// SetLevelOfDangerousID sets the LevelOfDangerous edge to LevelOfDangerous by id.
func (mu *MedicineUpdate) SetLevelOfDangerousID(id int) *MedicineUpdate {
	mu.mutation.SetLevelOfDangerousID(id)
	return mu
}

// SetNillableLevelOfDangerousID sets the LevelOfDangerous edge to LevelOfDangerous by id if the given value is not nil.
func (mu *MedicineUpdate) SetNillableLevelOfDangerousID(id *int) *MedicineUpdate {
	if id != nil {
		mu = mu.SetLevelOfDangerousID(*id)
	}
	return mu
}

// SetLevelOfDangerous sets the LevelOfDangerous edge to LevelOfDangerous.
func (mu *MedicineUpdate) SetLevelOfDangerous(l *LevelOfDangerous) *MedicineUpdate {
	return mu.SetLevelOfDangerousID(l.ID)
}

// SetMedicineTypeID sets the MedicineType edge to MedicineType by id.
func (mu *MedicineUpdate) SetMedicineTypeID(id int) *MedicineUpdate {
	mu.mutation.SetMedicineTypeID(id)
	return mu
}

// SetNillableMedicineTypeID sets the MedicineType edge to MedicineType by id if the given value is not nil.
func (mu *MedicineUpdate) SetNillableMedicineTypeID(id *int) *MedicineUpdate {
	if id != nil {
		mu = mu.SetMedicineTypeID(*id)
	}
	return mu
}

// SetMedicineType sets the MedicineType edge to MedicineType.
func (mu *MedicineUpdate) SetMedicineType(m *MedicineType) *MedicineUpdate {
	return mu.SetMedicineTypeID(m.ID)
}

// SetUnitOfMedicineID sets the UnitOfMedicine edge to UnitOfMedicine by id.
func (mu *MedicineUpdate) SetUnitOfMedicineID(id int) *MedicineUpdate {
	mu.mutation.SetUnitOfMedicineID(id)
	return mu
}

// SetNillableUnitOfMedicineID sets the UnitOfMedicine edge to UnitOfMedicine by id if the given value is not nil.
func (mu *MedicineUpdate) SetNillableUnitOfMedicineID(id *int) *MedicineUpdate {
	if id != nil {
		mu = mu.SetUnitOfMedicineID(*id)
	}
	return mu
}

// SetUnitOfMedicine sets the UnitOfMedicine edge to UnitOfMedicine.
func (mu *MedicineUpdate) SetUnitOfMedicine(u *UnitOfMedicine) *MedicineUpdate {
	return mu.SetUnitOfMedicineID(u.ID)
}

// AddDrugallergyIDs adds the drugallergys edge to DrugAllergy by ids.
func (mu *MedicineUpdate) AddDrugallergyIDs(ids ...int) *MedicineUpdate {
	mu.mutation.AddDrugallergyIDs(ids...)
	return mu
}

// AddDrugallergys adds the drugallergys edges to DrugAllergy.
func (mu *MedicineUpdate) AddDrugallergys(d ...*DrugAllergy) *MedicineUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.AddDrugallergyIDs(ids...)
}

// AddMedicinepresciptionIDs adds the medicinepresciption edge to Prescription by ids.
func (mu *MedicineUpdate) AddMedicinepresciptionIDs(ids ...int) *MedicineUpdate {
	mu.mutation.AddMedicinepresciptionIDs(ids...)
	return mu
}

// AddMedicinepresciption adds the medicinepresciption edges to Prescription.
func (mu *MedicineUpdate) AddMedicinepresciption(p ...*Prescription) *MedicineUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.AddMedicinepresciptionIDs(ids...)
}

// AddOrdermedicineIDs adds the ordermedicine edge to Order by ids.
func (mu *MedicineUpdate) AddOrdermedicineIDs(ids ...int) *MedicineUpdate {
	mu.mutation.AddOrdermedicineIDs(ids...)
	return mu
}

// AddOrdermedicine adds the ordermedicine edges to Order.
func (mu *MedicineUpdate) AddOrdermedicine(o ...*Order) *MedicineUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.AddOrdermedicineIDs(ids...)
}

// Mutation returns the MedicineMutation object of the builder.
func (mu *MedicineUpdate) Mutation() *MedicineMutation {
	return mu.mutation
}

// ClearLevelOfDangerous clears the LevelOfDangerous edge to LevelOfDangerous.
func (mu *MedicineUpdate) ClearLevelOfDangerous() *MedicineUpdate {
	mu.mutation.ClearLevelOfDangerous()
	return mu
}

// ClearMedicineType clears the MedicineType edge to MedicineType.
func (mu *MedicineUpdate) ClearMedicineType() *MedicineUpdate {
	mu.mutation.ClearMedicineType()
	return mu
}

// ClearUnitOfMedicine clears the UnitOfMedicine edge to UnitOfMedicine.
func (mu *MedicineUpdate) ClearUnitOfMedicine() *MedicineUpdate {
	mu.mutation.ClearUnitOfMedicine()
	return mu
}

// RemoveDrugallergyIDs removes the drugallergys edge to DrugAllergy by ids.
func (mu *MedicineUpdate) RemoveDrugallergyIDs(ids ...int) *MedicineUpdate {
	mu.mutation.RemoveDrugallergyIDs(ids...)
	return mu
}

// RemoveDrugallergys removes drugallergys edges to DrugAllergy.
func (mu *MedicineUpdate) RemoveDrugallergys(d ...*DrugAllergy) *MedicineUpdate {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return mu.RemoveDrugallergyIDs(ids...)
}

// RemoveMedicinepresciptionIDs removes the medicinepresciption edge to Prescription by ids.
func (mu *MedicineUpdate) RemoveMedicinepresciptionIDs(ids ...int) *MedicineUpdate {
	mu.mutation.RemoveMedicinepresciptionIDs(ids...)
	return mu
}

// RemoveMedicinepresciption removes medicinepresciption edges to Prescription.
func (mu *MedicineUpdate) RemoveMedicinepresciption(p ...*Prescription) *MedicineUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return mu.RemoveMedicinepresciptionIDs(ids...)
}

// RemoveOrdermedicineIDs removes the ordermedicine edge to Order by ids.
func (mu *MedicineUpdate) RemoveOrdermedicineIDs(ids ...int) *MedicineUpdate {
	mu.mutation.RemoveOrdermedicineIDs(ids...)
	return mu
}

// RemoveOrdermedicine removes ordermedicine edges to Order.
func (mu *MedicineUpdate) RemoveOrdermedicine(o ...*Order) *MedicineUpdate {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return mu.RemoveOrdermedicineIDs(ids...)
}

// Save executes the query and returns the number of rows/vertices matched by this operation.
func (mu *MedicineUpdate) Save(ctx context.Context) (int, error) {
	if v, ok := mu.mutation.Name(); ok {
		if err := medicine.NameValidator(v); err != nil {
			return 0, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Serial(); ok {
		if err := medicine.SerialValidator(v); err != nil {
			return 0, &ValidationError{Name: "serial", err: fmt.Errorf("ent: validator failed for field \"serial\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Brand(); ok {
		if err := medicine.BrandValidator(v); err != nil {
			return 0, &ValidationError{Name: "brand", err: fmt.Errorf("ent: validator failed for field \"brand\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Amount(); ok {
		if err := medicine.AmountValidator(v); err != nil {
			return 0, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Price(); ok {
		if err := medicine.PriceValidator(v); err != nil {
			return 0, &ValidationError{Name: "price", err: fmt.Errorf("ent: validator failed for field \"price\": %w", err)}
		}
	}
	if v, ok := mu.mutation.Howtouse(); ok {
		if err := medicine.HowtouseValidator(v); err != nil {
			return 0, &ValidationError{Name: "howtouse", err: fmt.Errorf("ent: validator failed for field \"howtouse\": %w", err)}
		}
	}

	var (
		err      error
		affected int
	)
	if len(mu.hooks) == 0 {
		affected, err = mu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			mu.mutation = mutation
			affected, err = mu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(mu.hooks) - 1; i >= 0; i-- {
			mut = mu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, mu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (mu *MedicineUpdate) SaveX(ctx context.Context) int {
	affected, err := mu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (mu *MedicineUpdate) Exec(ctx context.Context) error {
	_, err := mu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mu *MedicineUpdate) ExecX(ctx context.Context) {
	if err := mu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (mu *MedicineUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicine.Table,
			Columns: medicine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicine.FieldID,
			},
		},
	}
	if ps := mu.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := mu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldName,
		})
	}
	if value, ok := mu.mutation.Serial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldSerial,
		})
	}
	if value, ok := mu.mutation.Brand(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldBrand,
		})
	}
	if value, ok := mu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldAmount,
		})
	}
	if value, ok := mu.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldAmount,
		})
	}
	if value, ok := mu.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldPrice,
		})
	}
	if value, ok := mu.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldPrice,
		})
	}
	if value, ok := mu.mutation.Howtouse(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldHowtouse,
		})
	}
	if mu.mutation.LevelOfDangerousCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.LevelOfDangerousIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.MedicineTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MedicineTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if mu.mutation.UnitOfMedicineCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.UnitOfMedicineIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedDrugallergysIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.DrugallergysIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedMedicinepresciptionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := mu.mutation.MedicinepresciptionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := mu.mutation.RemovedOrdermedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.OrdermedicineTable,
			Columns: []string{medicine.OrdermedicineColumn},
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
	if nodes := mu.mutation.OrdermedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.OrdermedicineTable,
			Columns: []string{medicine.OrdermedicineColumn},
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
	if n, err = sqlgraph.UpdateNodes(ctx, mu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicine.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return 0, err
	}
	return n, nil
}

// MedicineUpdateOne is the builder for updating a single Medicine entity.
type MedicineUpdateOne struct {
	config
	hooks    []Hook
	mutation *MedicineMutation
}

// SetName sets the name field.
func (muo *MedicineUpdateOne) SetName(s string) *MedicineUpdateOne {
	muo.mutation.SetName(s)
	return muo
}

// SetSerial sets the serial field.
func (muo *MedicineUpdateOne) SetSerial(s string) *MedicineUpdateOne {
	muo.mutation.SetSerial(s)
	return muo
}

// SetBrand sets the brand field.
func (muo *MedicineUpdateOne) SetBrand(s string) *MedicineUpdateOne {
	muo.mutation.SetBrand(s)
	return muo
}

// SetAmount sets the amount field.
func (muo *MedicineUpdateOne) SetAmount(i int) *MedicineUpdateOne {
	muo.mutation.ResetAmount()
	muo.mutation.SetAmount(i)
	return muo
}

// AddAmount adds i to amount.
func (muo *MedicineUpdateOne) AddAmount(i int) *MedicineUpdateOne {
	muo.mutation.AddAmount(i)
	return muo
}

// SetPrice sets the price field.
func (muo *MedicineUpdateOne) SetPrice(i int) *MedicineUpdateOne {
	muo.mutation.ResetPrice()
	muo.mutation.SetPrice(i)
	return muo
}

// AddPrice adds i to price.
func (muo *MedicineUpdateOne) AddPrice(i int) *MedicineUpdateOne {
	muo.mutation.AddPrice(i)
	return muo
}

// SetHowtouse sets the howtouse field.
func (muo *MedicineUpdateOne) SetHowtouse(s string) *MedicineUpdateOne {
	muo.mutation.SetHowtouse(s)
	return muo
}

// SetLevelOfDangerousID sets the LevelOfDangerous edge to LevelOfDangerous by id.
func (muo *MedicineUpdateOne) SetLevelOfDangerousID(id int) *MedicineUpdateOne {
	muo.mutation.SetLevelOfDangerousID(id)
	return muo
}

// SetNillableLevelOfDangerousID sets the LevelOfDangerous edge to LevelOfDangerous by id if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableLevelOfDangerousID(id *int) *MedicineUpdateOne {
	if id != nil {
		muo = muo.SetLevelOfDangerousID(*id)
	}
	return muo
}

// SetLevelOfDangerous sets the LevelOfDangerous edge to LevelOfDangerous.
func (muo *MedicineUpdateOne) SetLevelOfDangerous(l *LevelOfDangerous) *MedicineUpdateOne {
	return muo.SetLevelOfDangerousID(l.ID)
}

// SetMedicineTypeID sets the MedicineType edge to MedicineType by id.
func (muo *MedicineUpdateOne) SetMedicineTypeID(id int) *MedicineUpdateOne {
	muo.mutation.SetMedicineTypeID(id)
	return muo
}

// SetNillableMedicineTypeID sets the MedicineType edge to MedicineType by id if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableMedicineTypeID(id *int) *MedicineUpdateOne {
	if id != nil {
		muo = muo.SetMedicineTypeID(*id)
	}
	return muo
}

// SetMedicineType sets the MedicineType edge to MedicineType.
func (muo *MedicineUpdateOne) SetMedicineType(m *MedicineType) *MedicineUpdateOne {
	return muo.SetMedicineTypeID(m.ID)
}

// SetUnitOfMedicineID sets the UnitOfMedicine edge to UnitOfMedicine by id.
func (muo *MedicineUpdateOne) SetUnitOfMedicineID(id int) *MedicineUpdateOne {
	muo.mutation.SetUnitOfMedicineID(id)
	return muo
}

// SetNillableUnitOfMedicineID sets the UnitOfMedicine edge to UnitOfMedicine by id if the given value is not nil.
func (muo *MedicineUpdateOne) SetNillableUnitOfMedicineID(id *int) *MedicineUpdateOne {
	if id != nil {
		muo = muo.SetUnitOfMedicineID(*id)
	}
	return muo
}

// SetUnitOfMedicine sets the UnitOfMedicine edge to UnitOfMedicine.
func (muo *MedicineUpdateOne) SetUnitOfMedicine(u *UnitOfMedicine) *MedicineUpdateOne {
	return muo.SetUnitOfMedicineID(u.ID)
}

// AddDrugallergyIDs adds the drugallergys edge to DrugAllergy by ids.
func (muo *MedicineUpdateOne) AddDrugallergyIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.AddDrugallergyIDs(ids...)
	return muo
}

// AddDrugallergys adds the drugallergys edges to DrugAllergy.
func (muo *MedicineUpdateOne) AddDrugallergys(d ...*DrugAllergy) *MedicineUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.AddDrugallergyIDs(ids...)
}

// AddMedicinepresciptionIDs adds the medicinepresciption edge to Prescription by ids.
func (muo *MedicineUpdateOne) AddMedicinepresciptionIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.AddMedicinepresciptionIDs(ids...)
	return muo
}

// AddMedicinepresciption adds the medicinepresciption edges to Prescription.
func (muo *MedicineUpdateOne) AddMedicinepresciption(p ...*Prescription) *MedicineUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.AddMedicinepresciptionIDs(ids...)
}

// AddOrdermedicineIDs adds the ordermedicine edge to Order by ids.
func (muo *MedicineUpdateOne) AddOrdermedicineIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.AddOrdermedicineIDs(ids...)
	return muo
}

// AddOrdermedicine adds the ordermedicine edges to Order.
func (muo *MedicineUpdateOne) AddOrdermedicine(o ...*Order) *MedicineUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.AddOrdermedicineIDs(ids...)
}

// Mutation returns the MedicineMutation object of the builder.
func (muo *MedicineUpdateOne) Mutation() *MedicineMutation {
	return muo.mutation
}

// ClearLevelOfDangerous clears the LevelOfDangerous edge to LevelOfDangerous.
func (muo *MedicineUpdateOne) ClearLevelOfDangerous() *MedicineUpdateOne {
	muo.mutation.ClearLevelOfDangerous()
	return muo
}

// ClearMedicineType clears the MedicineType edge to MedicineType.
func (muo *MedicineUpdateOne) ClearMedicineType() *MedicineUpdateOne {
	muo.mutation.ClearMedicineType()
	return muo
}

// ClearUnitOfMedicine clears the UnitOfMedicine edge to UnitOfMedicine.
func (muo *MedicineUpdateOne) ClearUnitOfMedicine() *MedicineUpdateOne {
	muo.mutation.ClearUnitOfMedicine()
	return muo
}

// RemoveDrugallergyIDs removes the drugallergys edge to DrugAllergy by ids.
func (muo *MedicineUpdateOne) RemoveDrugallergyIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.RemoveDrugallergyIDs(ids...)
	return muo
}

// RemoveDrugallergys removes drugallergys edges to DrugAllergy.
func (muo *MedicineUpdateOne) RemoveDrugallergys(d ...*DrugAllergy) *MedicineUpdateOne {
	ids := make([]int, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return muo.RemoveDrugallergyIDs(ids...)
}

// RemoveMedicinepresciptionIDs removes the medicinepresciption edge to Prescription by ids.
func (muo *MedicineUpdateOne) RemoveMedicinepresciptionIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.RemoveMedicinepresciptionIDs(ids...)
	return muo
}

// RemoveMedicinepresciption removes medicinepresciption edges to Prescription.
func (muo *MedicineUpdateOne) RemoveMedicinepresciption(p ...*Prescription) *MedicineUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return muo.RemoveMedicinepresciptionIDs(ids...)
}

// RemoveOrdermedicineIDs removes the ordermedicine edge to Order by ids.
func (muo *MedicineUpdateOne) RemoveOrdermedicineIDs(ids ...int) *MedicineUpdateOne {
	muo.mutation.RemoveOrdermedicineIDs(ids...)
	return muo
}

// RemoveOrdermedicine removes ordermedicine edges to Order.
func (muo *MedicineUpdateOne) RemoveOrdermedicine(o ...*Order) *MedicineUpdateOne {
	ids := make([]int, len(o))
	for i := range o {
		ids[i] = o[i].ID
	}
	return muo.RemoveOrdermedicineIDs(ids...)
}

// Save executes the query and returns the updated entity.
func (muo *MedicineUpdateOne) Save(ctx context.Context) (*Medicine, error) {
	if v, ok := muo.mutation.Name(); ok {
		if err := medicine.NameValidator(v); err != nil {
			return nil, &ValidationError{Name: "name", err: fmt.Errorf("ent: validator failed for field \"name\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Serial(); ok {
		if err := medicine.SerialValidator(v); err != nil {
			return nil, &ValidationError{Name: "serial", err: fmt.Errorf("ent: validator failed for field \"serial\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Brand(); ok {
		if err := medicine.BrandValidator(v); err != nil {
			return nil, &ValidationError{Name: "brand", err: fmt.Errorf("ent: validator failed for field \"brand\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Amount(); ok {
		if err := medicine.AmountValidator(v); err != nil {
			return nil, &ValidationError{Name: "amount", err: fmt.Errorf("ent: validator failed for field \"amount\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Price(); ok {
		if err := medicine.PriceValidator(v); err != nil {
			return nil, &ValidationError{Name: "price", err: fmt.Errorf("ent: validator failed for field \"price\": %w", err)}
		}
	}
	if v, ok := muo.mutation.Howtouse(); ok {
		if err := medicine.HowtouseValidator(v); err != nil {
			return nil, &ValidationError{Name: "howtouse", err: fmt.Errorf("ent: validator failed for field \"howtouse\": %w", err)}
		}
	}

	var (
		err  error
		node *Medicine
	)
	if len(muo.hooks) == 0 {
		node, err = muo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MedicineMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			muo.mutation = mutation
			node, err = muo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(muo.hooks) - 1; i >= 0; i-- {
			mut = muo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, muo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (muo *MedicineUpdateOne) SaveX(ctx context.Context) *Medicine {
	m, err := muo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return m
}

// Exec executes the query on the entity.
func (muo *MedicineUpdateOne) Exec(ctx context.Context) error {
	_, err := muo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (muo *MedicineUpdateOne) ExecX(ctx context.Context) {
	if err := muo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (muo *MedicineUpdateOne) sqlSave(ctx context.Context) (m *Medicine, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   medicine.Table,
			Columns: medicine.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: medicine.FieldID,
			},
		},
	}
	id, ok := muo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Medicine.ID for update")}
	}
	_spec.Node.ID.Value = id
	if value, ok := muo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldName,
		})
	}
	if value, ok := muo.mutation.Serial(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldSerial,
		})
	}
	if value, ok := muo.mutation.Brand(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldBrand,
		})
	}
	if value, ok := muo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldAmount,
		})
	}
	if value, ok := muo.mutation.AddedAmount(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldAmount,
		})
	}
	if value, ok := muo.mutation.Price(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldPrice,
		})
	}
	if value, ok := muo.mutation.AddedPrice(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: medicine.FieldPrice,
		})
	}
	if value, ok := muo.mutation.Howtouse(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: medicine.FieldHowtouse,
		})
	}
	if muo.mutation.LevelOfDangerousCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.LevelOfDangerousIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.MedicineTypeCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MedicineTypeIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if muo.mutation.UnitOfMedicineCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.UnitOfMedicineIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedDrugallergysIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.DrugallergysIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedMedicinepresciptionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := muo.mutation.MedicinepresciptionIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if nodes := muo.mutation.RemovedOrdermedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.OrdermedicineTable,
			Columns: []string{medicine.OrdermedicineColumn},
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
	if nodes := muo.mutation.OrdermedicineIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   medicine.OrdermedicineTable,
			Columns: []string{medicine.OrdermedicineColumn},
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
	m = &Medicine{config: muo.config}
	_spec.Assign = m.assignValues
	_spec.ScanValues = m.scanValues()
	if err = sqlgraph.UpdateNode(ctx, muo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{medicine.Label}
		} else if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return m, nil
}
