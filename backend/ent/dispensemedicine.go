// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/annotation"
	"github.com/sut63/team01/ent/bill"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/prescription"
)

// DispenseMedicine is the model entity for the DispenseMedicine schema.
type DispenseMedicine struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Datetime holds the value of the "datetime" field.
	Datetime time.Time `json:"datetime,omitempty"`
	// Note holds the value of the "note" field.
	Note string `json:"note,omitempty"`
	// Amountchangemedicine holds the value of the "amountchangemedicine" field.
	Amountchangemedicine int `json:"amountchangemedicine,omitempty"`
	// Detailchangemedicine holds the value of the "detailchangemedicine" field.
	Detailchangemedicine string `json:"detailchangemedicine,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the DispenseMedicineQuery when eager-loading is set.
	Edges           DispenseMedicineEdges `json:"edges"`
	annotation_id   *int
	pharmacist_id   *int
	prescription_id *int
}

// DispenseMedicineEdges holds the relations/edges for other nodes in the graph.
type DispenseMedicineEdges struct {
	// Pharmacist holds the value of the pharmacist edge.
	Pharmacist *Pharmacist
	// Annotation holds the value of the annotation edge.
	Annotation *Annotation
	// Prescription holds the value of the prescription edge.
	Prescription *Prescription
	// Bills holds the value of the Bills edge.
	Bills *Bill
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PharmacistOrErr returns the Pharmacist value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DispenseMedicineEdges) PharmacistOrErr() (*Pharmacist, error) {
	if e.loadedTypes[0] {
		if e.Pharmacist == nil {
			// The edge pharmacist was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: pharmacist.Label}
		}
		return e.Pharmacist, nil
	}
	return nil, &NotLoadedError{edge: "pharmacist"}
}

// AnnotationOrErr returns the Annotation value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DispenseMedicineEdges) AnnotationOrErr() (*Annotation, error) {
	if e.loadedTypes[1] {
		if e.Annotation == nil {
			// The edge annotation was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: annotation.Label}
		}
		return e.Annotation, nil
	}
	return nil, &NotLoadedError{edge: "annotation"}
}

// PrescriptionOrErr returns the Prescription value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DispenseMedicineEdges) PrescriptionOrErr() (*Prescription, error) {
	if e.loadedTypes[2] {
		if e.Prescription == nil {
			// The edge prescription was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: prescription.Label}
		}
		return e.Prescription, nil
	}
	return nil, &NotLoadedError{edge: "prescription"}
}

// BillsOrErr returns the Bills value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e DispenseMedicineEdges) BillsOrErr() (*Bill, error) {
	if e.loadedTypes[3] {
		if e.Bills == nil {
			// The edge Bills was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: bill.Label}
		}
		return e.Bills, nil
	}
	return nil, &NotLoadedError{edge: "Bills"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*DispenseMedicine) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullTime{},   // datetime
		&sql.NullString{}, // note
		&sql.NullInt64{},  // amountchangemedicine
		&sql.NullString{}, // detailchangemedicine
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*DispenseMedicine) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // annotation_id
		&sql.NullInt64{}, // pharmacist_id
		&sql.NullInt64{}, // prescription_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the DispenseMedicine fields.
func (dm *DispenseMedicine) assignValues(values ...interface{}) error {
	if m, n := len(values), len(dispensemedicine.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	dm.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullTime); !ok {
		return fmt.Errorf("unexpected type %T for field datetime", values[0])
	} else if value.Valid {
		dm.Datetime = value.Time
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field note", values[1])
	} else if value.Valid {
		dm.Note = value.String
	}
	if value, ok := values[2].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field amountchangemedicine", values[2])
	} else if value.Valid {
		dm.Amountchangemedicine = int(value.Int64)
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field detailchangemedicine", values[3])
	} else if value.Valid {
		dm.Detailchangemedicine = value.String
	}
	values = values[4:]
	if len(values) == len(dispensemedicine.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field annotation_id", value)
		} else if value.Valid {
			dm.annotation_id = new(int)
			*dm.annotation_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field pharmacist_id", value)
		} else if value.Valid {
			dm.pharmacist_id = new(int)
			*dm.pharmacist_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field prescription_id", value)
		} else if value.Valid {
			dm.prescription_id = new(int)
			*dm.prescription_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPharmacist queries the pharmacist edge of the DispenseMedicine.
func (dm *DispenseMedicine) QueryPharmacist() *PharmacistQuery {
	return (&DispenseMedicineClient{config: dm.config}).QueryPharmacist(dm)
}

// QueryAnnotation queries the annotation edge of the DispenseMedicine.
func (dm *DispenseMedicine) QueryAnnotation() *AnnotationQuery {
	return (&DispenseMedicineClient{config: dm.config}).QueryAnnotation(dm)
}

// QueryPrescription queries the prescription edge of the DispenseMedicine.
func (dm *DispenseMedicine) QueryPrescription() *PrescriptionQuery {
	return (&DispenseMedicineClient{config: dm.config}).QueryPrescription(dm)
}

// QueryBills queries the Bills edge of the DispenseMedicine.
func (dm *DispenseMedicine) QueryBills() *BillQuery {
	return (&DispenseMedicineClient{config: dm.config}).QueryBills(dm)
}

// Update returns a builder for updating this DispenseMedicine.
// Note that, you need to call DispenseMedicine.Unwrap() before calling this method, if this DispenseMedicine
// was returned from a transaction, and the transaction was committed or rolled back.
func (dm *DispenseMedicine) Update() *DispenseMedicineUpdateOne {
	return (&DispenseMedicineClient{config: dm.config}).UpdateOne(dm)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (dm *DispenseMedicine) Unwrap() *DispenseMedicine {
	tx, ok := dm.config.driver.(*txDriver)
	if !ok {
		panic("ent: DispenseMedicine is not a transactional entity")
	}
	dm.config.driver = tx.drv
	return dm
}

// String implements the fmt.Stringer.
func (dm *DispenseMedicine) String() string {
	var builder strings.Builder
	builder.WriteString("DispenseMedicine(")
	builder.WriteString(fmt.Sprintf("id=%v", dm.ID))
	builder.WriteString(", datetime=")
	builder.WriteString(dm.Datetime.Format(time.ANSIC))
	builder.WriteString(", note=")
	builder.WriteString(dm.Note)
	builder.WriteString(", amountchangemedicine=")
	builder.WriteString(fmt.Sprintf("%v", dm.Amountchangemedicine))
	builder.WriteString(", detailchangemedicine=")
	builder.WriteString(dm.Detailchangemedicine)
	builder.WriteByte(')')
	return builder.String()
}

// DispenseMedicines is a parsable slice of DispenseMedicine.
type DispenseMedicines []*DispenseMedicine

func (dm DispenseMedicines) config(cfg config) {
	for _i := range dm {
		dm[_i].config = cfg
	}
}
