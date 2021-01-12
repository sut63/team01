package controllers

import (
	"context"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/doctor"
	"github.com/sut63/team01/ent/medicine"
	"github.com/sut63/team01/ent/patientinfo"
	"github.com/sut63/team01/ent/prescription"
)

// PrescriptionController defines the struct for the Prescription controller
type PrescriptionController struct {
	client *ent.Client
	router gin.IRouter
}

// Prescription defines the struct for the Prescription
type Prescription struct {
	DoctorID      int
	PatientInfoID int
	MedicineID    int
	Value         string
	StatusQueue  string
}

// CreatePrescription handles POST requests for adding Prescription entities
// @Summary Create Prescription
// @Description Create Prescription
// @ID create-Prescription
// @Accept   json
// @Produce  json
// @Param Prescription body Prescription true "Prescription entity"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Prescription [post]
func (ctl *PrescriptionController) CreatePrescription(c *gin.Context) {
	obj := Prescription{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Prescription binding failed",
		})
		return
	}
	D, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(obj.DoctorID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "doctor not found",
		})
		return
	}

	P, err := ctl.client.PatientInfo.
		Query().
		Where(patientinfo.IDEQ(int(obj.PatientInfoID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "patientinfo not found",
		})
		return
	}

	M, err := ctl.client.Medicine.
		Query().
		Where(medicine.IDEQ(int(obj.MedicineID))).
		Only(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "medicine not found",
		})
		return
	}
		var Value int
	if v , err := strconv.ParseInt (obj.Value,10,64);err == nil{
		Value = int(v)
	}

	pr, err := ctl.client.Prescription.
		Create().
		SetValue(int(Value)).
		SetStatusQueue("No").
		SetPrescriptiondoctor(D).
		SetPrescriptionmedicine(M).
		SetPrescriptionpatient(P).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}
	c.JSON(200, pr)
}

// GetPrescription handles GET requests to retrieve a Prescription entity
// @Summary Get a Prescription entity by ID
// @Description get Prescription by ID
// @ID get-Prescription
// @Produce  json
// @Param id path int true "Prescription ID"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Prescription/{id} [get]
func (ctl *PrescriptionController) GetPrescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Prescription.
		Query().
		Where(prescription.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, u)
}

// ListPrescription handles request to get a list of Prescription entities
// @Summary List Prescription entities
// @Description list Prescription entities
// @ID list-Prescription
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Prescription [get]
func (ctl *PrescriptionController) ListPrescription(c *gin.Context) {
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
	Prescription, err := ctl.client.Prescription.
		Query().
		WithPrescriptiondoctor().
		WithPrescriptionmedicine().
		WithPrescriptionpatient().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, Prescription)
}

/*
// UpdatePrescription handles PUT requests to update a Prescription entity
// @Summary Update a Prescription entity by ID
// @Description update Prescription by ID
// @ID update-Prescription
// @Accept   json
// @Produce  json
// @Param id path int true "Prescription ID"
// @Param Prescription body ent.Prescription true "Prescription entity"
// @Success 200 {object} ent.Prescription
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Prescription/{id} [put]
func (ctl *PrescriptionController) UpdatePrescription(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.Prescription{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Prescription binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Prescription.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, u)
}
*/

// NewPrescriptionController creates and registers handles for the Prescription controller
func NewPrescriptionController(router gin.IRouter, client *ent.Client) *PrescriptionController {
	uc := &PrescriptionController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitPrescriptionController registers routes to the main engine
func (ctl *PrescriptionController) register() {
	Prescription := ctl.router.Group("/Prescription")

	Prescription.GET("", ctl.ListPrescription)

	// CRUD
	Prescription.POST("", ctl.CreatePrescription)
	//Prescription.GET(":id", ctl.GetPrescription)
	//Prescription.PUT(":id", ctl.UpdatePrescription)

}
