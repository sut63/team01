// Code generated by entc, DO NOT EDIT.

package dispensemedicine

const (
	// Label holds the string label denoting the dispensemedicine type in the database.
	Label = "dispense_medicine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldDatetime holds the string denoting the datetime field in the database.
	FieldDatetime = "datetime"
	// FieldNote holds the string denoting the note field in the database.
	FieldNote = "note"
	// FieldAmountchangemedicine holds the string denoting the amountchangemedicine field in the database.
	FieldAmountchangemedicine = "amountchangemedicine"
	// FieldDetailchangemedicine holds the string denoting the detailchangemedicine field in the database.
	FieldDetailchangemedicine = "detailchangemedicine"

	// EdgePharmacist holds the string denoting the pharmacist edge name in mutations.
	EdgePharmacist = "pharmacist"
	// EdgeAnnotation holds the string denoting the annotation edge name in mutations.
	EdgeAnnotation = "annotation"
	// EdgePrescription holds the string denoting the prescription edge name in mutations.
	EdgePrescription = "prescription"
	// EdgeBills holds the string denoting the bills edge name in mutations.
	EdgeBills = "Bills"

	// Table holds the table name of the dispensemedicine in the database.
	Table = "dispense_medicines"
	// PharmacistTable is the table the holds the pharmacist relation/edge.
	PharmacistTable = "dispense_medicines"
	// PharmacistInverseTable is the table name for the Pharmacist entity.
	// It exists in this package in order to avoid circular dependency with the "pharmacist" package.
	PharmacistInverseTable = "pharmacists"
	// PharmacistColumn is the table column denoting the pharmacist relation/edge.
	PharmacistColumn = "pharmacist_id"
	// AnnotationTable is the table the holds the annotation relation/edge.
	AnnotationTable = "dispense_medicines"
	// AnnotationInverseTable is the table name for the Annotation entity.
	// It exists in this package in order to avoid circular dependency with the "annotation" package.
	AnnotationInverseTable = "annotations"
	// AnnotationColumn is the table column denoting the annotation relation/edge.
	AnnotationColumn = "annotation_id"
	// PrescriptionTable is the table the holds the prescription relation/edge.
	PrescriptionTable = "dispense_medicines"
	// PrescriptionInverseTable is the table name for the Prescription entity.
	// It exists in this package in order to avoid circular dependency with the "prescription" package.
	PrescriptionInverseTable = "prescriptions"
	// PrescriptionColumn is the table column denoting the prescription relation/edge.
	PrescriptionColumn = "prescription_id"
	// BillsTable is the table the holds the Bills relation/edge.
	BillsTable = "bills"
	// BillsInverseTable is the table name for the Bill entity.
	// It exists in this package in order to avoid circular dependency with the "bill" package.
	BillsInverseTable = "bills"
	// BillsColumn is the table column denoting the Bills relation/edge.
	BillsColumn = "dispensemedicine_id"
)

// Columns holds all SQL columns for dispensemedicine fields.
var Columns = []string{
	FieldID,
	FieldDatetime,
	FieldNote,
	FieldAmountchangemedicine,
	FieldDetailchangemedicine,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the DispenseMedicine type.
var ForeignKeys = []string{
	"annotation_id",
	"pharmacist_id",
	"prescription_id",
}

var (
	// NoteValidator is a validator for the "note" field. It is called by the builders before save.
	NoteValidator func(string) error
	// AmountchangemedicineValidator is a validator for the "amountchangemedicine" field. It is called by the builders before save.
	AmountchangemedicineValidator func(int) error
	// DetailchangemedicineValidator is a validator for the "detailchangemedicine" field. It is called by the builders before save.
	DetailchangemedicineValidator func(string) error
)
