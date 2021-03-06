package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/sut63/team01/controllers"
	_ "github.com/sut63/team01/docs"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/annotation"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/doctor"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/payment"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/positioninpharmacist"
	"github.com/sut63/team01/ent/prescription"
)

//PatientInfos Structure
type PatientInfos struct {
	PatientInfo []PatientInfo
}

//PatientInfo Structure
type PatientInfo struct {
	CardNumber string
	Name       string
	Gender     string
	Age        int
}

//Annotations Structure
type Annotations struct {
	Annotation []Annotation
}

//Annotation Structure
type Annotation struct {
	Messages string
}

//Payments Structure
type Payments struct {
	Payment []Payment
}

//Payment Structure
type Payment struct {
	Pay string
}

//PositionInPharmacists Structure
type PositionInPharmacists struct {
	PositionInPharmacist []PositionInPharmacist
}

//PositionInPharmacist Structure
type PositionInPharmacist struct {
	Position string
}

//Pharmacists Structure
type Pharmacists struct {
	Pharmacist []Pharmacist
}

//Pharmacist Structure
type Pharmacist struct {
	Email                string
	Password             string
	Name                 string
	PositionInPharmacist int
}

//DispenseMedicines Structure
type DispenseMedicines struct {
	DispenseMedicine []DispenseMedicine
}

//DispenseMedicine Structure
type DispenseMedicine struct {
	Datetime             string
	Note                 string
	Amountchangemedicine int
	Detailchangemedicine string
	Prescription         int
	Annotation           int
	Pharmacist           int
}

//Prescriptions Strucre
type Prescriptions struct {
	Prescription []Prescription
}

//Prescription Strucre
type Prescription struct {
	DoctorID      int
	PatientInfoID int
	MedicineID    int
	Value         int
	Symptom       string
	Annotation    string
}

//Doctors Structure
type Doctors struct {
	Doctor []Doctor
}

//Doctor Structure
type Doctor struct {
	email    string
	password string
	name     string
}

// MedicineTypes structer input data
type MedicineTypes struct {
	MedicineType []MedicineType
}

// MedicineType structer input data
type MedicineType struct {
	Name string
}

// UnitOfMedicines structer input data
type UnitOfMedicines struct {
	UnitOfMedicine []UnitOfMedicine
}

// UnitOfMedicine structer input data
type UnitOfMedicine struct {
	Name string
}

// LevelOfDangeropress structer input data
type LevelOfDangeropress struct {
	LevelOfDangeropres []LevelOfDangeropres
}

// LevelOfDangeropres structer input data
type LevelOfDangeropres struct {
	Name string
}

//Bills Structure
type Bills struct {
	Bill []Bill
}

//Bill Structure
type Bill struct {
	Annotation       string
	Amount           int
	Payer            string
	Payments         int
	DispenseMedicine int
	Pharmacists      int
}

//Medicines Structure
type Medicines struct {
	Medicine []Medicine
}

//Medicine Structure
type Medicine struct {
	Name               string
	Serial             string
	Brand              string
	Amount             int
	Price              int
	Howtoprese         string
	LevelOfDangeropres int
	MedicineType       int
	UnitOfMedicine     int
}

//Companys Structure
type Companys struct {
	Company []Company
}

//Company Structure
type Company struct {
	name string
}

// @title SUT SA Example API Playlist Vidoe
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
	router := gin.Default()
	router.Use(cors.Default())

	client, err := ent.Open("sqlite3", "file:Ent-Database.db?cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("fail to open sqlite3: %v", err)
	}
	defer client.Close()

	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	v1 := router.Group("/api/v1")

	controllers.NewAnnotationController(v1, client)
	controllers.NewPharmacistController(v1, client)
	controllers.NewPositionInPharmacistController(v1, client)
	controllers.NewDispenseMedicineController(v1, client)
	controllers.NewDrugAllergyController(v1, client)
	controllers.NewPatientInfoController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewPrescriptionController(v1, client)
	controllers.NewMedicineTypeController(v1, client)
	controllers.NewUnitOfMedicineController(v1, client)
	controllers.NewLevelOfDangerousController(v1, client)
	controllers.NewMedicineController(v1, client)
	controllers.NewPaymentController(v1, client)
	controllers.NewBillController(v1, client)
	controllers.NewOrderController(v1, client)
	controllers.NewCompanyController(v1, client)

	//Set PatientInfos Data
	Patient := PatientInfos{
		PatientInfo: []PatientInfo{
			PatientInfo{"1720800134909", "Mr. Thanat Sukantatoon", "Male", 21},
			PatientInfo{"1720800125787", "Mr. Somchi Jaidee", "Male", 56},
			PatientInfo{"1920922478655", "Mrs. Somying GingGongKaw", "Female", 30},
			PatientInfo{"1578320214555", "Mr. Jaiped Jedsee", "Male", 80},
			PatientInfo{"1731231232562", "Mrs. Somporn Sornsin", "Female", 45},
			PatientInfo{"1896325211111", "Mr. Pongsak Samakkee", "Male", 99},
		},
	}
	for _, p := range Patient.PatientInfo {
		client.PatientInfo.
			Create().
			SetCardNumber(p.CardNumber).
			SetName(p.Name).
			SetGender(p.Gender).
			SetAge(p.Age).
			Save(context.Background())
	}

	//Set Annotations Data
	Annota := Annotations{
		Annotation: []Annotation{
			Annotation{"ไม่ระบุ"},
			Annotation{"เปลี่ยนตัวยาทดแทน ยี่ห้อเดิม"},
			Annotation{"เปลี่ยนตัวยาทดแทน ยี่ห้อใหม่"},
		},
	}
	for _, anno := range Annota.Annotation {
		client.Annotation.
			Create().
			SetMessages(anno.Messages).
			Save(context.Background())
	}

	//Set Payments Data
	Paym := Payments{
		Payment: []Payment{
			Payment{"การชำระผ่านธนาคารออนไลน์"},
			Payment{"การชำระผ่านบัตรเครดิต"},
			Payment{"การชำระผ่านเงินสด"},
		},
	}
	for _, Paym := range Paym.Payment {
		client.Payment.
			Create().
			SetPay(Paym.Pay).
			Save(context.Background())
	}

	//Set PositionInPharmacists Data
	PositPha := PositionInPharmacists{
		PositionInPharmacist: []PositionInPharmacist{
			PositionInPharmacist{"DrugAllergy"},
			PositionInPharmacist{"Order"},
			PositionInPharmacist{"Medicine"},
			PositionInPharmacist{"DispenseMedicine"},
			PositionInPharmacist{"Bill"},
		},
	}
	for _, poinphar := range PositPha.PositionInPharmacist {
		client.PositionInPharmacist.
			Create().
			SetPosition(poinphar.Position).
			Save(context.Background())
	}

	//Set Pharmacists Data
	Pharma := Pharmacists{
		Pharmacist: []Pharmacist{
			Pharmacist{"code@gmail.com", "c1234", "Pharm. Nacodetip Hanchai", 4},
			Pharmacist{"anna1231@gmail.com", "anna1231", "Pharm. Anna Saithai", 3},
			Pharmacist{"pay@gmail.com", "p1234", "Pharm. Manisara Insuwan", 1},
			Pharmacist{"pat@hotmail.com", "pa1234", "Pharm. Somchai Poonsuk", 5},
			Pharmacist{"dorn@gmail.com", "d1234", "Pharm. Supanan Pongsuwan", 2},
			Pharmacist{"kuntarit5010@gmail.com", "Kun@0881234", "Pharm. Kuntarit Wannasak", 1},
		},
	}
	for _, phar := range Pharma.Pharmacist {

		posinphar, err := client.PositionInPharmacist.
			Query().
			Where(positioninpharmacist.IDEQ(int(phar.PositionInPharmacist))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Pharmacist.
			Create().
			SetEmail(phar.Email).
			SetPassword(phar.Password).
			SetName(phar.Name).
			SetPositioninpharmacist(posinphar).
			Save(context.Background())
	}

	//Set Doctor Data
	Doct := Doctors{
		Doctor: []Doctor{
			Doctor{"jennie01@gmail.com", "je1111", "Ph.D. Jennie Kim"},
			Doctor{"jisoo02@gmail.com", "ji2222", "Ph.D. Kim Chisu"},
			Doctor{"rose03@gmail.com", "ro3333", "Ph.D. Roseanne Park"},
			Doctor{"lisa04@gmail.com", "li4444", "Ph.D. Lalisa Manoban"},
		},
	}
	for _, doc := range Doct.Doctor {
		client.Doctor.
			Create().
			SetEmail(doc.email).
			SetPassword(doc.password).
			SetName(doc.name).
			Save(context.Background())
	}

	// Set MedicineTypes Data
	MedicineTypes := MedicineTypes{
		MedicineType: []MedicineType{
			MedicineType{"General Sales List"},
			MedicineType{"Pharmacy Medicines"},
			MedicineType{"Precription Only Medicines"},
			MedicineType{"Controlled Drugs"},
		},
	}

	for _, t := range MedicineTypes.MedicineType {
		client.MedicineType.
			Create().
			SetName(t.Name).
			Save(context.Background())
	}

	// Set UnitOfMedcines Data
	UnitOfMedicines := UnitOfMedicines{
		UnitOfMedicine: []UnitOfMedicine{
			UnitOfMedicine{"Tablet"},
			UnitOfMedicine{"Bottom"},
			UnitOfMedicine{"Pill"},
			UnitOfMedicine{"Capsule"},
		},
	}

	for _, u := range UnitOfMedicines.UnitOfMedicine {
		client.UnitOfMedicine.
			Create().
			SetName(u.Name).
			Save(context.Background())
	}

	// Set LevelOfDangeropress Data
	LevelOfDangeropress := LevelOfDangeropress{
		LevelOfDangeropres: []LevelOfDangeropres{
			LevelOfDangeropres{"Safe"},
			LevelOfDangeropres{"Caution"},
			LevelOfDangeropres{"Danger"},
		},
	}

	for _, l := range LevelOfDangeropress.LevelOfDangeropres {
		client.LevelOfDangerous.
			Create().
			SetName(l.Name).
			Save(context.Background())
	}

	// Set Medicines Data
	Medicines := Medicines{
		Medicine: []Medicine{
			Medicine{"TestMedicine", "T000", "AGC Inc.", 200000, 300, "prese to create another medicine", 2, 3, 2},
			Medicine{"PainKiller", "P001", "AGC Inc.", 150000, 20, "prese to cure your pain", 1, 1, 3},
		},
	}

	for _, m := range Medicines.Medicine {
		client.Medicine.
			Create().
			SetName(m.Name).
			SetSerial(m.Serial).
			SetBrand(m.Brand).
			SetAmount(m.Amount).
			SetPrice(m.Price).
			SetHowtouse(m.Howtoprese).
			SetLevelOfDangerousID(m.LevelOfDangeropres).
			SetMedicineTypeID(m.MedicineType).
			SetUnitOfMedicineID(m.UnitOfMedicine).
			Save(context.Background())
	}

	//Set Prescription Data
	Pres := Prescriptions{
		Prescription: []Prescription{
			Prescription{1, 1, 1, 15, "ปวดหัว", "ไม่มี"},
		},
	}

	for _, pci := range Pres.Prescription {

		pif, err := client.PatientInfo.
			Query().
			Where(patientinfo.IDEQ(int(pci.PatientInfoID))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dtr, err := client.Doctor.
			Query().
			Where(doctor.IDEQ(int(pci.DoctorID))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		mdc, err := client.Medicine.
			Query().
			Where(medicine.IDEQ(int(pci.MedicineID))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Prescription.
			Create().
			SetValue(pci.Value).
			SetSymptom(pci.Symptom).
			SetAnnotation(pci.Annotation).
			SetPrescriptionpatient(pif).
			SetPrescriptiondoctor(dtr).
			SetPrescriptionmedicine(mdc).
			Save(context.Background())
	}

	//Set DispenseMedicine Data
	DisMedic := DispenseMedicines{
		DispenseMedicine: []DispenseMedicine{
			DispenseMedicine{"2021-01-12T01:27:50.52+07:00", "ไม่มี", 0, "-", 1, 1, 1},
		},
	}

	for _, dim := range DisMedic.DispenseMedicine {

		pres, err := client.Prescription.
			Query().
			Where(prescription.IDEQ(int(dim.Prescription))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		Annota, err := client.Annotation.
			Query().
			Where(annotation.IDEQ(int(dim.Annotation))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		pml, err := client.Pharmacist.
			Query().
			Where(pharmacist.IDEQ(int(dim.Pharmacist))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		datetim, err := time.Parse(time.RFC3339, dim.Datetime)

		client.DispenseMedicine.
			Create().
			SetDatetime(datetim).
			SetNote(dim.Note).
			SetAmountchangemedicine(int(dim.Amountchangemedicine)).
			SetDetailchangemedicine(dim.Detailchangemedicine).
			SetPrescription(pres).
			SetAnnotation(Annota).
			SetPharmacist(pml).
			Save(context.Background())
	}

	// Set Bills Data
	Bills := Bills{
		Bill: []Bill{
			Bill{"-", 100, "Patarapong Chareongpol", 1, 1, 1},
		},
	}

	for _, b := range Bills.Bill {

		py, err := client.Payment.
			Query().
			Where(payment.IDEQ(int(b.Payments))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		dis, err := client.DispenseMedicine.
			Query().
			Where(dispensemedicine.IDEQ(int(b.DispenseMedicine))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		ph, err := client.Pharmacist.
			Query().
			Where(pharmacist.IDEQ(int(b.Pharmacists))).
			Only(context.Background())
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		client.Bill.
			Create().
			SetAmount(b.Amount).
			SetDispenseMedicines(dis).
			SetPayments(py).
			SetAnnotation(b.Annotation).
			SetPharmacists(ph).
			SetPayer(b.Payer).
			Save(context.Background())
	}

	companys := Companys{
		Company: []Company{
			Company{"MedicCompany"},
			Company{"BandageCompany"},
			Company{"ABCompany"},
			Company{"Team1Company"},
		},
	}

	for _, pay := range companys.Company {
		client.Company.
			Create().
			SetName(pay.name).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
