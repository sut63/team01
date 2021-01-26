// Code generated by entc, DO NOT EDIT.

package drugallergy

const (
	// Label holds the string label denoting the drugallergy type in the database.
	Label = "drug_allergy"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldSymptom holds the string denoting the symptom field in the database.
	FieldSymptom = "symptom"
	// FieldCongenitalDisease holds the string denoting the congenitaldisease field in the database.
	FieldCongenitalDisease = "congenital_disease"
	// FieldAnnotation holds the string denoting the annotation field in the database.
	FieldAnnotation = "annotation"
	// FieldDateTime holds the string denoting the datetime field in the database.
	FieldDateTime = "date_time"

	// EdgePatient holds the string denoting the patient edge name in mutations.
	EdgePatient = "patient"
	// EdgeMedicine holds the string denoting the medicine edge name in mutations.
	EdgeMedicine = "medicine"
	// EdgePharmacist holds the string denoting the pharmacist edge name in mutations.
	EdgePharmacist = "pharmacist"

	// Table holds the table name of the drugallergy in the database.
	Table = "drug_allergies"
	// PatientTable is the table the holds the patient relation/edge.
	PatientTable = "drug_allergies"
	// PatientInverseTable is the table name for the PatientInfo entity.
	// It exists in this package in order to avoid circular dependency with the "patientinfo" package.
	PatientInverseTable = "patient_infos"
	// PatientColumn is the table column denoting the patient relation/edge.
	PatientColumn = "patient_id"
	// MedicineTable is the table the holds the medicine relation/edge.
	MedicineTable = "drug_allergies"
	// MedicineInverseTable is the table name for the Medicine entity.
	// It exists in this package in order to avoid circular dependency with the "medicine" package.
	MedicineInverseTable = "medicines"
	// MedicineColumn is the table column denoting the medicine relation/edge.
	MedicineColumn = "medicine_id"
	// PharmacistTable is the table the holds the pharmacist relation/edge.
	PharmacistTable = "drug_allergies"
	// PharmacistInverseTable is the table name for the Pharmacist entity.
	// It exists in this package in order to avoid circular dependency with the "pharmacist" package.
	PharmacistInverseTable = "pharmacists"
	// PharmacistColumn is the table column denoting the pharmacist relation/edge.
	PharmacistColumn = "pharmacist_id"
)

// Columns holds all SQL columns for drugallergy fields.
var Columns = []string{
	FieldID,
	FieldSymptom,
	FieldCongenitalDisease,
	FieldAnnotation,
	FieldDateTime,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the DrugAllergy type.
var ForeignKeys = []string{
	"medicine_id",
	"patient_id",
	"pharmacist_id",
}

var (
	// SymptomValidator is a validator for the "symptom" field. It is called by the builders before save.
	SymptomValidator func(string) error
	// CongenitalDiseaseValidator is a validator for the "congenitalDisease" field. It is called by the builders before save.
	CongenitalDiseaseValidator func(string) error
	// AnnotationValidator is a validator for the "annotation" field. It is called by the builders before save.
	AnnotationValidator func(string) error
)
