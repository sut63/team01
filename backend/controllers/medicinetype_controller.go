package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/medicinetype"
)

// MedicineTypeController defines the struct for the MedicineType controller
type MedicineTypeController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateMedicineType handles POST requests for adding MedicineType entities
// @Summary Create MedicineType
// @Description Create MedicineType
// @ID create-MedicineType
// @Accept   json
// @Produce  json
// @Param MedicineType body ent.MedicineType true "MedicineType entity"
// @Success 200 {object} ent.MedicineType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /MedicineType [post]
func (ctl *MedicineTypeController) CreateMedicineType(c *gin.Context) {
	obj := ent.MedicineType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "MedicineType binding failed",
		})
		return
	}
	u, err := ctl.client.MedicineType.
		Create().
		SetName(obj.Name).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, u)
}

// GetMedicineType handles GET requests to retrieve a MedicineType entity
// @Summary Get a MedicineType entity by ID
// @Description get MedicineType by ID
// @ID get-MedicineType
// @Produce  json
// @Param id path int true "MedicineType ID"
// @Success 200 {object} ent.MedicineType
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /MedicineType/{id} [get]
func (ctl *MedicineTypeController) GetMedicineType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.MedicineType.
		Query().
		Where(medicinetype.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, u)
}

// ListMedicineType handles request to get a list of MedicineType entities
// @Summary List MedicineType entities
// @Description list MedicineType entities
// @ID list-MedicineType
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.MedicineType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /MedicineType [get]
func (ctl *MedicineTypeController) ListMedicineType(c *gin.Context) {
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
	MedicineType, err := ctl.client.MedicineType.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, MedicineType)
}

// DeleteMedicineType handles DELETE requests to delete a MedicineType entity
// @Summary Delete a MedicineType entity by ID
// @Description get MedicineType by ID
// @ID delete-MedicineType
// @Produce  json
// @Param id path int true "MedicineType ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /MedicineType/{id} [delete]
func (ctl *MedicineTypeController) DeleteMedicineType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.MedicineType.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdateMedicineType handles PUT requests to update a MedicineType entity
// @Summary Update a MedicineType entity by ID
// @Description update MedicineType by ID
// @ID update-MedicineType
// @Accept   json
// @Produce  json
// @Param id path int true "MedicineType ID"
// @Param MedicineType body ent.MedicineType true "MedicineType entity"
// @Success 200 {object} ent.MedicineType
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /MedicineType/{id} [put]
func (ctl *MedicineTypeController) UpdateMedicineType(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.MedicineType{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "MedicineType binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.MedicineType.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, u)
}

// NewMedicineTypeController creates and registers handles for the MedicineType controller
func NewMedicineTypeController(router gin.IRouter, client *ent.Client) *MedicineTypeController {
	uc := &MedicineTypeController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitMedicineTypeController registers routes to the main engine
func (ctl *MedicineTypeController) register() {
	MedicineType := ctl.router.Group("/MedicineType")

	MedicineType.GET("", ctl.ListMedicineType)

	// CRUD
	MedicineType.POST("", ctl.CreateMedicineType)
	MedicineType.GET(":id", ctl.GetMedicineType)
	MedicineType.PUT(":id", ctl.UpdateMedicineType)
	MedicineType.DELETE(":id", ctl.DeleteMedicineType)
}
