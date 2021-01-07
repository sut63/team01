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
	_ "github.com/sut63/team01/docs" //ไม่ต้องลบมันจะหายแดงหลังจาก swag
	"github.com/sut63/team01/ent"
)

/*type Users struct {
	User []User
}

type User struct {
	Name  string
	Email string
}

type Playlists struct {
	Playlist []Playlist
}

type Playlist struct {
	Title string
	Owner int
}

type Videos struct {
	Video []Video
}

type Video struct {
	Name  string
	Url   string
	Owner int
}*/

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

//DispenseMedicines Structure
type DispenseMedicines struct {
	DispenseMedicine []DispenseMedicine
}

//DispenseMedicine Structure
type DispenseMedicine struct {
	Datetime     string
	Prescription int
	Annotation   int
	Pharmacist   int
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

	/*controllers.NewUserController(v1, client)
	controllers.NewVideoController(v1, client)
	controllers.NewResolutionController(v1, client)
	controllers.NewPlaylistController(v1, client)
	controllers.NewPlaylistVideoController(v1, client)*/

	controllers.NewAnnotationController(v1, client)
	controllers.NewPharmacistController(v1, client)
	controllers.NewDispenseMedicineController(v1, client)

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

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run()
}
