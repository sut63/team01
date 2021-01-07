// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/patientinfo"
)

// PatientInfo is the model entity for the PatientInfo schema.
type PatientInfo struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CardNumber holds the value of the "cardNumber" field.
	CardNumber string `json:"cardNumber,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PatientInfoQuery when eager-loading is set.
	Edges PatientInfoEdges `json:"edges"`
}

// PatientInfoEdges holds the relations/edges for other nodes in the graph.
type PatientInfoEdges struct {
	// Drugallergys holds the value of the drugallergys edge.
	Drugallergys []*DrugAllergy
	// Patientprescription holds the value of the patientprescription edge.
	Patientprescription []*Prescription
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// DrugallergysOrErr returns the Drugallergys value or an error if the edge
// was not loaded in eager-loading.
func (e PatientInfoEdges) DrugallergysOrErr() ([]*DrugAllergy, error) {
	if e.loadedTypes[0] {
		return e.Drugallergys, nil
	}
	return nil, &NotLoadedError{edge: "drugallergys"}
}

// PatientprescriptionOrErr returns the Patientprescription value or an error if the edge
// was not loaded in eager-loading.
func (e PatientInfoEdges) PatientprescriptionOrErr() ([]*Prescription, error) {
	if e.loadedTypes[1] {
		return e.Patientprescription, nil
	}
	return nil, &NotLoadedError{edge: "patientprescription"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PatientInfo) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // cardNumber
		&sql.NullString{}, // name
		&sql.NullString{}, // gender
		&sql.NullInt64{},  // age
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PatientInfo fields.
func (pi *PatientInfo) assignValues(values ...interface{}) error {
	if m, n := len(values), len(patientinfo.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pi.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field cardNumber", values[0])
	} else if value.Valid {
		pi.CardNumber = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field name", values[1])
	} else if value.Valid {
		pi.Name = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field gender", values[2])
	} else if value.Valid {
		pi.Gender = value.String
	}
	if value, ok := values[3].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field age", values[3])
	} else if value.Valid {
		pi.Age = int(value.Int64)
	}
	return nil
}

// QueryDrugallergys queries the drugallergys edge of the PatientInfo.
func (pi *PatientInfo) QueryDrugallergys() *DrugAllergyQuery {
	return (&PatientInfoClient{config: pi.config}).QueryDrugallergys(pi)
}

// QueryPatientprescription queries the patientprescription edge of the PatientInfo.
func (pi *PatientInfo) QueryPatientprescription() *PrescriptionQuery {
	return (&PatientInfoClient{config: pi.config}).QueryPatientprescription(pi)
}

// Update returns a builder for updating this PatientInfo.
// Note that, you need to call PatientInfo.Unwrap() before calling this method, if this PatientInfo
// was returned from a transaction, and the transaction was committed or rolled back.
func (pi *PatientInfo) Update() *PatientInfoUpdateOne {
	return (&PatientInfoClient{config: pi.config}).UpdateOne(pi)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pi *PatientInfo) Unwrap() *PatientInfo {
	tx, ok := pi.config.driver.(*txDriver)
	if !ok {
		panic("ent: PatientInfo is not a transactional entity")
	}
	pi.config.driver = tx.drv
	return pi
}

// String implements the fmt.Stringer.
func (pi *PatientInfo) String() string {
	var builder strings.Builder
	builder.WriteString("PatientInfo(")
	builder.WriteString(fmt.Sprintf("id=%v", pi.ID))
	builder.WriteString(", cardNumber=")
	builder.WriteString(pi.CardNumber)
	builder.WriteString(", name=")
	builder.WriteString(pi.Name)
	builder.WriteString(", gender=")
	builder.WriteString(pi.Gender)
	builder.WriteString(", age=")
	builder.WriteString(fmt.Sprintf("%v", pi.Age))
	builder.WriteByte(')')
	return builder.String()
}

// PatientInfos is a parsable slice of PatientInfo.
type PatientInfos []*PatientInfo

func (pi PatientInfos) config(cfg config) {
	for _i := range pi {
		pi[_i].config = cfg
	}
}