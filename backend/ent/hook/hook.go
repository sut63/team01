// Code generated by entc, DO NOT EDIT.

package hook

import (
	"context"
	"fmt"

	"github.com/sut63/team01/ent"
)

// The AnnotationFunc type is an adapter to allow the use of ordinary
// function as Annotation mutator.
type AnnotationFunc func(context.Context, *ent.AnnotationMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f AnnotationFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.AnnotationMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.AnnotationMutation", m)
	}
	return f(ctx, mv)
}

// The BillFunc type is an adapter to allow the use of ordinary
// function as Bill mutator.
type BillFunc func(context.Context, *ent.BillMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f BillFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.BillMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.BillMutation", m)
	}
	return f(ctx, mv)
}

// The CompanyFunc type is an adapter to allow the use of ordinary
// function as Company mutator.
type CompanyFunc func(context.Context, *ent.CompanyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f CompanyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.CompanyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.CompanyMutation", m)
	}
	return f(ctx, mv)
}

// The DispenseMedicineFunc type is an adapter to allow the use of ordinary
// function as DispenseMedicine mutator.
type DispenseMedicineFunc func(context.Context, *ent.DispenseMedicineMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DispenseMedicineFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DispenseMedicineMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DispenseMedicineMutation", m)
	}
	return f(ctx, mv)
}

// The DoctorFunc type is an adapter to allow the use of ordinary
// function as Doctor mutator.
type DoctorFunc func(context.Context, *ent.DoctorMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DoctorFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DoctorMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DoctorMutation", m)
	}
	return f(ctx, mv)
}

// The DrugAllergyFunc type is an adapter to allow the use of ordinary
// function as DrugAllergy mutator.
type DrugAllergyFunc func(context.Context, *ent.DrugAllergyMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f DrugAllergyFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.DrugAllergyMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.DrugAllergyMutation", m)
	}
	return f(ctx, mv)
}

// The LevelOfDangerousFunc type is an adapter to allow the use of ordinary
// function as LevelOfDangerous mutator.
type LevelOfDangerousFunc func(context.Context, *ent.LevelOfDangerousMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f LevelOfDangerousFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.LevelOfDangerousMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.LevelOfDangerousMutation", m)
	}
	return f(ctx, mv)
}

// The MedicineFunc type is an adapter to allow the use of ordinary
// function as Medicine mutator.
type MedicineFunc func(context.Context, *ent.MedicineMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MedicineFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MedicineMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MedicineMutation", m)
	}
	return f(ctx, mv)
}

// The MedicineTypeFunc type is an adapter to allow the use of ordinary
// function as MedicineType mutator.
type MedicineTypeFunc func(context.Context, *ent.MedicineTypeMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f MedicineTypeFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.MedicineTypeMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.MedicineTypeMutation", m)
	}
	return f(ctx, mv)
}

// The OrderFunc type is an adapter to allow the use of ordinary
// function as Order mutator.
type OrderFunc func(context.Context, *ent.OrderMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f OrderFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.OrderMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.OrderMutation", m)
	}
	return f(ctx, mv)
}

// The PatientInfoFunc type is an adapter to allow the use of ordinary
// function as PatientInfo mutator.
type PatientInfoFunc func(context.Context, *ent.PatientInfoMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PatientInfoFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PatientInfoMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PatientInfoMutation", m)
	}
	return f(ctx, mv)
}

// The PaymentFunc type is an adapter to allow the use of ordinary
// function as Payment mutator.
type PaymentFunc func(context.Context, *ent.PaymentMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PaymentFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PaymentMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PaymentMutation", m)
	}
	return f(ctx, mv)
}

// The PharmacistFunc type is an adapter to allow the use of ordinary
// function as Pharmacist mutator.
type PharmacistFunc func(context.Context, *ent.PharmacistMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PharmacistFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PharmacistMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PharmacistMutation", m)
	}
	return f(ctx, mv)
}

// The PositionInPharmacistFunc type is an adapter to allow the use of ordinary
// function as PositionInPharmacist mutator.
type PositionInPharmacistFunc func(context.Context, *ent.PositionInPharmacistMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PositionInPharmacistFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PositionInPharmacistMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PositionInPharmacistMutation", m)
	}
	return f(ctx, mv)
}

// The PrescriptionFunc type is an adapter to allow the use of ordinary
// function as Prescription mutator.
type PrescriptionFunc func(context.Context, *ent.PrescriptionMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f PrescriptionFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.PrescriptionMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.PrescriptionMutation", m)
	}
	return f(ctx, mv)
}

// The StatusFunc type is an adapter to allow the use of ordinary
// function as Status mutator.
type StatusFunc func(context.Context, *ent.StatusMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f StatusFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.StatusMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.StatusMutation", m)
	}
	return f(ctx, mv)
}

// The UnitOfMedicineFunc type is an adapter to allow the use of ordinary
// function as UnitOfMedicine mutator.
type UnitOfMedicineFunc func(context.Context, *ent.UnitOfMedicineMutation) (ent.Value, error)

// Mutate calls f(ctx, m).
func (f UnitOfMedicineFunc) Mutate(ctx context.Context, m ent.Mutation) (ent.Value, error) {
	mv, ok := m.(*ent.UnitOfMedicineMutation)
	if !ok {
		return nil, fmt.Errorf("unexpected mutation type %T. expect *ent.UnitOfMedicineMutation", m)
	}
	return f(ctx, mv)
}

// Condition is a hook condition function.
type Condition func(context.Context, ent.Mutation) bool

// And groups conditions with the AND operator.
func And(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if !first(ctx, m) || !second(ctx, m) {
			return false
		}
		for _, cond := range rest {
			if !cond(ctx, m) {
				return false
			}
		}
		return true
	}
}

// Or groups conditions with the OR operator.
func Or(first, second Condition, rest ...Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		if first(ctx, m) || second(ctx, m) {
			return true
		}
		for _, cond := range rest {
			if cond(ctx, m) {
				return true
			}
		}
		return false
	}
}

// Not negates a given condition.
func Not(cond Condition) Condition {
	return func(ctx context.Context, m ent.Mutation) bool {
		return !cond(ctx, m)
	}
}

// HasOp is a condition testing mutation operation.
func HasOp(op ent.Op) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		return m.Op().Is(op)
	}
}

// HasAddedFields is a condition validating `.AddedField` on fields.
func HasAddedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.AddedField(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.AddedField(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasClearedFields is a condition validating `.FieldCleared` on fields.
func HasClearedFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if exists := m.FieldCleared(field); !exists {
			return false
		}
		for _, field := range fields {
			if exists := m.FieldCleared(field); !exists {
				return false
			}
		}
		return true
	}
}

// HasFields is a condition validating `.Field` on fields.
func HasFields(field string, fields ...string) Condition {
	return func(_ context.Context, m ent.Mutation) bool {
		if _, exists := m.Field(field); !exists {
			return false
		}
		for _, field := range fields {
			if _, exists := m.Field(field); !exists {
				return false
			}
		}
		return true
	}
}

// If executes the given hook under condition.
//
//	Hook.If(ComputeAverage, And(HasFields(...), HasAddedFields(...)))
//
func If(hk ent.Hook, cond Condition) ent.Hook {
	return func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if cond(ctx, m) {
				return hk(next).Mutate(ctx, m)
			}
			return next.Mutate(ctx, m)
		})
	}
}

// On executes the given hook only for the given operation.
//
//	hook.On(Log, ent.Delete|ent.Create)
//
func On(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, HasOp(op))
}

// Unless skips the given hook only for the given operation.
//
//	hook.Unless(Log, ent.Update|ent.UpdateOne)
//
func Unless(hk ent.Hook, op ent.Op) ent.Hook {
	return If(hk, Not(HasOp(op)))
}

// Reject returns a hook that rejects all operations that match op.
//
//	func (T) Hooks() []ent.Hook {
//		return []ent.Hook{
//			Reject(ent.Delete|ent.Update),
//		}
//	}
//
func Reject(op ent.Op) ent.Hook {
	hk := func(ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(_ context.Context, m ent.Mutation) (ent.Value, error) {
			return nil, fmt.Errorf("%s operation is not allowed", m.Op())
		})
	}
	return On(hk, op)
}

// Chain acts as a list of hooks and is effectively immutable.
// Once created, it will always hold the same set of hooks in the same order.
type Chain struct {
	hooks []ent.Hook
}

// NewChain creates a new chain of hooks.
func NewChain(hooks ...ent.Hook) Chain {
	return Chain{append([]ent.Hook(nil), hooks...)}
}

// Hook chains the list of hooks and returns the final hook.
func (c Chain) Hook() ent.Hook {
	return func(mutator ent.Mutator) ent.Mutator {
		for i := len(c.hooks) - 1; i >= 0; i-- {
			mutator = c.hooks[i](mutator)
		}
		return mutator
	}
}

// Append extends a chain, adding the specified hook
// as the last ones in the mutation flow.
func (c Chain) Append(hooks ...ent.Hook) Chain {
	newHooks := make([]ent.Hook, 0, len(c.hooks)+len(hooks))
	newHooks = append(newHooks, c.hooks...)
	newHooks = append(newHooks, hooks...)
	return Chain{newHooks}
}

// Extend extends a chain, adding the specified chain
// as the last ones in the mutation flow.
func (c Chain) Extend(chain Chain) Chain {
	return c.Append(chain.hooks...)
}
