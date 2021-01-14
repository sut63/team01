package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/levelofdangerous"
	"github.com/sut63/team01/ent/medicinetype"
	"github.com/sut63/team01/ent/unitofmedicine"
)

// MedicineController defines the struct for the Medicine controller
type MedicineController struct {
	client *ent.Client
	router gin.IRouter
}

// Medicine defines the struct for the Medicine
type Medicine struct {
	Name               string
	Serial             string
	Brand              string
	Amount             int
	Price              int
	Howtouse           string
	MedicineTypeID     int
	LevelOfDangerousID int
	UnitOfMedicineID   int
}

// CreateMedicine handles POST requests for adding Medicine entities
// @Summary Create Medicine
// @Description Create Medicine
// @ID create-Medicine
// @Accept   json
// @Produce  json
// @Param Medicine body Medicine true "Medicine entity"
// @Success 200 {object} ent.Medicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Medicine [post]
func (ctl *MedicineController) CreateMedicine(c *gin.Context) {
	obj := Medicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Medicine binding failed",
		})
		return
	}

	T, err := ctl.client.MedicineType.
		Query().
		Where(medicinetype.IDEQ(int(obj.MedicineTypeID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "province not found",
		})
		return
	}

	L, err := ctl.client.LevelOfDangerous.
		Query().
		Where(levelofdangerous.IDEQ(int(obj.LevelOfDangerousID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "type not found",
		})
		return
	}

	U, err := ctl.client.UnitOfMedicine.
		Query().
		Where(unitofmedicine.IDEQ(int(obj.UnitOfMedicineID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "district not found",
		})
		return
	}

	m, err := ctl.client.Medicine.
		Create().
		SetName(obj.Name).
		SetSerial(obj.Serial).
		SetBrand(obj.Brand).
		SetAmount(obj.Amount).
		SetPrice(obj.Price).
		SetHowtouse(obj.Howtouse).
		SetMedicineType(T).
		SetLevelOfDangerous(L).
		SetUnitOfMedicine(U).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, gin.H{
		"status": true,
		"data":   m,
	})
}

// ListMedicine handles request to get a list of Medicine entities
// @Summary List Medicine entities
// @Description list Medicine entities
// @ID list-Medicine
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Medicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Medicine [get]
func (ctl *MedicineController) ListMedicine(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}
	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}
	Medicine, err := ctl.client.Medicine.
		Query().
		WithMedicineType().
		WithUnitOfMedicine().
		WithLevelOfDangerous().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, Medicine)
}

// NewMedicineController creates and registers handles for the Medicine controller
func NewMedicineController(router gin.IRouter, client *ent.Client) *MedicineController {
	uc := &MedicineController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitMedicineController registers routes to the main engine
func (ctl *MedicineController) register() {
	Medicine := ctl.router.Group("/Medicine")

	Medicine.GET("", ctl.ListMedicine)

	// CRUD
	Medicine.POST("", ctl.CreateMedicine)

}
