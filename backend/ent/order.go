// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/company"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/order"
	"github.com/sut63/team01/ent/pharmacist"
)

// Order is the model entity for the Order schema.
type Order struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Hospitalid holds the value of the "hospitalid" field.
	Hospitalid string `json:"hospitalid,omitempty"`
	// Addedtime holds the value of the "addedtime" field.
	Addedtime time.Time `json:"addedtime,omitempty"`
	// Price holds the value of the "price" field.
	Price int `json:"price,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the OrderQuery when eager-loading is set.
	Edges                      OrderEdges `json:"edges"`
	company_ordercompany       *int
	medicine_ordermedicine     *int
	pharmacist_orderpharmacist *int
}

// OrderEdges holds the relations/edges for other nodes in the graph.
type OrderEdges struct {
	// Medicine holds the value of the medicine edge.
	Medicine *Medicine
	// Company holds the value of the company edge.
	Company *Company
	// Pharmacist holds the value of the pharmacist edge.
	Pharmacist *Pharmacist
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// MedicineOrErr returns the Medicine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) MedicineOrErr() (*Medicine, error) {
	if e.loadedTypes[0] {
		if e.Medicine == nil {
			// The edge medicine was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicine.Label}
		}
		return e.Medicine, nil
	}
	return nil, &NotLoadedError{edge: "medicine"}
}

// CompanyOrErr returns the Company value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) CompanyOrErr() (*Company, error) {
	if e.loadedTypes[1] {
		if e.Company == nil {
			// The edge company was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: company.Label}
		}
		return e.Company, nil
	}
	return nil, &NotLoadedError{edge: "company"}
}

// PharmacistOrErr returns the Pharmacist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e OrderEdges) PharmacistOrErr() (*Pharmacist, error) {
	if e.loadedTypes[2] {
		if e.Pharmacist == nil {
			// The edge pharmacist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pharmacist.Label}
		}
		return e.Pharmacist, nil
	}
	return nil, &NotLoadedError{edge: "pharmacist"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Order) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // hospitalid
		&sql.NullTime{},   // addedtime
		&sql.NullInt64{},  // price
		&sql.NullInt64{},  // amount
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Order) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // company_ordercompany
		&sql.NullInt64{}, // medicine_ordermedicine
		&sql.NullInt64{}, // pharmacist_orderpharmacist
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Order fields.
func (o *Order) assignValues(values ...interface{}) error {
	if m, n := len(values), len(order.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	o.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field hospitalid", values[0])
	} else if value.Valid {
		o.Hospitalid = value.String
	}
	if value, ok := values[1].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field addedtime", values[1])
	} else if value.Valid {
		o.Addedtime = value.Time
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field price", values[2])
	} else if value.Valid {
		o.Price = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field amount", values[3])
	} else if value.Valid {
		o.Amount = int(value.Int64)
	}
	values = values[4:]
	if len(values) == len(order.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field company_ordercompany", value)
		} else if value.Valid {
			o.company_ordercompany = new(int)
			*o.company_ordercompany = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field medicine_ordermedicine", value)
		} else if value.Valid {
			o.medicine_ordermedicine = new(int)
			*o.medicine_ordermedicine = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field pharmacist_orderpharmacist", value)
		} else if value.Valid {
			o.pharmacist_orderpharmacist = new(int)
			*o.pharmacist_orderpharmacist = int(value.Int64)
		}
	}
	return nil
}

// QueryMedicine queries the medicine edge of the Order.
func (o *Order) QueryMedicine() *MedicineQuery {
	return (&OrderClient{config: o.config}).QueryMedicine(o)
}

// QueryCompany queries the company edge of the Order.
func (o *Order) QueryCompany() *CompanyQuery {
	return (&OrderClient{config: o.config}).QueryCompany(o)
}

// QueryPharmacist queries the pharmacist edge of the Order.
func (o *Order) QueryPharmacist() *PharmacistQuery {
	return (&OrderClient{config: o.config}).QueryPharmacist(o)
}

// Update returns a builder for updating this Order.
// Note that, you need to call Order.Unwrap() before calling this method, if this Order
// was returned from a transaction, and the transaction was committed or rolled back.
func (o *Order) Update() *OrderUpdateOne {
	return (&OrderClient{config: o.config}).UpdateOne(o)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (o *Order) Unwrap() *Order {
	tx, ok := o.config.driver.(*txDriver)
	if !ok {
		panic("ent: Order is not a transactional entity")
	}
	o.config.driver = tx.drv
	return o
}

// String implements the fmt.Stringer.
func (o *Order) String() string {
	var builder strings.Builder
	builder.WriteString("Order(")
	builder.WriteString(fmt.Sprintf("id=%v", o.ID))
	builder.WriteString(", hospitalid=")
	builder.WriteString(o.Hospitalid)
	builder.WriteString(", addedtime=")
	builder.WriteString(o.Addedtime.Format(time.ANSIC))
	builder.WriteString(", price=")
	builder.WriteString(fmt.Sprintf("%v", o.Price))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", o.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// Orders is a parsable slice of Order.
type Orders []*Order

func (o Orders) config(cfg config) {
	for _i := range o {
		o[_i].config = cfg
	}
}
