// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/doctor"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/prescription"
)

// Prescription is the model entity for the Prescription schema.
type Prescription struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Value holds the value of the "Value" field.
	Value int `json:"Value,omitempty"`
	// Symptom holds the value of the "Symptom" field.
	Symptom string `json:"Symptom,omitempty"`
	// Annotation holds the value of the "Annotation" field.
	Annotation string `json:"Annotation,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the PrescriptionQuery when eager-loading is set.
	Edges       PrescriptionEdges `json:"edges"`
	doctor_id   *int
	medicine_id *int
	patient_id  *int
}

// PrescriptionEdges holds the relations/edges for other nodes in the graph.
type PrescriptionEdges struct {
	// Prescriptionpatient holds the value of the prescriptionpatient edge.
	Prescriptionpatient *PatientInfo
	// Prescriptiondoctor holds the value of the prescriptiondoctor edge.
	Prescriptiondoctor *Doctor
	// Prescriptionmedicine holds the value of the prescriptionmedicine edge.
	Prescriptionmedicine *Medicine
	// Dispensemedicine holds the value of the dispensemedicine edge.
	Dispensemedicine *DispenseMedicine
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [4]bool
}

// PrescriptionpatientOrErr returns the Prescriptionpatient value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) PrescriptionpatientOrErr() (*PatientInfo, error) {
	if e.loadedTypes[0] {
		if e.Prescriptionpatient == nil {
			// The edge prescriptionpatient was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: patientinfo.Label}
		}
		return e.Prescriptionpatient, nil
	}
	return nil, &NotLoadedError{edge: "prescriptionpatient"}
}

// PrescriptiondoctorOrErr returns the Prescriptiondoctor value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) PrescriptiondoctorOrErr() (*Doctor, error) {
	if e.loadedTypes[1] {
		if e.Prescriptiondoctor == nil {
			// The edge prescriptiondoctor was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: doctor.Label}
		}
		return e.Prescriptiondoctor, nil
	}
	return nil, &NotLoadedError{edge: "prescriptiondoctor"}
}

// PrescriptionmedicineOrErr returns the Prescriptionmedicine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) PrescriptionmedicineOrErr() (*Medicine, error) {
	if e.loadedTypes[2] {
		if e.Prescriptionmedicine == nil {
			// The edge prescriptionmedicine was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: medicine.Label}
		}
		return e.Prescriptionmedicine, nil
	}
	return nil, &NotLoadedError{edge: "prescriptionmedicine"}
}

// DispensemedicineOrErr returns the Dispensemedicine value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e PrescriptionEdges) DispensemedicineOrErr() (*DispenseMedicine, error) {
	if e.loadedTypes[3] {
		if e.Dispensemedicine == nil {
			// The edge dispensemedicine was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: dispensemedicine.Label}
		}
		return e.Dispensemedicine, nil
	}
	return nil, &NotLoadedError{edge: "dispensemedicine"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Prescription) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullInt64{},  // Value
		&sql.NullString{}, // Symptom
		&sql.NullString{}, // Annotation
	}
}

// fkValues returns the types for scanning foreign-keys values from sql.Rows.
func (*Prescription) fkValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{}, // doctor_id
		&sql.NullInt64{}, // medicine_id
		&sql.NullInt64{}, // patient_id
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Prescription fields.
func (pr *Prescription) assignValues(values ...interface{}) error {
	if m, n := len(values), len(prescription.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	pr.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullInt64); !ok {
		return fmt.Errorf("unexpected type %T for field Value", values[0])
	} else if value.Valid {
		pr.Value = int(value.Int64)
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Symptom", values[1])
	} else if value.Valid {
		pr.Symptom = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Annotation", values[2])
	} else if value.Valid {
		pr.Annotation = value.String
	}
	values = values[3:]
	if len(values) == len(prescription.ForeignKeys) {
		if value, ok := values[0].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field doctor_id", value)
		} else if value.Valid {
			pr.doctor_id = new(int)
			*pr.doctor_id = int(value.Int64)
		}
		if value, ok := values[1].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field medicine_id", value)
		} else if value.Valid {
			pr.medicine_id = new(int)
			*pr.medicine_id = int(value.Int64)
		}
		if value, ok := values[2].(*sql.NullInt64); !ok {
			return fmt.Errorf("unexpected type %T for edge-field patient_id", value)
		} else if value.Valid {
			pr.patient_id = new(int)
			*pr.patient_id = int(value.Int64)
		}
	}
	return nil
}

// QueryPrescriptionpatient queries the prescriptionpatient edge of the Prescription.
func (pr *Prescription) QueryPrescriptionpatient() *PatientInfoQuery {
	return (&PrescriptionClient{config: pr.config}).QueryPrescriptionpatient(pr)
}

// QueryPrescriptiondoctor queries the prescriptiondoctor edge of the Prescription.
func (pr *Prescription) QueryPrescriptiondoctor() *DoctorQuery {
	return (&PrescriptionClient{config: pr.config}).QueryPrescriptiondoctor(pr)
}

// QueryPrescriptionmedicine queries the prescriptionmedicine edge of the Prescription.
func (pr *Prescription) QueryPrescriptionmedicine() *MedicineQuery {
	return (&PrescriptionClient{config: pr.config}).QueryPrescriptionmedicine(pr)
}

// QueryDispensemedicine queries the dispensemedicine edge of the Prescription.
func (pr *Prescription) QueryDispensemedicine() *DispenseMedicineQuery {
	return (&PrescriptionClient{config: pr.config}).QueryDispensemedicine(pr)
}

// Update returns a builder for updating this Prescription.
// Note that, you need to call Prescription.Unwrap() before calling this method, if this Prescription
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Prescription) Update() *PrescriptionUpdateOne {
	return (&PrescriptionClient{config: pr.config}).UpdateOne(pr)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (pr *Prescription) Unwrap() *Prescription {
	tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Prescription is not a transactional entity")
	}
	pr.config.driver = tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Prescription) String() string {
	var builder strings.Builder
	builder.WriteString("Prescription(")
	builder.WriteString(fmt.Sprintf("id=%v", pr.ID))
	builder.WriteString(", Value=")
	builder.WriteString(fmt.Sprintf("%v", pr.Value))
	builder.WriteString(", Symptom=")
	builder.WriteString(pr.Symptom)
	builder.WriteString(", Annotation=")
	builder.WriteString(pr.Annotation)
	builder.WriteByte(')')
	return builder.String()
}

// Prescriptions is a parsable slice of Prescription.
type Prescriptions []*Prescription

func (pr Prescriptions) config(cfg config) {
	for _i := range pr {
		pr[_i].config = cfg
	}
}
