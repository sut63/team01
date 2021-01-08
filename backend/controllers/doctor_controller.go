package controllers

import (
	"context"
	//"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/sut63/team01/ent"
	"github.com/sut63/team01/ent/doctor"
)

// DoctorController defines the struct for the Doctor controller
type DoctorController struct {
	client *ent.Client
	router gin.IRouter
}

// Doctor defines the struct for the Doctor
type Doctor struct {
	email              string
	password           string
	name               string
	
}

// CreateDoctor handles POST requests for adding Doctor entities
// @Summary Create Doctor
// @Description Create Doctor
// @ID create-Doctor
// @Accept   json
// @Produce  json
// @Param Doctor body ent.Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Doctor [post]
func (ctl *DoctorController) CreateDoctor(c *gin.Context) {
	obj := ent.Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor binding failed",
		})
		return
	}
	u, err := ctl.client.Doctor.
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
 
 
// GetDoctor handles GET requests to retrieve a Doctor entity
// @Summary Get a Doctor entity by ID
// @Description get Doctor by ID
// @ID get-Doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Doctor/{id} [get]
func (ctl *DoctorController) GetDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	u, err := ctl.client.Doctor.
		Query().
		Where(doctor.IDEQ(int(id))).
		Only(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(200, u)
}

// ListDoctor handles request to get a list of Doctor entities
// @Summary List Doctor entities
// @Description list Doctor entities
// @ID list-Doctor
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Doctor [get]
func (ctl *DoctorController) ListDoctor(c *gin.Context) {
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
		if err == nil {offset = int(offset64)
		}
	}
	Doctor, err := ctl.client.Doctor.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, Doctor)
}

/*
// DeleteDoctor handles DELETE requests to delete a Doctor entity
// @Summary Delete a Doctor entity by ID
// @Description get Doctor by ID
// @ID delete-Doctor
// @Produce  json
// @Param id path int true "Doctor ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Doctor/{id} [delete]
func (ctl *DoctorController) DeleteDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	err = ctl.client.Doctor.
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
*/

// UpdateDoctor handles PUT requests to update a Doctor entity
// @Summary Update a Doctor entity by ID
// @Description update Doctor by ID
// @ID update-Doctor
// @Accept   json
// @Produce  json
// @Param id path int true "Doctor ID"
// @Param Doctor body ent.Doctor true "Doctor entity"
// @Success 200 {object} ent.Doctor
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /Doctor/{id} [put]
func (ctl *DoctorController) UpdateDoctor(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	obj := ent.Doctor{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "Doctor binding failed",
		})
		return
	}
	obj.ID = int(id)
	u, err := ctl.client.Doctor.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{"error": "update failed"})
		return
	}
	c.JSON(200, u)
}

// NewDoctorController creates and registers handles for the Doctor controller
func NewDoctorController(router gin.IRouter, client *ent.Client) *DoctorController {
	uc := &DoctorController{
		client: client,
		router: router,
	}
	uc.register()
	return uc
}

// InitDoctorController registers routes to the main engine
func (ctl *DoctorController) register() {
	Doctor := ctl.router.Group("/Doctor")

	Doctor.GET("", ctl.ListDoctor)

	// CRUD
	Doctor.POST("", ctl.CreateDoctor) 
	Doctor.GET(":id", ctl.GetDoctor)
	Doctor.PUT(":id", ctl.UpdateDoctor)
	//Doctor.DELETE(":id", ctl.DeleteDoctor)
}