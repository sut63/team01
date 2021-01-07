// Code generated by entc, DO NOT EDIT.

package medicine

const (
	// Label holds the string label denoting the medicine type in the database.
	Label = "medicine"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSerial holds the string denoting the serial field in the database.
	FieldSerial = "serial"
	// FieldBrand holds the string denoting the brand field in the database.
	FieldBrand = "brand"
	// FieldAmount holds the string denoting the amount field in the database.
	FieldAmount = "amount"
	// FieldPrice holds the string denoting the price field in the database.
	FieldPrice = "price"
	// FieldHowtouse holds the string denoting the howtouse field in the database.
	FieldHowtouse = "howtouse"

	// EdgeLevelOfDangerous holds the string denoting the levelofdangerous edge name in mutations.
	EdgeLevelOfDangerous = "LevelOfDangerous"
	// EdgeMedicineType holds the string denoting the medicinetype edge name in mutations.
	EdgeMedicineType = "MedicineType"
	// EdgeUnitOfMedicine holds the string denoting the unitofmedicine edge name in mutations.
	EdgeUnitOfMedicine = "UnitOfMedicine"
	// EdgeDrugallergys holds the string denoting the drugallergys edge name in mutations.
	EdgeDrugallergys = "drugallergys"
	// EdgeMedicinepresciption holds the string denoting the medicinepresciption edge name in mutations.
	EdgeMedicinepresciption = "medicinepresciption"
	// EdgeOrder holds the string denoting the order edge name in mutations.
	EdgeOrder = "order"

	// Table holds the table name of the medicine in the database.
	Table = "medicines"
	// LevelOfDangerousTable is the table the holds the LevelOfDangerous relation/edge.
	LevelOfDangerousTable = "medicines"
	// LevelOfDangerousInverseTable is the table name for the LevelOfDangerous entity.
	// It exists in this package in order to avoid circular dependency with the "levelofdangerous" package.
	LevelOfDangerousInverseTable = "level_of_dangerous"
	// LevelOfDangerousColumn is the table column denoting the LevelOfDangerous relation/edge.
	LevelOfDangerousColumn = "level_of_dangerous_medicine"
	// MedicineTypeTable is the table the holds the MedicineType relation/edge.
	MedicineTypeTable = "medicines"
	// MedicineTypeInverseTable is the table name for the MedicineType entity.
	// It exists in this package in order to avoid circular dependency with the "medicinetype" package.
	MedicineTypeInverseTable = "medicine_types"
	// MedicineTypeColumn is the table column denoting the MedicineType relation/edge.
	MedicineTypeColumn = "medicine_type_medicine"
	// UnitOfMedicineTable is the table the holds the UnitOfMedicine relation/edge.
	UnitOfMedicineTable = "medicines"
	// UnitOfMedicineInverseTable is the table name for the UnitOfMedicine entity.
	// It exists in this package in order to avoid circular dependency with the "unitofmedicine" package.
	UnitOfMedicineInverseTable = "unit_of_medicines"
	// UnitOfMedicineColumn is the table column denoting the UnitOfMedicine relation/edge.
	UnitOfMedicineColumn = "unit_of_medicine_medicine"
	// DrugallergysTable is the table the holds the drugallergys relation/edge.
	DrugallergysTable = "drug_allergies"
	// DrugallergysInverseTable is the table name for the DrugAllergy entity.
	// It exists in this package in order to avoid circular dependency with the "drugallergy" package.
	DrugallergysInverseTable = "drug_allergies"
	// DrugallergysColumn is the table column denoting the drugallergys relation/edge.
	DrugallergysColumn = "medicine_id"
	// MedicinepresciptionTable is the table the holds the medicinepresciption relation/edge.
	MedicinepresciptionTable = "prescriptions"
	// MedicinepresciptionInverseTable is the table name for the Prescription entity.
	// It exists in this package in order to avoid circular dependency with the "prescription" package.
	MedicinepresciptionInverseTable = "prescriptions"
	// MedicinepresciptionColumn is the table column denoting the medicinepresciption relation/edge.
	MedicinepresciptionColumn = "medicine_id"
	// OrderTable is the table the holds the order relation/edge. The primary key declared below.
	OrderTable = "medicine_order"
	// OrderInverseTable is the table name for the Order entity.
	// It exists in this package in order to avoid circular dependency with the "order" package.
	OrderInverseTable = "orders"
)

// Columns holds all SQL columns for medicine fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldSerial,
	FieldBrand,
	FieldAmount,
	FieldPrice,
	FieldHowtouse,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the Medicine type.
var ForeignKeys = []string{
	"level_of_dangerous_medicine",
	"medicine_type_medicine",
	"unit_of_medicine_medicine",
}

var (
	// OrderPrimaryKey and OrderColumn2 are the table columns denoting the
	// primary key for the order relation (M2M).
	OrderPrimaryKey = []string{"medicine_id", "order_id"}
)