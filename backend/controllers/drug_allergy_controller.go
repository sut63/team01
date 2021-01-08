package controllers

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/drugallergy"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/pharmacist"
)

// DrugAllergyController defines the struct for the drug-allergy controller
type DrugAllergyController struct {
	client *ent.Client
	router gin.IRouter
}

type DrugAllergy struct {
	Patient    int
	Medicine   int
	Pharmacist int
	DateTime   string
}

// CreateDrugAllergy handles POST requests for adding DrugAllergy entities
// @Summary Create drug-allergy
// @Description Create drug-allergy
// @ID create-drug-allergy
// @Accept   json
// @Produce  json
// @Param drug-allergy body ent.DrugAllergy true "DrugAllergy entity"
// @Success 200 {object} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugallergys [post]
func (ctl *DrugAllergyController) CreateDrugAllergy(c *gin.Context) {
	obj := DrugAllergy{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Drug Allergy binding failed",
		})
		return
	}

	p, err := ctl.client.PatientInfo.
		Query().
		Where(patientinfo.IDEQ(int(obj.Patient))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "patient not found",
		})
		return
	}

	m, err := ctl.client.Medicine.
		Query().
		Where(medicine.IDEQ(int(obj.Medicine))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "medicine not found",
		})
		return
	}

	pha, err := ctl.client.Pharmacist.
		Query().
		Where(pharmacist.IDEQ(int(obj.Pharmacist))).
		Only(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "pharmacist not found",
		})
		return
	}

	time, err := time.Parse(time.RFC3339, obj.DateTime)

	da, err := ctl.client.DrugAllergy.
		Create().
		SetDateTime(time).
		SetPatient(p).
		SetMedicine(m).
		SetPharmacist(pha).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, da)
}

// GetDrugAllergy handles GET requests to retrieve a DrugAllergy entity
// @Summary Get a DrugAllergy entity by ID
// @Description get DrugAllergy by ID
// @ID get-drug-allergy
// @Produce  json
// @Param id path int true "DrugAllergy ID"
// @Success 200 {object} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugallergys/{id} [get]
func (ctl *DrugAllergyController) GetDrugAllergy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	da, err := ctl.client.DrugAllergy.
		Query().
		Where(drugallergy.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, da)
}

// ListDrugAllergy handles request to get a list of DrugAllergy entities
// @Summary List DrugAllergy entities
// @Description list DrugAllergy entities
// @ID list-drug-allergy
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugallergys [get]
func (ctl *DrugAllergyController) ListDrugAllergy(c *gin.Context) {
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

	DrugAllergys, err := ctl.client.DrugAllergy.
		Query().
		WithPatient().
		WithPharmacist().
		WithMedicine().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, DrugAllergys)
}

// DeleteDrugAllergy handles DELETE requests to delete a DrugAllergy entity
// @Summary Delete a DrugAllergy entity by ID
// @Description get DrugAllergy by ID
// @ID delete-drug-allergy
// @Produce  json
// @Param id path int true "DrugAllergy ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugallergy/{id} [delete]
func (ctl *DrugAllergyController) DeleteDrugAllergy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.DrugAllergy.
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

// UpdateDrugAllergy handles PUT requests to update a DrugAllergy entity
// @Summary Update a DrugAllergy entity by ID
// @Description update DrugAllergy by ID
// @ID update-drug-allergy
// @Accept   json
// @Produce  json
// @Param id path int true "DrugAllergy ID"
// @Param drugallergy body ent.DrugAllergy true "DrugAllergy entity"
// @Success 200 {object} ent.DrugAllergy
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /drugallergys/{id} [put]
func (ctl *DrugAllergyController) UpdateDrugAllergy(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.DrugAllergy{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "drug allergy binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.DrugAllergy.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}

	c.JSON(200, u)
}

// NewDrugAllergyController creates and registers handles for the DrugAllergy controller
func NewDrugAllergyController(router gin.IRouter, client *ent.Client) *DrugAllergyController {
	da := &DrugAllergyController{
		client: client,
		router: router,
	}
	da.register()
	return da
}

// InitDrugAllergyController registers routes to the main engine
func (ctl *DrugAllergyController) register() {
	drugallergys := ctl.router.Group("/drugallergys")

	drugallergys.GET("", ctl.ListDrugAllergy)

	// CRUD
	drugallergys.POST("", ctl.CreateDrugAllergy)
	drugallergys.GET(":id", ctl.GetDrugAllergy)
	drugallergys.PUT(":id", ctl.UpdateDrugAllergy)
	drugallergys.DELETE(":id", ctl.DeleteDrugAllergy)
}
