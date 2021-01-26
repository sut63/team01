// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"github.com/facebookincubator/ent/dialect/sql/schema"
	"github.com/facebookincubator/ent/schema/field"
)

var (
	// AnnotationsColumns holds the columns for the "annotations" table.
	AnnotationsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "messages", Type: field.TypeString},
	}
	// AnnotationsTable holds the schema information for the "annotations" table.
	AnnotationsTable = &schema.Table{
		Name:        "annotations",
		Columns:     AnnotationsColumns,
		PrimaryKey:  []*schema.Column{AnnotationsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// BillsColumns holds the columns for the "bills" table.
	BillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "amount", Type: field.TypeInt},
		{Name: "annotation", Type: field.TypeString},
		{Name: "dispensemedicine_id", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "payment_id", Type: field.TypeInt, Nullable: true},
		{Name: "pharmacist_id", Type: field.TypeInt, Nullable: true},
	}
	// BillsTable holds the schema information for the "bills" table.
	BillsTable = &schema.Table{
		Name:       "bills",
		Columns:    BillsColumns,
		PrimaryKey: []*schema.Column{BillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "bills_dispense_medicines_Bills",
				Columns: []*schema.Column{BillsColumns[3]},

				RefColumns: []*schema.Column{DispenseMedicinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bills_payments_Bills",
				Columns: []*schema.Column{BillsColumns[4]},

				RefColumns: []*schema.Column{PaymentsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "bills_pharmacists_Bills",
				Columns: []*schema.Column{BillsColumns[5]},

				RefColumns: []*schema.Column{PharmacistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// CompaniesColumns holds the columns for the "companies" table.
	CompaniesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// CompaniesTable holds the schema information for the "companies" table.
	CompaniesTable = &schema.Table{
		Name:        "companies",
		Columns:     CompaniesColumns,
		PrimaryKey:  []*schema.Column{CompaniesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DispenseMedicinesColumns holds the columns for the "dispense_medicines" table.
	DispenseMedicinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "datetime", Type: field.TypeTime},
		{Name: "annotation_id", Type: field.TypeInt, Nullable: true},
		{Name: "pharmacist_id", Type: field.TypeInt, Nullable: true},
		{Name: "prescription_id", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// DispenseMedicinesTable holds the schema information for the "dispense_medicines" table.
	DispenseMedicinesTable = &schema.Table{
		Name:       "dispense_medicines",
		Columns:    DispenseMedicinesColumns,
		PrimaryKey: []*schema.Column{DispenseMedicinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "dispense_medicines_annotations_dispensemedicine",
				Columns: []*schema.Column{DispenseMedicinesColumns[2]},

				RefColumns: []*schema.Column{AnnotationsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dispense_medicines_pharmacists_dispensemedicine",
				Columns: []*schema.Column{DispenseMedicinesColumns[3]},

				RefColumns: []*schema.Column{PharmacistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "dispense_medicines_prescriptions_dispensemedicine",
				Columns: []*schema.Column{DispenseMedicinesColumns[4]},

				RefColumns: []*schema.Column{PrescriptionsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DoctorsColumns holds the columns for the "doctors" table.
	DoctorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// DoctorsTable holds the schema information for the "doctors" table.
	DoctorsTable = &schema.Table{
		Name:        "doctors",
		Columns:     DoctorsColumns,
		PrimaryKey:  []*schema.Column{DoctorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// DrugAllergiesColumns holds the columns for the "drug_allergies" table.
	DrugAllergiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "date_time", Type: field.TypeTime},
		{Name: "medicine_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "pharmacist_id", Type: field.TypeInt, Nullable: true},
	}
	// DrugAllergiesTable holds the schema information for the "drug_allergies" table.
	DrugAllergiesTable = &schema.Table{
		Name:       "drug_allergies",
		Columns:    DrugAllergiesColumns,
		PrimaryKey: []*schema.Column{DrugAllergiesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "drug_allergies_medicines_drugallergys",
				Columns: []*schema.Column{DrugAllergiesColumns[2]},

				RefColumns: []*schema.Column{MedicinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drug_allergies_patient_infos_drugallergys",
				Columns: []*schema.Column{DrugAllergiesColumns[3]},

				RefColumns: []*schema.Column{PatientInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "drug_allergies_pharmacists_drugallergys",
				Columns: []*schema.Column{DrugAllergiesColumns[4]},

				RefColumns: []*schema.Column{PharmacistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// LevelOfDangerousColumns holds the columns for the "level_of_dangerous" table.
	LevelOfDangerousColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// LevelOfDangerousTable holds the schema information for the "level_of_dangerous" table.
	LevelOfDangerousTable = &schema.Table{
		Name:        "level_of_dangerous",
		Columns:     LevelOfDangerousColumns,
		PrimaryKey:  []*schema.Column{LevelOfDangerousColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// MedicinesColumns holds the columns for the "medicines" table.
	MedicinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
		{Name: "serial", Type: field.TypeString},
		{Name: "brand", Type: field.TypeString},
		{Name: "amount", Type: field.TypeInt},
		{Name: "price", Type: field.TypeInt},
		{Name: "howtouse", Type: field.TypeString},
		{Name: "level_of_dangerous_medicine", Type: field.TypeInt, Nullable: true},
		{Name: "medicine_type_medicine", Type: field.TypeInt, Nullable: true},
		{Name: "unit_of_medicine_medicine", Type: field.TypeInt, Nullable: true},
	}
	// MedicinesTable holds the schema information for the "medicines" table.
	MedicinesTable = &schema.Table{
		Name:       "medicines",
		Columns:    MedicinesColumns,
		PrimaryKey: []*schema.Column{MedicinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "medicines_level_of_dangerous_Medicine",
				Columns: []*schema.Column{MedicinesColumns[7]},

				RefColumns: []*schema.Column{LevelOfDangerousColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medicines_medicine_types_Medicine",
				Columns: []*schema.Column{MedicinesColumns[8]},

				RefColumns: []*schema.Column{MedicineTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "medicines_unit_of_medicines_Medicine",
				Columns: []*schema.Column{MedicinesColumns[9]},

				RefColumns: []*schema.Column{UnitOfMedicinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MedicineTypesColumns holds the columns for the "medicine_types" table.
	MedicineTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// MedicineTypesTable holds the schema information for the "medicine_types" table.
	MedicineTypesTable = &schema.Table{
		Name:        "medicine_types",
		Columns:     MedicineTypesColumns,
		PrimaryKey:  []*schema.Column{MedicineTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// OrdersColumns holds the columns for the "orders" table.
	OrdersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "addedtime", Type: field.TypeTime},
		{Name: "price", Type: field.TypeInt},
		{Name: "amount", Type: field.TypeInt},
		{Name: "company_ordercompany", Type: field.TypeInt, Nullable: true},
		{Name: "medicine_ordermedicine", Type: field.TypeInt, Nullable: true},
		{Name: "pharmacist_orderpharmacist", Type: field.TypeInt, Nullable: true},
	}
	// OrdersTable holds the schema information for the "orders" table.
	OrdersTable = &schema.Table{
		Name:       "orders",
		Columns:    OrdersColumns,
		PrimaryKey: []*schema.Column{OrdersColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "orders_companies_ordercompany",
				Columns: []*schema.Column{OrdersColumns[4]},

				RefColumns: []*schema.Column{CompaniesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orders_medicines_ordermedicine",
				Columns: []*schema.Column{OrdersColumns[5]},

				RefColumns: []*schema.Column{MedicinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "orders_pharmacists_orderpharmacist",
				Columns: []*schema.Column{OrdersColumns[6]},

				RefColumns: []*schema.Column{PharmacistsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PatientInfosColumns holds the columns for the "patient_infos" table.
	PatientInfosColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "card_number", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "gender", Type: field.TypeString},
		{Name: "age", Type: field.TypeInt},
	}
	// PatientInfosTable holds the schema information for the "patient_infos" table.
	PatientInfosTable = &schema.Table{
		Name:        "patient_infos",
		Columns:     PatientInfosColumns,
		PrimaryKey:  []*schema.Column{PatientInfosColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PaymentsColumns holds the columns for the "payments" table.
	PaymentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "pay", Type: field.TypeString},
	}
	// PaymentsTable holds the schema information for the "payments" table.
	PaymentsTable = &schema.Table{
		Name:        "payments",
		Columns:     PaymentsColumns,
		PrimaryKey:  []*schema.Column{PaymentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PharmacistsColumns holds the columns for the "pharmacists" table.
	PharmacistsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "password", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
	}
	// PharmacistsTable holds the schema information for the "pharmacists" table.
	PharmacistsTable = &schema.Table{
		Name:        "pharmacists",
		Columns:     PharmacistsColumns,
		PrimaryKey:  []*schema.Column{PharmacistsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// PrescriptionsColumns holds the columns for the "prescriptions" table.
	PrescriptionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "value", Type: field.TypeInt},
		{Name: "symptom", Type: field.TypeString},
		{Name: "annotation", Type: field.TypeString},
		{Name: "doctor_id", Type: field.TypeInt, Nullable: true},
		{Name: "medicine_id", Type: field.TypeInt, Nullable: true},
		{Name: "patient_id", Type: field.TypeInt, Nullable: true},
		{Name: "status_id", Type: field.TypeInt, Nullable: true},
	}
	// PrescriptionsTable holds the schema information for the "prescriptions" table.
	PrescriptionsTable = &schema.Table{
		Name:       "prescriptions",
		Columns:    PrescriptionsColumns,
		PrimaryKey: []*schema.Column{PrescriptionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:  "prescriptions_doctors_doctorprescription",
				Columns: []*schema.Column{PrescriptionsColumns[4]},

				RefColumns: []*schema.Column{DoctorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "prescriptions_medicines_medicinepresciption",
				Columns: []*schema.Column{PrescriptionsColumns[5]},

				RefColumns: []*schema.Column{MedicinesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "prescriptions_patient_infos_patientprescription",
				Columns: []*schema.Column{PrescriptionsColumns[6]},

				RefColumns: []*schema.Column{PatientInfosColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:  "prescriptions_status_statusprescription",
				Columns: []*schema.Column{PrescriptionsColumns[7]},

				RefColumns: []*schema.Column{StatusColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// StatusColumns holds the columns for the "status" table.
	StatusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "status", Type: field.TypeString, Unique: true},
	}
	// StatusTable holds the schema information for the "status" table.
	StatusTable = &schema.Table{
		Name:        "status",
		Columns:     StatusColumns,
		PrimaryKey:  []*schema.Column{StatusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// UnitOfMedicinesColumns holds the columns for the "unit_of_medicines" table.
	UnitOfMedicinesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "name", Type: field.TypeString},
	}
	// UnitOfMedicinesTable holds the schema information for the "unit_of_medicines" table.
	UnitOfMedicinesTable = &schema.Table{
		Name:        "unit_of_medicines",
		Columns:     UnitOfMedicinesColumns,
		PrimaryKey:  []*schema.Column{UnitOfMedicinesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AnnotationsTable,
		BillsTable,
		CompaniesTable,
		DispenseMedicinesTable,
		DoctorsTable,
		DrugAllergiesTable,
		LevelOfDangerousTable,
		MedicinesTable,
		MedicineTypesTable,
		OrdersTable,
		PatientInfosTable,
		PaymentsTable,
		PharmacistsTable,
		PrescriptionsTable,
		StatusTable,
		UnitOfMedicinesTable,
	}
)

func init() {
	BillsTable.ForeignKeys[0].RefTable = DispenseMedicinesTable
	BillsTable.ForeignKeys[1].RefTable = PaymentsTable
	BillsTable.ForeignKeys[2].RefTable = PharmacistsTable
	DispenseMedicinesTable.ForeignKeys[0].RefTable = AnnotationsTable
	DispenseMedicinesTable.ForeignKeys[1].RefTable = PharmacistsTable
	DispenseMedicinesTable.ForeignKeys[2].RefTable = PrescriptionsTable
	DrugAllergiesTable.ForeignKeys[0].RefTable = MedicinesTable
	DrugAllergiesTable.ForeignKeys[1].RefTable = PatientInfosTable
	DrugAllergiesTable.ForeignKeys[2].RefTable = PharmacistsTable
	MedicinesTable.ForeignKeys[0].RefTable = LevelOfDangerousTable
	MedicinesTable.ForeignKeys[1].RefTable = MedicineTypesTable
	MedicinesTable.ForeignKeys[2].RefTable = UnitOfMedicinesTable
	OrdersTable.ForeignKeys[0].RefTable = CompaniesTable
	OrdersTable.ForeignKeys[1].RefTable = MedicinesTable
	OrdersTable.ForeignKeys[2].RefTable = PharmacistsTable
	PrescriptionsTable.ForeignKeys[0].RefTable = DoctorsTable
	PrescriptionsTable.ForeignKeys[1].RefTable = MedicinesTable
	PrescriptionsTable.ForeignKeys[2].RefTable = PatientInfosTable
	PrescriptionsTable.ForeignKeys[3].RefTable = StatusTable
}
