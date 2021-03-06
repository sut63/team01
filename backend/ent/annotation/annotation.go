// Code generated by entc, DO NOT EDIT.

package annotation

const (
	// Label holds the string label denoting the annotation type in the database.
	Label = "annotation"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldMessages holds the string denoting the messages field in the database.
	FieldMessages = "messages"

	// EdgeDispensemedicine holds the string denoting the dispensemedicine edge name in mutations.
	EdgeDispensemedicine = "dispensemedicine"

	// Table holds the table name of the annotation in the database.
	Table = "annotations"
	// DispensemedicineTable is the table the holds the dispensemedicine relation/edge.
	DispensemedicineTable = "dispense_medicines"
	// DispensemedicineInverseTable is the table name for the DispenseMedicine entity.
	// It exists in this package in order to avoid circular dependency with the "dispensemedicine" package.
	DispensemedicineInverseTable = "dispense_medicines"
	// DispensemedicineColumn is the table column denoting the dispensemedicine relation/edge.
	DispensemedicineColumn = "annotation_id"
)

// Columns holds all SQL columns for annotation fields.
var Columns = []string{
	FieldID,
	FieldMessages,
}

var (
	// MessagesValidator is a validator for the "messages" field. It is called by the builders before save.
	MessagesValidator func(string) error
)
