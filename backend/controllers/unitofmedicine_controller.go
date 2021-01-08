package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/unitofmedicine"
)

// UnitOfMedicineController defines the struct for the UnitOfMedicine controller
type UnitOfMedicineController struct {
	client *ent.Client
	router gin.IRouter
}

// CreateUnitOfMedicine handles POST requests for adding UnitOfMedicine entities
// @Summary Create UnitOfMedicine
// @Description Create UnitOfMedicine
// @ID create-UnitOfMedicine
// @Accept   json
// @Produce  json
// @Param UnitOfMedicine body ent.UnitOfMedicine true "UnitOfMedicine entity"
// @Success 200 {object} ent.UnitOfMedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /UnitOfMedicine [post]
func (ctl *UnitOfMedicineController) CreateUnitOfMedicine(c *gin.Context) {
	obj := ent.UnitOfMedicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "UnitOfMedicine binding failed",
		})
		return
	}
	u, err := ctl.client.UnitOfMedicine.
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

// GetUnitOfMedicine handles GET requests to retrieve a UnitOfMedicine entity
// @Summary Get a UnitOfMedicine entity by ID
// @Description get UnitOfMedicine by ID
// @ID get-UnitOfMedicine
// @Produce  json
// @Param id path int true "UnitOfMedicine ID"
// @Success 200 {object} ent.UnitOfMedicine
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /UnitOfMedicine/{id} [get]
func (ctl *UnitOfMedicineController) GetUnitOfMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.UnitOfMedicine.
		Query().
		Where(unitofmedicine.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, u)
}

// ListUnitOfMedicine handles request to get a list of UnitOfMedicine entities
// @Summary List UnitOfMedicine entities
// @Description list UnitOfMedicine entities
// @ID list-UnitOfMedicine
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.UnitOfMedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /UnitOfMedicine [get]
func (ctl *UnitOfMedicineController) ListUnitOfMedicine(c *gin.Context) {
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
	UnitOfMedicine, err := ctl.client.UnitOfMedicine.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, UnitOfMedicine)
}

// DeleteUnitOfMedicine handles DELETE requests to delete a UnitOfMedicine entity
// @Summary Delete a UnitOfMedicine entity by ID
// @Description get UnitOfMedicine by ID
// @ID delete-UnitOfMedicine
// @Produce  json
// @Param id path int true "UnitOfMedicine ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /UnitOfMedicine/{id} [delete]
func (ctl *UnitOfMedicineController) DeleteUnitOfMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.UnitOfMedicine.
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

// UpdateUnitOfMedicine handles PUT requests to update a UnitOfMedicine entity
// @Summary Update a UnitOfMedicine entity by ID
// @Description update UnitOfMedicine by ID
// @ID update-UnitOfMedicine
// @Accept   json
// @Produce  json
// @Param id path int true "UnitOfMedicine ID"
// @Param UnitOfMedicine body ent.UnitOfMedicine true "UnitOfMedicine entity"
// @Success 200 {object} ent.UnitOfMedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /UnitOfMedicine/{id} [put]
func (ctl *UnitOfMedicineController) UpdateUnitOfMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.UnitOfMedicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "UnitOfMedicine binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.UnitOfMedicine.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, u)
}

// NewUnitOfMedicineController creates and registers handles for the UnitOfMedicine controller
func NewUnitOfMedicineController(router gin.IRouter, client *ent.Client) *UnitOfMedicineController {
	uc := &UnitOfMedicineController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitUnitOfMedicineController registers routes to the main engine
func (ctl *UnitOfMedicineController) register() {
	UnitOfMedicine := ctl.router.Group("/UnitOfMedicine")

	UnitOfMedicine.GET("", ctl.ListUnitOfMedicine)

	// CRUD
	UnitOfMedicine.POST("", ctl.CreateUnitOfMedicine)
	UnitOfMedicine.GET(":id", ctl.GetUnitOfMedicine)
	UnitOfMedicine.PUT(":id", ctl.UpdateUnitOfMedicine)
	UnitOfMedicine.DELETE(":id", ctl.DeleteUnitOfMedicine)
}
