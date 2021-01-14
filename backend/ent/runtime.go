// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/sut63/team01/ent/annotation"
	"github.com/sut63/team01/ent/bill"
	"github.com/sut63/team01/ent/company"
	"github.com/sut63/team01/ent/doctor"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/prescription"
	"github.com/sut63/team01/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	annotationFields := schema.Annotation{}.Fields()
	_ = annotationFields
	// annotationDescMessages is the schema descriptor for messages field.
	annotationDescMessages := annotationFields[0].Descriptor()
	// annotation.MessagesValidator is a validator for the "messages" field. It is called by the builders before save.
	annotation.MessagesValidator = annotationDescMessages.Validators[0].(func(string) error)
	billFields := schema.Bill{}.Fields()
	_ = billFields
	// billDescAmount is the schema descriptor for amount field.
	billDescAmount := billFields[0].Descriptor()
	// bill.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	bill.AmountValidator = billDescAmount.Validators[0].(func(int) error)
	companyFields := schema.Company{}.Fields()
	_ = companyFields
	// companyDescName is the schema descriptor for Name field.
	companyDescName := companyFields[0].Descriptor()
	// company.NameValidator is a validator for the "Name" field. It is called by the builders before save.
	company.NameValidator = companyDescName.Validators[0].(func(string) error)
	doctorFields := schema.Doctor{}.Fields()
	_ = doctorFields
	// doctorDescEmail is the schema descriptor for email field.
	doctorDescEmail := doctorFields[0].Descriptor()
	// doctor.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	doctor.EmailValidator = doctorDescEmail.Validators[0].(func(string) error)
	// doctorDescPassword is the schema descriptor for password field.
	doctorDescPassword := doctorFields[1].Descriptor()
	// doctor.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	doctor.PasswordValidator = doctorDescPassword.Validators[0].(func(string) error)
	// doctorDescName is the schema descriptor for name field.
	doctorDescName := doctorFields[2].Descriptor()
	// doctor.NameValidator is a validator for the "name" field. It is called by the builders before save.
	doctor.NameValidator = doctorDescName.Validators[0].(func(string) error)
	medicineFields := schema.Medicine{}.Fields()
	_ = medicineFields
	// medicineDescName is the schema descriptor for name field.
	medicineDescName := medicineFields[0].Descriptor()
	// medicine.NameValidator is a validator for the "name" field. It is called by the builders before save.
	medicine.NameValidator = medicineDescName.Validators[0].(func(string) error)
	// medicineDescSerial is the schema descriptor for serial field.
	medicineDescSerial := medicineFields[1].Descriptor()
	// medicine.SerialValidator is a validator for the "serial" field. It is called by the builders before save.
	medicine.SerialValidator = medicineDescSerial.Validators[0].(func(string) error)
	// medicineDescBrand is the schema descriptor for brand field.
	medicineDescBrand := medicineFields[2].Descriptor()
	// medicine.BrandValidator is a validator for the "brand" field. It is called by the builders before save.
	medicine.BrandValidator = medicineDescBrand.Validators[0].(func(string) error)
	// medicineDescAmount is the schema descriptor for amount field.
	medicineDescAmount := medicineFields[3].Descriptor()
	// medicine.AmountValidator is a validator for the "amount" field. It is called by the builders before save.
	medicine.AmountValidator = medicineDescAmount.Validators[0].(func(int) error)
	// medicineDescPrice is the schema descriptor for price field.
	medicineDescPrice := medicineFields[4].Descriptor()
	// medicine.PriceValidator is a validator for the "price" field. It is called by the builders before save.
	medicine.PriceValidator = medicineDescPrice.Validators[0].(func(int) error)
	// medicineDescHowtouse is the schema descriptor for howtouse field.
	medicineDescHowtouse := medicineFields[5].Descriptor()
	// medicine.HowtouseValidator is a validator for the "howtouse" field. It is called by the builders before save.
	medicine.HowtouseValidator = medicineDescHowtouse.Validators[0].(func(string) error)
	patientinfoFields := schema.PatientInfo{}.Fields()
	_ = patientinfoFields
	// patientinfoDescCardNumber is the schema descriptor for cardNumber field.
	patientinfoDescCardNumber := patientinfoFields[0].Descriptor()
	// patientinfo.CardNumberValidator is a validator for the "cardNumber" field. It is called by the builders before save.
	patientinfo.CardNumberValidator = patientinfoDescCardNumber.Validators[0].(func(string) error)
	// patientinfoDescName is the schema descriptor for name field.
	patientinfoDescName := patientinfoFields[1].Descriptor()
	// patientinfo.NameValidator is a validator for the "name" field. It is called by the builders before save.
	patientinfo.NameValidator = patientinfoDescName.Validators[0].(func(string) error)
	// patientinfoDescGender is the schema descriptor for gender field.
	patientinfoDescGender := patientinfoFields[2].Descriptor()
	// patientinfo.GenderValidator is a validator for the "gender" field. It is called by the builders before save.
	patientinfo.GenderValidator = patientinfoDescGender.Validators[0].(func(string) error)
	// patientinfoDescAge is the schema descriptor for age field.
	patientinfoDescAge := patientinfoFields[3].Descriptor()
	// patientinfo.AgeValidator is a validator for the "age" field. It is called by the builders before save.
	patientinfo.AgeValidator = patientinfoDescAge.Validators[0].(func(int) error)
	pharmacistFields := schema.Pharmacist{}.Fields()
	_ = pharmacistFields
	// pharmacistDescEmail is the schema descriptor for email field.
	pharmacistDescEmail := pharmacistFields[0].Descriptor()
	// pharmacist.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	pharmacist.EmailValidator = pharmacistDescEmail.Validators[0].(func(string) error)
	// pharmacistDescPassword is the schema descriptor for password field.
	pharmacistDescPassword := pharmacistFields[1].Descriptor()
	// pharmacist.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	pharmacist.PasswordValidator = pharmacistDescPassword.Validators[0].(func(string) error)
	// pharmacistDescName is the schema descriptor for name field.
	pharmacistDescName := pharmacistFields[2].Descriptor()
	// pharmacist.NameValidator is a validator for the "name" field. It is called by the builders before save.
	pharmacist.NameValidator = pharmacistDescName.Validators[0].(func(string) error)
	prescriptionFields := schema.Prescription{}.Fields()
	_ = prescriptionFields
	// prescriptionDescValue is the schema descriptor for Value field.
	prescriptionDescValue := prescriptionFields[0].Descriptor()
	// prescription.ValueValidator is a validator for the "Value" field. It is called by the builders before save.
	prescription.ValueValidator = prescriptionDescValue.Validators[0].(func(int) error)
}
