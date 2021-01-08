package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/annotation"
	"github.com/sut63/team01/ent/dispensemedicine"
	"github.com/sut63/team01/ent/pharmacist"
	"github.com/sut63/team01/ent/prescription"
)

// DispenseMedicineController defines the struct for the dispensemedicine controller
type DispenseMedicineController struct {
	client *ent.Client
	router gin.IRouter
}

//DispenseMedicine Structure
type DispenseMedicine struct {
	Datetime     string
	Prescription int
	Annotation   int
	Pharmacist   int
}

// CreateDispenseMedicine handles POST requests for adding dispensemedicine entities
// @Summary Create dispensemedicine
// @Description Create dispensemedicine
// @ID create-dispensemedicine
// @Accept   json
// @Produce  json
// @Param dispensemedicine body DispenseMedicine true "DispenseMedicine entity"
// @Success 200 {object} ent.DispenseMedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dispensemedicines [post]
func (ctl *DispenseMedicineController) CreateDispenseMedicine(c *gin.Context) {
	obj := DispenseMedicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dispensemedicine binding failed",
		})
		return
	}

	psc, err := ctl.client.Prescription.
		Query().
		Where(prescription.IDEQ(int(obj.Prescription))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "prescription not found",
		})
		return
	}

	an, err := ctl.client.Annotation.
		Query().
		Where(annotation.IDEQ(int(obj.Annotation))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "annotation not found",
		})
		return
	}

	pm, err := ctl.client.Pharmacist.
		Query().
		Where(pharmacist.IDEQ(int(obj.Pharmacist))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Datetime)
	dm, err := ctl.client.DispenseMedicine.
		Create().
		SetDatetime(time).
		SetPrescription(psc).
		SetAnnotation(an).
		SetPharmacist(pm).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving dispensemedicine failed",
		})
		return
	}

	c.JSON(200, dm)
}

// GetDispenseMedicine handles GET requests to retrieve a dispensemedicine entity
// @Summary Get a dispensemedicine entity by ID
// @Description get dispensemedicine by ID
// @ID get-dispensemedicine
// @Produce  json
// @Param id path int true "DispenseMedicine ID"
// @Success 200 {object} ent.DispenseMedicine
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dispensemedicines/{id} [get]
func (ctl *DispenseMedicineController) GetDispenseMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	dm, err := ctl.client.DispenseMedicine.
		Query().
		WithPrescription().
		WithAnnotation().
		WithPharmacist().
		Where(dispensemedicine.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, dm)
}

// ListDispenseMedicine handles request to get a list of dispensemedicine entities
// @Summary List dispensemedicine entities
// @Description list dispensemedicine entities
// @ID list-dispensemedicine
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DispenseMedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dispensemedicines [get]
func (ctl *DispenseMedicineController) ListDispenseMedicine(c *gin.Context) {
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

	dispensemedicines, err := ctl.client.DispenseMedicine.
		Query().
		WithPrescription().
		WithAnnotation().
		WithPharmacist().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, dispensemedicines)
}

// DeleteDispenseMedicine handles DELETE requests to delete a dispensemedicine entity
// @Summary Delete a dispensemedicine entity by ID
// @Description get dispensemedicine by ID
// @ID delete-dispensemedicine
// @Produce  json
// @Param id path int true "DispenseMedicine ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dispensemedicines/{id} [delete]
func (ctl *DispenseMedicineController) DeleteDispenseMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.DispenseMedicine.
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

// UpdateDispenseMedicine handles PUT requests to update a dispensemedicine entity
// @Summary Update a dispensemedicine entity by ID
// @Description update dispensemedicine by ID
// @ID update-dispensemedicine
// @Accept   json
// @Produce  json
// @Param id path int true "DispenseMedicine ID"
// @Param dispensemedicine body ent.DispenseMedicine true "DispenseMedicine entity"
// @Success 200 {object} ent.DispenseMedicine
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /dispensemedicines/{id} [put]
func (ctl *DispenseMedicineController) UpdateDispenseMedicine(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := DispenseMedicine{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "dispensemedicine binding failed",
		})
		return
	}

	psc, err := ctl.client.Prescription.
		Query().
		Where(prescription.IDEQ(int(obj.Prescription))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	an, err := ctl.client.Annotation.
		Query().
		Where(annotation.IDEQ(int(obj.Annotation))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	pm, err := ctl.client.Pharmacist.
		Query().
		Where(pharmacist.IDEQ(int(obj.Pharmacist))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.Datetime)
	dm, err := ctl.client.DispenseMedicine.
		UpdateOneID(int(id)).
		SetDatetime(time).
		SetPrescription(psc).
		SetAnnotation(an).
		SetPharmacist(pm).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update dispensemedicine failed"})
		return
	}

	c.JSON(200, dm)
}

// NewDispenseMedicineController creates and registers handles for the DispenseMedicineController
func NewDispenseMedicineController(router gin.IRouter, client *ent.Client) *DispenseMedicineController {
	dmc := &DispenseMedicineController{
		client: client,
		router: router,
	}
	dmc.register()
	return dmc
}

// InitDispenseMedicineController registers routes to the main engine
func (ctl *DispenseMedicineController) register() {
	dispensemedicines := ctl.router.Group("/dispensemedicines")

	dispensemedicines.GET("", ctl.ListDispenseMedicine)

	// CRUD
	dispensemedicines.POST("", ctl.CreateDispenseMedicine)
	dispensemedicines.GET(":id", ctl.GetDispenseMedicine)
	dispensemedicines.PUT(":id", ctl.UpdateDispenseMedicine)
	dispensemedicines.DELETE(":id", ctl.DeleteDispenseMedicine)
}
