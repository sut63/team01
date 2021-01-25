// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/bill"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/payment"
	"github.com/sut63/team01/ent/pharmacist"
)

// Bill is the model entity for the Bill schema.
type Bill struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// Annotation holds the value of the "annotation" field.
	Annotation string `json:"annotation,omitempty"`
	// Payer holds the value of the "payer" field.
	Payer string `json:"payer,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BillQuery when eager-loading is set.
	Edges               BillEdges `json:"edges"`
	dispensemedicine_id *int
	payment_id          *int
	pharmacist_id       *int
}

// BillEdges holds the relations/edges for other nodes in the graph.
type BillEdges struct {
	// Pharmacists holds the value of the pharmacists edge.
	Pharmacists *Pharmacist
	// Payments holds the value of the payments edge.
	Payments *Payment
	// DispenseMedicines holds the value of the dispenseMedicines edge.
	DispenseMedicines *DispenseMedicine
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// PharmacistsOrErr returns the Pharmacists value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) PharmacistsOrErr() (*Pharmacist, error) {
	if e.loadedTypes[0] {
		if e.Pharmacists == nil {
			// The edge pharmacists was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pharmacist.Label}
		}
		return e.Pharmacists, nil
	}
	return nil, &NotLoadedError{edge: "pharmacists"}
}

// PaymentsOrErr returns the Payments value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) PaymentsOrErr() (*Payment, error) {
	if e.loadedTypes[1] {
		if e.Payments == nil {
			// The edge payments was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: payment.Label}
		}
		return e.Payments, nil
	}
	return nil, &NotLoadedError{edge: "payments"}
}

// DispenseMedicinesOrErr returns the DispenseMedicines value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BillEdges) DispenseMedicinesOrErr() (*DispenseMedicine, error) {
	if e.loadedTypes[2] {
		if e.DispenseMedicines == nil {
			// The edge dispenseMedicines was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dispensemedicine.Label}
		}
		return e.DispenseMedicines, nil
	}
	return nil, &NotLoadedError{edge: "dispenseMedicines"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Bill) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // amount
		&sql.NullString{}, // annotation
		&sql.NullString{}, // payer
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Bill) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // dispensemedicine_id
		&sql.NullInt64{}, // payment_id
		&sql.NullInt64{}, // pharmacist_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Bill fields.
func (b *Bill) assignValues(values ...interface{}) error {
	if m, n := len(values), len(bill.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	b.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field amount", values[0])
	} else if value.Valid {
		b.Amount = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field annotation", values[1])
	} else if value.Valid {
		b.Annotation = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field payer", values[2])
	} else if value.Valid {
		b.Payer = value.String
	}
	values = values[3:]
	if len(values) == len(bill.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field dispensemedicine_id", value)
		} else if value.Valid {
			b.dispensemedicine_id = new(int)
			*b.dispensemedicine_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field payment_id", value)
		} else if value.Valid {
			b.payment_id = new(int)
			*b.payment_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field pharmacist_id", value)
		} else if value.Valid {
			b.pharmacist_id = new(int)
			*b.pharmacist_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPharmacists queries the pharmacists edge of the Bill.
func (b *Bill) QueryPharmacists() *PharmacistQuery {
	return (&BillClient{config: b.config}).QueryPharmacists(b)
}

// QueryPayments queries the payments edge of the Bill.
func (b *Bill) QueryPayments() *PaymentQuery {
	return (&BillClient{config: b.config}).QueryPayments(b)
}

// QueryDispenseMedicines queries the dispenseMedicines edge of the Bill.
func (b *Bill) QueryDispenseMedicines() *DispenseMedicineQuery {
	return (&BillClient{config: b.config}).QueryDispenseMedicines(b)
}

// Update returns a builder for updating this Bill.
// Note that, you need to call Bill.Unwrap() before calling this method, if this Bill
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Bill) Update() *BillUpdateOne {
	return (&BillClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (b *Bill) Unwrap() *Bill {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Bill is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Bill) String() string {
	var builder strings.Builder
	builder.WriteString("Bill(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", b.Amount))
	builder.WriteString(", annotation=")
	builder.WriteString(b.Annotation)
	builder.WriteString(", payer=")
	builder.WriteString(b.Payer)
	builder.WriteByte(')')
	return builder.String()
}

// Bills is a parsable slice of Bill.
type Bills []*Bill

func (b Bills) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
