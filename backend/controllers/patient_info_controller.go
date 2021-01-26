package controllers

import (
	"context"
	"fmt"
	"strconv"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent"
	"github.com/gin-gonic/gin"
)

// PatientInfoController defines the struct for the PatientInfo controller
type PatientInfoController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePatientInfo handles POST requests for adding PatientInfo entities
// @Summary Create PatientInfo
// @Description Create PatientInfo
// @ID create-Ppatient-info
// @Accept   json
// @Produce  json
// @Param PatientInfo body ent.PatientInfo true "PatientInfo entity"
// @Success 200 {object} ent.PatientInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientinfos [post]
func (ctl *PatientInfoController) CreatePatientInfo(c *gin.Context) {
	obj := ent.PatientInfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Patient Info binding failed",
		})
		return
	}

	p, err := ctl.client.PatientInfo.
		Create().
		SetCardNumber(obj.CardNumber).
		SetName(obj.Name).
		SetGender(obj.Gender).
		SetAge(obj.Age).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, p)
}

// GetPatientInfo handles GET requests to retrieve a PatientInfo entity
// @Summary Get a PatientInfo entity by ID
// @Description get PatientInfo by ID
// @ID get-patient-info
// @Produce  json
// @Param id path int true "PatientInfo ID"
// @Success 200 {object} ent.PatientInfo
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientinfos/{id} [get]
func (ctl *PatientInfoController) GetPatientInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	p, err := ctl.client.PatientInfo.
		Query().
		Where(patientinfo.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, p)
}

// ListPatientInfo handles request to get a list of PatientInfo entities
// @Summary List PatientInfo entities
// @Description list PatientInfo entities
// @ID list-patient-info
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.PatientInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientinfos [get]
func (ctl *PatientInfoController) ListPatientInfo(c *gin.Context) {
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

	patientinfos, err := ctl.client.PatientInfo.
		Query().
		WithDrugallergys().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, patientinfos)
}

// DeletePatientInfo handles DELETE requests to delete a PatientInfo entity
// @Summary Delete a PatientInfo entity by ID
// @Description get PatientInfo by ID
// @ID delete-patient-info
// @Produce  json
// @Param id path int true "PatientInfo ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientinfos/{id} [delete]
func (ctl *PatientInfoController) DeletePatientInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.PatientInfo.
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

// UpdatePatientInfo handles PUT requests to update a PatientInfo entity
// @Summary Update a PatientInfo entity by ID
// @Description update PatientInfo by ID
// @ID update-PatientInfo
// @Accept   json
// @Produce  json
// @Param id path int true "PatientInfo ID"
// @Param PatientInfo body ent.PatientInfo true "PatientInfo entity"
// @Success 200 {object} ent.PatientInfo
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /patientoinfos/{id} [put]
func (ctl *PatientInfoController) UpdatePatientInfo(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.PatientInfo{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Patient Info binding failed",
		})
		return
	}
	obj.ID = int(id)
	fmt.Println(obj.ID)
	p, err := ctl.client.PatientInfo.
		UpdateOneID(int(id)).
		SetCardNumber(obj.CardNumber).
		SetName(obj.Name).
		SetGender(obj.Gender).
		SetAge(obj.Age).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, p)
}

// NewPatientInfoController creates and registers handles for the PatientInfo controller
func NewPatientInfoController(router gin.IRouter, client *ent.Client) *PatientInfoController {
	pi := &PatientInfoController{
		client: client,
		router: router,
	}
	pi.register()
	return pi
}

func (ctl *PatientInfoController) register() {
	patientinfos := ctl.router.Group("/patientinfos")

	patientinfos.GET("", ctl.ListPatientInfo)

	// CRUD
	patientinfos.POST("", ctl.CreatePatientInfo)
	patientinfos.GET(":id", ctl.GetPatientInfo)
	patientinfos.PUT(":id", ctl.UpdatePatientInfo)
	patientinfos.DELETE(":id", ctl.DeletePatientInfo)
}