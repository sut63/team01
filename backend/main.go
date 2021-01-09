package main

import (
	"context"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/sut63/team01/controllers"
	_ "github.com/sut63/team01/docs"
	"github.com/sut63/team01/ent"
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

//Pharmacists Structure
type Pharmacists struct {
	Pharmacist []Pharmacist
}

//Pharmacist Structure
type Pharmacist struct {
	Email    string
	Password string
	Name     string
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

// LevelOfDangerouss structer input data
type LevelOfDangerouss struct {
	LevelOfDangerous []LevelOfDangerous
}

// LevelOfDangerous structer input data
type LevelOfDangerous struct {
	Name string
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
	controllers.NewDispenseMedicineController(v1, client)
	controllers.NewDrugAllergyController(v1, client)
	controllers.NewPatientInfoController(v1, client)
	controllers.NewDoctorController(v1, client)
	controllers.NewPrescriptionController(v1, client)
	controllers.NewMedicineTypeController(v1, client)
	controllers.NewUnitOfMedicineController(v1, client)
	controllers.NewLevelOfDangerousController(v1, client)
	controllers.NewMedicineController(v1, client)

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

	//Set Pharmacists Data
	Pharma := Pharmacists{
		Pharmacist: []Pharmacist{
			Pharmacist{"codetib926@hotmail.com", "code926", "Pharm. Nacodetip Hanchai"},
			Pharmacist{"anna1231@gmail.com", "anna1231", "Pharm. Anna Saithai"},
			Pharmacist{"manisara@gmail.com", "mani098765", "Pharm. Manisara Insuwan"},
			Pharmacist{"somchai@hotmail.com", "s096666", "Pharm. Somchai Poonsuk"},
			Pharmacist{"supanan5009@gmail.com", "Su14520", "Pharm. Supanan Pongsuwan"},
			Pharmacist{"kuntarit5010@gmail.com", "Kun@0881234", "Pharm. Kuntarit Wannasak"},
		},
	}
	for _, phar := range Pharma.Pharmacist {
		client.Pharmacist.
			Create().
			SetEmail(phar.Email).
			SetPassword(phar.Password).
			SetName(phar.Name).
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

	// Set LevelOfDangerouss Data
	LevelOfDangerouss := LevelOfDangerouss{
		LevelOfDangerous: []LevelOfDangerous{
			LevelOfDangerous{"Safe"},
			LevelOfDangerous{"Caution"},
			LevelOfDangerous{"Danger"},
		},
	}

	for _, l := range LevelOfDangerouss.LevelOfDangerous {
		client.LevelOfDangerous.
			Create().
			SetName(l.Name).
			Save(context.Background())
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
