// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/levelofdangerous"
)

// LevelOfDangerous is the model entity for the LevelOfDangerous schema.
type LevelOfDangerous struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LevelOfDangerousQuery when eager-loading is set.
	Edges LevelOfDangerousEdges `json:"edges"`
}

// LevelOfDangerousEdges holds the relations/edges for other nodes in the graph.
type LevelOfDangerousEdges struct {
	// Medicine holds the value of the Medicine edge.
	Medicine []*Medicine
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MedicineOrErr returns the Medicine value or an error if the edge
// was not loaded in eager-loading.
func (e LevelOfDangerousEdges) MedicineOrErr() ([]*Medicine, error) {
	if e.loadedTypes[0] {
		return e.Medicine, nil
	}
	return nil, &NotLoadedError{edge: "Medicine"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LevelOfDangerous) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // name
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LevelOfDangerous fields.
func (lod *LevelOfDangerous) assignValues(values ...interface{}) error {
	if m, n := len(values), len(levelofdangerous.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	lod.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[0])
	} else if value.Valid {
		lod.Name = value.String
	}
	return nil
}

// QueryMedicine queries the Medicine edge of the LevelOfDangerous.
func (lod *LevelOfDangerous) QueryMedicine() *MedicineQuery {
	return (&LevelOfDangerousClient{config: lod.config}).QueryMedicine(lod)
}

// Update returns a builder for updating this LevelOfDangerous.
// Note that, you need to call LevelOfDangerous.Unwrap() before calling this method, if this LevelOfDangerous
// was returned from a transaction, and the transaction was committed or rolled back.
func (lod *LevelOfDangerous) Update() *LevelOfDangerousUpdateOne {
	return (&LevelOfDangerousClient{config: lod.config}).UpdateOne(lod)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (lod *LevelOfDangerous) Unwrap() *LevelOfDangerous {
	tx, ok := lod.config.driver.(*txDriver)
	if !ok {
		panic("ent: LevelOfDangerous is not a transactional entity")
	}
	lod.config.driver = tx.drv
	return lod
}

// String implements the fmt.Stringer.
func (lod *LevelOfDangerous) String() string {
	var builder strings.Builder
	builder.WriteString("LevelOfDangerous(")
	builder.WriteString(fmt.Sprintf("id=%v", lod.ID))
	builder.WriteString(", name=")
	builder.WriteString(lod.Name)
	builder.WriteByte(')')
	return builder.String()
}

// LevelOfDangerousSlice is a parsable slice of LevelOfDangerous.
type LevelOfDangerousSlice []*LevelOfDangerous

func (lod LevelOfDangerousSlice) config(cfg config) {
	for _i := range lod {
		lod[_i].config = cfg
	}
}
