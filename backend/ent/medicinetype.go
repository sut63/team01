// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/medicinetype"
)

// MedicineType is the model entity for the MedicineType schema.
type MedicineType struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the MedicineTypeQuery when eager-loading is set.
	Edges MedicineTypeEdges `json:"edges"`
}

// MedicineTypeEdges holds the relations/edges for other nodes in the graph.
type MedicineTypeEdges struct {
	// Medicine holds the value of the Medicine edge.
	Medicine []*Medicine
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MedicineOrErr returns the Medicine value or an error if the edge
// was not loaded in eager-loading.
func (e MedicineTypeEdges) MedicineOrErr() ([]*Medicine, error) {
	if e.loadedTypes[0] {
		return e.Medicine, nil
	}
	return nil, &NotLoadedError{edge: "Medicine"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*MedicineType) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the MedicineType fields.
func (mt *MedicineType) assignValues(values ...interface{}) error {
	if m, n := len(values), len(medicinetype.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	mt.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		mt.Name = value.String
	}
	return nil
}

// QueryMedicine queries the Medicine edge of the MedicineType.
func (mt *MedicineType) QueryMedicine() *MedicineQuery {
	return (&MedicineTypeClient{config: mt.config}).QueryMedicine(mt)
}

// Update returns a builder for updating this MedicineType.
// Note that, you need to call MedicineType.Unwrap() before calling this method, if this MedicineType
// was returned from a transaction, and the transaction was committed or rolled back.
func (mt *MedicineType) Update() *MedicineTypeUpdateOne {
	return (&MedicineTypeClient{config: mt.config}).UpdateOne(mt)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (mt *MedicineType) Unwrap() *MedicineType {
	tx, ok := mt.config.driver.(*txDriver)
	if !ok {
		panic("ent: MedicineType is not a transactional entity")
	}
	mt.config.driver = tx.drv
	return mt
}

// String implements the fmt.Stringer.
func (mt *MedicineType) String() string {
	var builder strings.Builder
	builder.WriteString("MedicineType(")
	builder.WriteString(fmt.Sprintf("id=%v", mt.ID))
	builder.WriteString(", name=")
	builder.WriteString(mt.Name)
	builder.WriteByte(')')
	return builder.String()
}

// MedicineTypes is a parsable slice of MedicineType.
type MedicineTypes []*MedicineType

func (mt MedicineTypes) config(cfg config) {
	for _i := range mt {
		mt[_i].config = cfg
	}
}
