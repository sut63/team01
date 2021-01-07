// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/payment"
)

// Payment is the model entity for the Payment schema.
type Payment struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Pay holds the value of the "pay" field.
	Pay string `json:"pay,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PaymentQuery when eager-loading is set.
	Edges PaymentEdges `json:"edges"`
}

// PaymentEdges holds the relations/edges for other nodes in the graph.
type PaymentEdges struct {
	// Bills holds the value of the Bills edge.
	Bills []*Bill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// BillsOrErr returns the Bills value or an error if the edge
// was not loaded in eager-loading.
func (e PaymentEdges) BillsOrErr() ([]*Bill, error) {
	if e.loadedTypes[0] {
		return e.Bills, nil
	}
	return nil, &NotLoadedError{edge: "Bills"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Payment) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // pay
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Payment fields.
func (pa *Payment) assignValues(values ...interface{}) error {
	if m, n := len(values), len(payment.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pa.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field pay", values[0])
	} else if value.Valid {
		pa.Pay = value.String
	}
	return nil
}

// QueryBills queries the Bills edge of the Payment.
func (pa *Payment) QueryBills() *BillQuery {
	return (&PaymentClient{config: pa.config}).QueryBills(pa)
}

// Update returns a builder for updating this Payment.
// Note that, you need to call Payment.Unwrap() before calling this method, if this Payment
// was returned from a transaction, and the transaction was committed or rolled back.
func (pa *Payment) Update() *PaymentUpdateOne {
	return (&PaymentClient{config: pa.config}).UpdateOne(pa)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pa *Payment) Unwrap() *Payment {
	tx, ok := pa.config.driver.(*txDriver)
	if !ok {
		panic("ent: Payment is not a transactional entity")
	}
	pa.config.driver = tx.drv
	return pa
}

// String implements the fmt.Stringer.
func (pa *Payment) String() string {
	var builder strings.Builder
	builder.WriteString("Payment(")
	builder.WriteString(fmt.Sprintf("id=%v", pa.ID))
	builder.WriteString(", pay=")
	builder.WriteString(pa.Pay)
	builder.WriteByte(')')
	return builder.String()
}

// Payments is a parsable slice of Payment.
type Payments []*Payment

func (pa Payments) config(cfg config) {
	for _i := range pa {
		pa[_i].config = cfg
	}
}
