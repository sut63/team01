// Code generated by entc, DO NOT EDIT.

package unitofmedicine

const (
	// Label holds the string label denoting the unitofmedicine type in the database.
	Label = "unit_of_medicine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"

	// EdgeMedicine holds the string denoting the medicine edge name in mutations.
	EdgeMedicine = "Medicine"

	// Table holds the table name of the unitofmedicine in the database.
	Table = "unit_of_medicines"
	// MedicineTable is the table the holds the Medicine relation/edge.
	MedicineTable = "medicines"
	// MedicineInverseTable is the table name for the Medicine entity.
	// It exists in this package in order to avoid circular dependency with the "medicine" package.
	MedicineInverseTable = "medicines"
	// MedicineColumn is the table column denoting the Medicine relation/edge.
	MedicineColumn = "unit_of_medicine_medicine"
)

// Columns holds all SQL columns for unitofmedicine fields.
var Columns = []string{
	FieldID,
	FieldName,
}
